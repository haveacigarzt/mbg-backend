package main

import (
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

	input.Filters.Sort = app.readString(qs, "sort", "id")

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
