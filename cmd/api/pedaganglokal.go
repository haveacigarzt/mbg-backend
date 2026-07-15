package main

import (
	"errors"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
)

func (app *application) listPedagangLokalHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama        string
		JenisProduk string
		SPPG        string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.JenisProduk = app.readString(qs, "jenis_produk", "")
	input.SPPG = app.readString(qs, "sppg", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "-created_at")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
		"alamat",
		"-alamat",
		"jenis_produk",
		"-jenis_produk",
		"created_at",
		"-created_at",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pedagang, metadata, err := app.models.PedagangLokal.GetAll(
		input.Nama,
		input.JenisProduk,
		input.SPPG,
		input.Filters,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"pedagang_lokal": pedagang,
		"metadata":       metadata,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createPedagangLokalHandler(w http.ResponseWriter, r *http.Request) {
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
		Nama        string  `json:"nama"`
		Alamat      string  `json:"alamat"`
		NoHP        string  `json:"no_hp"`
		Longitude   float64 `json:"longitude"`
		Latitude    float64 `json:"latitude"`
		JenisProduk string  `json:"jenis_produk"`
	}

	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pedagangLokal := &data.PedagangLokal{
		Nama:        input.Nama,
		Alamat:      input.Alamat,
		NoHP:        input.NoHP,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		JenisProduk: input.JenisProduk,
		SPPGID:      sppg.ID,
	}

	v := validator.New()

	if data.ValidatePedagangLokal(v, pedagangLokal); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.PedagangLokal.Insert(pedagangLokal)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"message": "pedagang lokal successfully created"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePedagangLokalHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	pedagangLokal, err := app.models.PedagangLokal.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if sppg.ID != pedagangLokal.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		Nama        *string  `json:"nama"`
		Alamat      *string  `json:"alamat"`
		NoHP        *string  `json:"no_hp"`
		Longitude   *float64 `json:"longitude"`
		Latitude    *float64 `json:"latitude"`
		JenisProduk *string  `json:"jenis_produk"`
	}

	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Nama != nil {
		pedagangLokal.Nama = *input.Nama
	}
	if input.Alamat != nil {
		pedagangLokal.Alamat = *input.Alamat
	}
	if input.NoHP != nil {
		pedagangLokal.NoHP = *input.NoHP
	}
	if input.Longitude != nil {
		pedagangLokal.Longitude = *input.Longitude
	}
	if input.Latitude != nil {
		pedagangLokal.Latitude = *input.Latitude
	}
	if input.JenisProduk != nil {
		pedagangLokal.JenisProduk = *input.JenisProduk
	}

	v := validator.New()

	if data.ValidatePedagangLokal(v, pedagangLokal); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.PedagangLokal.Update(pedagangLokal)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "pedagang lokal successfully updated"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deletePedagangLokalHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	pedagangLokal, err := app.models.PedagangLokal.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if sppg.ID != pedagangLokal.SPPGID {
		app.notPermittedResponse(w, r)
		return
	}

	err = app.models.PedagangLokal.Delete(id)
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
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "pedagang lokal successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
