package main

import (
	"errors"
	"fmt"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
)

func (app *application) createPosyanduHandler(w http.ResponseWriter, r *http.Request) {
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
		Nama           string  `json:"nama"`
		Alamat         string  `json:"alamat"`
		Kecamatan_ID   int64   `json:"kecamatan_id"`
		Kelurahan_ID   int64   `json:"kelurahan_id"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		JumlahBalita   int     `json:"jumlah_balita"`
		JumlahIbuHamil int     `json:"jumlah_ibu_hamil"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Movie struct.
	posyandu := &data.Posyandu{
		Nama:           input.Nama,
		Alamat:         input.Alamat,
		Kecamatan_ID:   input.Kecamatan_ID,
		Kelurahan_ID:   input.Kelurahan_ID,
		JumlahBalita:   input.JumlahBalita,
		JumlahIbuHamil: input.JumlahIbuHamil,
		Latitude:       input.Latitude,
		Longitude:      input.Longitude,
		SPPGID:         sppg.ID,
	}

	// Initialize a new Validator instance.
	v := validator.New()

	if data.ValidatePosyandu(v, posyandu); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Posyandu.Insert(posyandu)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posyandu/%d", posyandu.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"posyandu": posyandu}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getPosyanduHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	posyandu, err := app.models.Posyandu.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"posyandu": posyandu}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePosyanduHandler(w http.ResponseWriter, r *http.Request) {
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

	posyandu, err := app.models.Posyandu.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if sppg.ID != posyandu.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		Nama           *string  `json:"nama"`
		Alamat         *string  `json:"alamat"`
		Kecamatan_ID   *int64   `json:"kecamatan_id"`
		Kelurahan_ID   *int64   `json:"kelurahan_id"`
		Latitude       *float64 `json:"latitude"`
		Longitude      *float64 `json:"longitude"`
		JumlahBalita   *int     `json:"jumlah_balita"`
		JumlahIbuHamil *int     `json:"jumlah_ibu_hamil"`
	}
	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Nama != nil {
		posyandu.Nama = *input.Nama
	}

	if input.Alamat != nil {
		posyandu.Alamat = *input.Alamat
	}

	if input.Kecamatan_ID != nil {
		posyandu.Kecamatan_ID = *input.Kecamatan_ID
	}

	if input.Kelurahan_ID != nil {
		posyandu.Kelurahan_ID = *input.Kelurahan_ID
	}

	if input.Latitude != nil {
		posyandu.Latitude = *input.Latitude
	}

	if input.Longitude != nil {
		posyandu.Longitude = *input.Longitude
	}

	if input.JumlahBalita != nil {
		posyandu.JumlahBalita = *input.JumlahBalita
	}

	if input.JumlahIbuHamil != nil {
		posyandu.JumlahIbuHamil = *input.JumlahIbuHamil
	}

	// Validate the updated movie record, sending the client a 422 Unprocessable Entity
	// response if any checks fail.
	v := validator.New()
	if data.ValidatePosyandu(v, posyandu); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Pass the updated movie record to our new Update() method.
	err = app.models.Posyandu.Update(posyandu)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	posyandu, err = app.models.Posyandu.Get(posyandu.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"posyandu": posyandu}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deletePosyanduHandler(w http.ResponseWriter, r *http.Request) {
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

	posyandu, err := app.models.Posyandu.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// fmt.Println("sppg ID =", sppg.ID)
	// fmt.Println("posyandu SPPG ID =", posyandu.SPPGID)

	if sppg.ID != posyandu.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	err = app.models.Posyandu.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Return a 200 OK status code along with a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "posyandu successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) listPosyanduHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama           string
		Kecamatan_ID   int64
		Kelurahan_ID   int64
		SPPGID         int64
		JumlahBalita   int
		JumlahIbuHamil int
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Kecamatan_ID = app.readInt64(qs, "kecamatan_id", 0, v)
	input.Kelurahan_ID = app.readInt64(qs, "kelurahan_id", 0, v)
	input.SPPGID = app.readInt64(qs, "sppg_id", 0, v)
	input.JumlahBalita = app.readInt(qs, "jumlah_balita", 0, v)
	input.JumlahIbuHamil = app.readInt(qs, "jumlah_ibu_hamil", 0, v)
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
		"alamat",
		"-alamat",
		"jumlah_balita",
		"-jumlah_balita",
		"jumlah_ibu_hamil",
		"-jumlah_ibu_hamil",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	posyandu, metadata, err := app.models.Posyandu.GetAll(input.Nama, input.Kecamatan_ID, input.Kelurahan_ID, input.SPPGID, input.JumlahBalita, input.JumlahIbuHamil, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "posyandu": posyandu}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
