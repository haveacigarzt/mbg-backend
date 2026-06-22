package main

import (
	"errors"
	"fmt"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
)

func (app *application) createSekolahHandler(w http.ResponseWriter, r *http.Request) {
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
		Nama         string  `json:"nama"`
		Alamat       string  `json:"alamat"`
		Tingkat      string  `json:"tingkat"`
		JumlahSiswa  int     `json:"jumlah_siswa"`
		Kecamatan_ID int64   `json:"kecamatan_id"`
		Kelurahan_ID int64   `json:"kelurahan_id"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Copy the values from the input struct to a new Movie struct.
	sekolah := &data.Sekolah{
		Nama:         input.Nama,
		Alamat:       input.Alamat,
		Tingkat:      input.Tingkat,
		Kecamatan_ID: input.Kecamatan_ID,
		Kelurahan_ID: input.Kelurahan_ID,
		JumlahSiswa:  input.JumlahSiswa,
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
		SPPGID:       sppg.ID,
	}

	// Initialize a new Validator instance.
	v := validator.New()

	if data.ValidateSekolah(v, sekolah); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Sekolah.Insert(sekolah)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/sekolah/%d", sekolah.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"sekolah": sekolah}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getSekolahHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	sekolah, err := app.models.Sekolah.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sekolah": sekolah}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateSekolahHandler(w http.ResponseWriter, r *http.Request) {
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

	sekolah, err := app.models.Sekolah.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if sppg.ID != sekolah.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		Nama         *string  `json:"nama"`
		Alamat       *string  `json:"alamat"`
		Tingkat      *string  `json:"tingkat"`
		JumlahSiswa  *int     `json:"jumlah_siswa"`
		Kecamatan_ID *int64   `json:"kecamatan_id"`
		Kelurahan_ID *int64   `json:"kelurahan_id"`
		Latitude     *float64 `json:"latitude"`
		Longitude    *float64 `json:"longitude"`
	}
	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Nama != nil {
		sekolah.Nama = *input.Nama
	}

	if input.Alamat != nil {
		sekolah.Alamat = *input.Alamat
	}

	if input.Tingkat != nil {
		sekolah.Tingkat = *input.Tingkat
	}

	if input.Kecamatan_ID != nil {
		sekolah.Kecamatan_ID = *input.Kecamatan_ID
	}

	if input.Kelurahan_ID != nil {
		sekolah.Kelurahan_ID = *input.Kelurahan_ID
	}

	if input.JumlahSiswa != nil {
		sekolah.JumlahSiswa = *input.JumlahSiswa
	}

	if input.Latitude != nil {
		sekolah.Latitude = *input.Latitude
	}

	if input.Longitude != nil {
		sekolah.Longitude = *input.Longitude
	}
	// Validate the updated movie record, sending the client a 422 Unprocessable Entity
	// response if any checks fail.
	v := validator.New()
	if data.ValidateSekolah(v, sekolah); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Pass the updated movie record to our new Update() method.
	err = app.models.Sekolah.Update(sekolah)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	sekolah, err = app.models.Sekolah.Get(sekolah.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Write the updated movie record in a JSON response.
	err = app.writeJSON(w, http.StatusOK, envelope{"sekolah": sekolah}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteSekolahHandler(w http.ResponseWriter, r *http.Request) {
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

	sekolah, err := app.models.Sekolah.Get(id)
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
	// fmt.Println("sekolah SPPG ID =", sekolah.SPPGID)

	if sppg.ID != sekolah.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	err = app.models.Sekolah.Delete(id)
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
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "sekolah successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) listSekolahHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama         string
		Tingkat      string
		Kecamatan_ID int64
		Kelurahan_ID int64
		SPPGID       int64
		JumlahSiswa  int
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Tingkat = app.readString(qs, "tingkat", "")
	input.Kecamatan_ID = app.readInt64(qs, "kecamatan_id", 0, v)
	input.Kelurahan_ID = app.readInt64(qs, "kelurahan_id", 0, v)
	input.SPPGID = app.readInt64(qs, "sppg_id", 0, v)
	input.JumlahSiswa = app.readInt(qs, "jumlah_siswa", 0, v)
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
		"jumlah_siswa",
		"-jumlah_siswa",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	sekolah, metadata, err := app.models.Sekolah.GetAll(input.Nama, input.Tingkat, input.Kecamatan_ID, input.Kelurahan_ID, input.SPPGID, input.JumlahSiswa, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"sekolah": sekolah, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
