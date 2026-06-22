package main

import (
	"context"
	"errors"
	"fmt"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
	"time"
)

func (app *application) createDriverHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.GetByUserID(currUser.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Nama         string `json:"nama"`
		NomorTelepon string `json:"nomor_telepon"`

		User struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	driver := &data.Driver{
		Nama:         input.Nama,
		NomorTelepon: input.NomorTelepon,
		StatusAktif:  true,
		SPPGID:       sppg.ID,
	}

	driverValidator := validator.New()

	if data.ValidateDriver(driverValidator, driver); !driverValidator.Valid() {
		app.failedValidationResponse(w, r, driverValidator.Errors)
		return
	}

	user := &data.User{
		Name:      input.User.Name,
		Email:     input.User.Email,
		RoleID:    4,
		Activated: true,
	}

	err = user.Password.Set(input.User.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	userValidator := validator.New()

	if data.ValidateUser(userValidator, user); !userValidator.Valid() {
		app.failedValidationResponse(w, r, userValidator.Errors)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	tx, err := app.models.DB.BeginTx(ctx, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	defer tx.Rollback()

	err = app.models.Users.InsertTx(ctx, tx, user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			userValidator.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, userValidator.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	driver.UserID = user.ID

	err = app.models.Drivers.InsertTx(ctx, tx, driver)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicatePhoneNumber):
			driverValidator.AddError(
				"nomor_telepon",
				"a driver with this phone number already exists",
			)
			app.failedValidationResponse(w, r, driverValidator.Errors)

		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/drivers/%d", driver.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{
		"driver": driver,
	}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listDriverHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.GetByUserID(user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		StatusAktif *bool
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.StatusAktif = app.readOptionalBool(qs, "status_aktif", v)
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"status_aktif",
		"-status_aktif",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	drivers, metadata, err := app.models.Drivers.GetAll(sppg.ID, input.StatusAktif, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "drivers": drivers}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getDriverHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.GetByUserID(user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	driver, err := app.models.Drivers.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if sppg.ID != driver.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"driver": driver}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getDriverMeHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID != 4 {
		app.notPermittedResponse(w, r)
		return
	}

	driver, err := app.models.Drivers.GetByUserIDJoinSPPG(user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"driver": driver}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateDriverHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.GetByUserID(user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	driver, err := app.models.Drivers.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if sppg.ID != driver.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		Nama         *string `json:"nama"`
		NomorTelepon *string `json:"nomor_telepon"`
		StatusAktif  *bool   `json:"status_aktif"`
	}
	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Nama != nil {
		driver.Nama = *input.Nama
	}

	if input.NomorTelepon != nil {
		driver.NomorTelepon = *input.NomorTelepon
	}

	if input.StatusAktif != nil {
		driver.StatusAktif = *input.StatusAktif
	}

	// Validate the updated movie record, sending the client a 422 Unprocessable Entity
	// response if any checks fail.
	v := validator.New()
	if data.ValidateDriver(v, driver); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Pass the updated movie record to our new Update() method.
	err = app.models.Drivers.Update(driver)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"driver": driver}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteDriverHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.GetByUserID(user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	driver, err := app.models.Drivers.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	count, err := app.models.Pengiriman.CountByDriverID(driver.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if count > 0 {
		app.badRequestResponse(w, r, errors.New("Driver masih memiliki riwayat pengiriman"))
		return
	}

	if sppg.ID != driver.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	tx, err := app.models.DB.BeginTx(ctx, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	defer tx.Rollback()

	err = app.models.Drivers.DeleteTx(ctx, tx, id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.models.Users.DeleteTx(ctx, tx, driver.UserID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Return a 200 OK status code along with a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "driver successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
