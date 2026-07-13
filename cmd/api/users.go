package main

import (
	"context"
	"errors"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
	"time"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:      input.Name,
		Email:     input.Email,
		Activated: false,
		RoleID:    2,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	v := validator.New()

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "Email user yang dikirimkan sudah digunakan oleh akun lain")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Add the "sekolah:read" permission for the new user.
	// err = app.models.Permissions.AddForUser(user.ID, "sekolah:read")
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }

	// After the user record has been created in the database, generate a new activation
	// token for the user.
	token, err := app.models.Tokens.New(user.ID, 3*24*time.Hour, data.ScopeActivation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.background(func() {
		data := map[string]any{
			"activationToken": token.Plaintext,
			"userID":          user.ID,
		}
		err = app.mailer.Send(user.Email, "user_welcome.tmpl", data)
		if err != nil {
			app.logger.Error(err.Error())
		}
	})

	err = app.writeJSON(w, http.StatusAccepted, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) activateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the plaintext activation token from the request body.
	var input struct {
		TokenPlaintext string `json:"token"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Validate the plaintext token provided by the client.
	v := validator.New()
	if data.ValidateTokenPlaintext(v, input.TokenPlaintext); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Retrieve the details of the user associated with the token using the
	// GetForToken() method (which we will create in a minute). If no matching record
	// is found, then we let the client know that the token they provided is not valid.
	user, err := app.models.Users.GetForToken(data.ScopeActivation, input.TokenPlaintext)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			v.AddError("token", "invalid or expired activation token")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Update the user's activation status.
	user.Activated = true
	// Save the updated user record in our database, checking for any edit conflicts in
	// the same way that we did for our movie records.
	err = app.models.Users.Update(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// If everything went successfully, then we delete all activation tokens for the
	// user.
	err = app.models.Tokens.DeleteAllForUser(data.ScopeActivation, user.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send the updated user details to the client in a JSON response.
	err = app.writeJSON(w, http.StatusOK, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type AuthUserResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Activated bool      `json:"activated"`
	Role      RoleInfo  `json:"role"`
}

type RoleInfo struct {
	RoleID   int64 `json:"role_id"`
	IDInRole int64 `json:"id_in_role"`
}

func (app *application) authUserHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	if user.ID == 0 {
		app.authenticationRequiredResponse(w, r)
		return
	}

	idInRole := int64(0)

	if user.RoleID == 3 {
		sppg, err := app.models.SPPG.GetByUserID(user.ID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
		idInRole = sppg.ID
	}
	if user.RoleID == 4 {
		driver, err := app.models.Drivers.GetByUserID(user.ID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
		idInRole = driver.ID
	}

	response := AuthUserResponse{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Activated: user.Activated,
		Role: RoleInfo{
			RoleID:   user.RoleID,
			IDInRole: idInRole,
		},
	}

	err := app.writeJSON(w, http.StatusOK, envelope{
		"user": response,
	}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getAkunHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	if user.ID == 0 {
		app.authenticationRequiredResponse(w, r)
		return
	}

	var input struct {
		Name     string
		Status   string
		RoleID   string
		Instansi string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.Status = app.readString(qs, "status", "")
	input.RoleID = app.readString(qs, "role_id", "")
	input.Instansi = app.readString(qs, "instansi", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"name",
		"-name",
		"email",
		"-email",
		"role_id",
		"-role_id",
		"terakhir_aktif",
		"-terakhir_aktif",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// check if input.Status either "aktif", atau "nonaktif", selain itu error
	if input.Status != "" && input.Status != "aktif" && input.Status != "nonaktif" {
		v.AddError("status", "must be either 'aktif' or 'nonaktif'")
	}
	// check if input.RoleID either "1", "2", "3" "4", selain itu error
	validRoles := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": true,
	}

	if input.RoleID != "" && !validRoles[input.RoleID] {
		v.AddError("role_id", "invalid role_id")
	}
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// fmt.Println(input.Status)
	// fmt.Println(input.RoleID)

	akun, metadata, err := app.models.Users.GetAll(input.Name, input.Status, input.RoleID, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "akun": akun}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getAkunSummaryHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	if user.ID == 0 {
		app.authenticationRequiredResponse(w, r)
		return
	}

	summary, err := app.models.Users.GetSummary()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"summary": summary}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createAkunHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	if user.ID == 0 {
		app.authenticationRequiredResponse(w, r)
		return
	}

	var input struct {
		Name     string `json:"name"`
		RoleID   int64  `json:"role_id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateAkunInput(v, input); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	newUser := &data.User{
		Name:      input.Name,
		Email:     input.Email,
		RoleID:    input.RoleID,
		Activated: true,
	}

	err = newUser.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	data.ValidateUser(v, newUser)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	tx, err := app.models.DB.BeginTx(ctx, nil)
	if err != nil {
		tx.Rollback()
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.models.Users.InsertAndGetIDTx(ctx, tx, newUser)
	if err != nil {
		tx.Rollback()

		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "Email user yang dikirimkan sudah digunakan oleh akun lain")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	if input.RoleID == 3 {
		sppg := &data.SPPG{
			UserID:         newUser.ID,
			SosmedURL:      []string{},
			StatusAktif:    true,
			KapasitasPorsi: 0,
			Kecamatan_ID:   nil,
			Kelurahan_ID:   nil,
			Nama:           &newUser.Name,
		}

		err = app.models.SPPG.InsertTx(ctx, tx, sppg)
		if err != nil {
			tx.Rollback()
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{
		"message": "User created",
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateAkunHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)

	if user.ID == 0 {
		app.authenticationRequiredResponse(w, r)
		return
	}

	var input data.AkunUpdate

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateAkunUpdate(v, input); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Ambil user yang akan diupdate
	account, err := app.models.Users.Get(input.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	account.Activated = *input.Activated

	err = app.models.Users.Update(account)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"message": "User updated",
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteAkunHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)

	if currUser.ID == 0 {
		app.authenticationRequiredResponse(w, r)
		return
	}

	var input data.AkunDelete

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if data.ValidateAkunDelete(v, input); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.Get(input.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.models.Users.Delete(user)

	err = app.writeJSON(w, http.StatusOK, envelope{
		"message": "User deleted",
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
