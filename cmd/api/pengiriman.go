package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
	"time"
)

func (app *application) createPengirimanHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	fmt.Println(user.RoleID)
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

	var input data.CreatePengirimanInput

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	for _, tujuan := range input.Tujuan {
		if tujuan.TujuanType == "sekolah" {
			sekolah, err := app.models.Sekolah.Get(tujuan.TujuanID)
			if err != nil {
				app.notFoundResponse(w, r)
				return
			}

			fmt.Println("sekolahsppgid: ", sekolah.SPPGID, "sppgid: ", sppg.ID)

			if sekolah.SPPGID != sppg.ID {
				app.notPermittedResponse(w, r)
				return
			}
		}

		if tujuan.TujuanType == "posyandu" {
			posyandu, err := app.models.Posyandu.Get(tujuan.TujuanID)
			if err != nil {
				app.notFoundResponse(w, r)
				return
			}
			fmt.Println("posyandusppgid: ", posyandu.SPPGID, "sppgid: ", sppg.ID)

			if posyandu.SPPGID != sppg.ID {
				app.notPermittedResponse(w, r)
				return
			}
		}
	}

	// Copy the values from the input struct to a new Movie struct.
	pengiriman := &data.CreatePengirimanInput{
		Status: "menunggu",
		SPPGID: sppg.ID,
		Tujuan: input.Tujuan,
	}

	// Initialize a new Validator instance.
	v := validator.New()

	if data.ValidateInputPengiriman(v, pengiriman); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
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

	for _, tujuan := range input.Tujuan {
		pengiriman := &data.Pengiriman{
			SPPGID:     sppg.ID,
			TujuanType: tujuan.TujuanType,
			TujuanID:   tujuan.TujuanID,
			Status:     "menunggu",
		}

		err := app.models.Pengiriman.InsertTx(ctx, tx, pengiriman)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// ws
	payload := map[string]any{
		"type": "pengiriman:create",
	}
	jsonData, _ := json.Marshal(payload)
	room := fmt.Sprintf("sppg:%d", pengiriman.SPPGID)
	app.hub.BroadcastToRoom(room, jsonData)
	fmt.Println("pengiriman:create")

	headers := make(http.Header)

	err = app.writeJSON(w, http.StatusCreated, envelope{"pengiriman": pengiriman}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPengirimanHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Status     string
		SPPGID     int64
		TujuanType string
		Tanggal    string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Status = app.readString(qs, "status", "")
	input.SPPGID = app.readInt64(qs, "sppg_id", 0, v)
	input.TujuanType = app.readString(qs, "tujuan_type", "")
	input.Tanggal = app.readString(qs, "tanggal", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"status",
		"-status",
		"sppg_id",
		"-sppg_id",
		"tujuan_type",
		"-tujuan_type",
		"tujuan_nama",
		"-tujuan_nama",
		"created_at",
		"-created_at",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID == 4 {
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
		input.SPPGID = driver.SPPGID
	}

	var start *time.Time
	var end *time.Time

	if input.Tanggal != "" {
		t, err := time.Parse("2006-01-02", input.Tanggal)
		if err != nil {
			v.AddError("tanggal", "format tanggal harus YYYY-MM-DD")
		}

		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		s := t
		e := t.AddDate(0, 0, 1)

		start = &s
		end = &e
	}

	pengiriman, metadata, err := app.models.Pengiriman.GetAll(input.Status, input.SPPGID, start, end, input.TujuanType, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "pengiriman": pengiriman}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getPengirimanHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
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

	user := app.contextGetUser(r)
	if user.RoleID == 4 {
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
		if pengiriman.SPPGID != driver.SPPGID {
			app.notPermittedResponse(w, r)
			return
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pengiriman": pengiriman}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getPengirimanAktifByDriverHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID != 4 {
		app.notPermittedResponse(w, r)
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

	pengiriman, err := app.models.Pengiriman.GetAktifByDriverID(driver.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pengiriman": pengiriman}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePengirimanHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	user := app.contextGetUser(r)
	fmt.Println("roleid: ", user.RoleID)
	if user.RoleID != 4 {
		app.notPermittedResponse(w, r)
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

	fmt.Println("driverID: ", driver.ID)

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

	if pengiriman.Status != "menunggu" || pengiriman.DriverID != nil {
		fmt.Println(*pengiriman.DriverID, " and ", driver.ID)
		if *pengiriman.DriverID != driver.ID {
			app.notPermittedResponse(w, r)
			return
		}
	} else {
		fmt.Println("menunggu")
		exist, err := app.models.Pengiriman.GetIsDelivering(driver.ID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
		if exist {
			fmt.Println("exist")
			app.badRequestResponse(
				w,
				r,
				errors.New("Anda masih memiliki pengiriman yang belum selesai"),
			)
			return
		}
	}

	var input struct {
		Status string `json:"status"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Status == pengiriman.Status {
		app.badRequestResponse(
			w,
			r,
			errors.New("status pengiriman tidak berubah"),
		)
		return
	}

	switch pengiriman.Status {

	case "menunggu":
		switch input.Status {

		case "berangkat":
			now := time.Now()
			pengiriman.WaktuBerangkat = &now

		case "dibatalkan":
			now := time.Now()
			pengiriman.WaktuSelesai = &now

		default:
			app.badRequestResponse(
				w,
				r,
				errors.New(
					"status hanya dapat diubah menjadi berangkat atau dibatalkan",
				),
			)
			return
		}

	case "berangkat":
		switch input.Status {

		case "sampai":
			now := time.Now()
			pengiriman.WaktuSelesai = &now

		case "dibatalkan":
			now := time.Now()
			pengiriman.WaktuSelesai = &now

		default:
			app.badRequestResponse(
				w,
				r,
				errors.New(
					"status hanya dapat diubah menjadi sampai atau dibatalkan",
				),
			)
			return
		}

	case "sampai":
		app.badRequestResponse(
			w,
			r,
			errors.New("pengiriman sudah sampai"),
		)
		return

	case "dibatalkan":
		app.badRequestResponse(
			w,
			r,
			errors.New("pengiriman sudah dibatalkan"),
		)
		return

	default:
		app.serverErrorResponse(
			w,
			r,
			fmt.Errorf("unknown status: %s", pengiriman.Status),
		)
		return
	}

	pengiriman.Status = input.Status
	pengiriman.DriverID = &driver.ID
	pengiriman.DriverNama = &driver.Nama

	if input.Status == "sampai" || input.Status == "dibatalkan" {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		tx, err := app.models.DB.BeginTx(ctx, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
		defer tx.Rollback()

		err = app.models.Pengiriman.UpdateTx(ctx, tx, pengiriman)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrEditConflict):
				app.editConflictResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}

		err = app.models.Tracking.DeleteByPengirimanTx(ctx, tx, pengiriman.ID)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

		err = tx.Commit()
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}

	} else {

		err = app.models.Pengiriman.Update(pengiriman)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrEditConflict):
				app.editConflictResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}
	}

	// ws
	payload := map[string]any{
		"type": "pengiriman:update",
		"data": map[string]any{
			"pengiriman_id":   pengiriman.ID,
			"status":          pengiriman.Status,
			"waktu_berangkat": pengiriman.WaktuBerangkat,
			"waktu_selesai":   pengiriman.WaktuSelesai,
			"driver_nama":     pengiriman.DriverNama,
		},
	}

	jsonData, _ := json.Marshal(payload)
	room := fmt.Sprintf("sppg:%d", pengiriman.SPPGID)
	app.hub.BroadcastToRoom(room, jsonData)
	app.hub.BroadcastToRoom("open", jsonData)
	fmt.Println("pengiriman:update")

	err = app.writeJSON(w, http.StatusOK, envelope{"pengiriman": pengiriman}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
