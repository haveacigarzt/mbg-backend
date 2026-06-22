package main

import (
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "availablee",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) kecamatanHandler(w http.ResponseWriter, r *http.Request) {
	kecamatan, err := app.models.Kecamatan.GetAll()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"kecamatan": kecamatan}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) kelurahanHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the kecamatan ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	kelurahan, err := app.models.Kelurahan.GetAll(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"kelurahan": kelurahan}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
