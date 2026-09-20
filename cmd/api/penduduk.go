package main

import (
	"errors"
	"mbg/internal/data"
	"net/http"
)

func (app *application) getPendudukHandler(w http.ResponseWriter, r *http.Request) {

	nik, err := app.readStringParam(r, "nik")
	if err != nil || len(nik) != 16 {
		err = app.writeJSON(w, http.StatusOK, envelope{"penduduk": nil}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	penduduk, err := app.models.Penduduk.Get(nik)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			err = app.writeJSON(w, http.StatusOK, envelope{"penduduk": nil}, nil)
			if err != nil {
				app.serverErrorResponse(w, r, err)
			}
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"penduduk": penduduk}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getPesertaDidikByNISNHandler(w http.ResponseWriter, r *http.Request) {

	nisn, err := app.readStringParam(r, "nisn")
	if err != nil || len(nisn) != 10 {
		err = app.writeJSON(w, http.StatusOK, envelope{"peserta_didik": nil}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	peserta_didik, err := app.models.PesertaDidik.GetByNISN(nisn)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			err = app.writeJSON(w, http.StatusOK, envelope{"peserta_didik": nil}, nil)
			if err != nil {
				app.serverErrorResponse(w, r, err)
			}
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"peserta_didik": peserta_didik}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
