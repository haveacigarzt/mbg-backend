package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
	"strconv"
	"strings"
)

func (app *application) createTrackingHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID != 4 {
		app.notPermittedResponse(w, r)
		return
	}

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	driver, err := app.models.Drivers.GetByUserID(user.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	pengiriman, err := app.models.Pengiriman.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if pengiriman.Status != "berangkat" {
		app.badRequestResponse(
			w,
			r,
			errors.New("pengiriman tidak sedang berlangsung"),
		)
		return
	}

	if pengiriman.DriverID == nil || driver.ID != *pengiriman.DriverID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Speed     float64 `json:"speed"`
		Accuracy  float64 `json:"accuracy"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	tracking := &data.Tracking{
		Latitude:     input.Latitude,
		Longitude:    input.Longitude,
		Speed:        input.Speed,
		Accuracy:     input.Accuracy,
		PengirimanID: id,
	}

	trackingValidator := validator.New()

	if data.ValidateTracking(trackingValidator, tracking); !trackingValidator.Valid() {
		app.failedValidationResponse(w, r, trackingValidator.Errors)
		return
	}

	err = app.models.Tracking.Insert(tracking)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// fmt.Println(tracking.ID)

	tracking, err = app.models.Tracking.Get(tracking.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// ws: untuk publik
	payload := map[string]any{
		"type": "tracking:created",
		"data": tracking,
	}
	fmt.Println("tracking:created", tracking.ID)
	jsonData, _ := json.Marshal(payload)
	// room := fmt.Sprintf("sppg:%d", pengiriman.SPPGID)
	app.hub.BroadcastToRoom("open", jsonData)

	err = app.writeJSON(w, http.StatusCreated, envelope{
		"tracking": tracking,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getTrackingHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID == 4 {
		app.notPermittedResponse(w, r)
		return
	}

	type inputGet struct {
		PengirimanIDs []int64
		data.Filters
	}

	var input inputGet
	v := validator.New()

	qs := r.URL.Query()

	ids := qs.Get("pengiriman_ids")

	if ids != "" {
		for _, idStr := range strings.Split(ids, ",") {
			id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
			if err != nil {
				v.AddError("pengiriman_ids", "must contain valid integer IDs")
				break
			}

			input.PengirimanIDs = append(input.PengirimanIDs, id)
		}
	}

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"created_at",
		"-created_at",
	}

	// fmt.Println("validating")

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// fmt.Println(input)

	tracking, metadata, err := app.models.Tracking.GetAll(input.PengirimanIDs, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{
		"tracking": tracking, "metadata": metadata,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
