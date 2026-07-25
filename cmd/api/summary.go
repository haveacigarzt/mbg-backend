package main

import (
	"errors"
	"mbg/internal/data"
	"net/http"
)

func (app *application) sumPenerimaManfaatHandler(w http.ResponseWriter, r *http.Request) {

	summary, err := app.models.Summary.GetSummaryPenerimaManfaat()
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"summary_penerima_manfaat": summary}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) sumDapurHandler(w http.ResponseWriter, r *http.Request) {

	summary, err := app.models.Summary.GetSummaryDapur()
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"summary_dapur": summary}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
