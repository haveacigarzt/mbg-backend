package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
	"time"
)

func (app *application) createSPPGHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama           string   `json:"nama"`
		Alamat         string   `json:"alamat"`
		SosmedURL      []string `json:"sosmed_url"`
		KepalaSPPG     string   `json:"kepala_sppg"`
		NomorTelepon   string   `json:"nomor_telepon"`
		Email          string   `json:"email"`
		Latitude       float64  `json:"latitude"`
		Longitude      float64  `json:"longitude"`
		Kecamatan_ID   int64    `json:"kecamatan_id"`
		Kelurahan_ID   int64    `json:"kelurahan_id"`
		KapasitasPorsi int      `json:"kapasitas_porsi"`
		StatusAktif    bool     `json:"status_aktif"`

		User struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	sppg := &data.SPPG{
		Nama:           &input.Nama,
		Alamat:         &input.Alamat,
		SosmedURL:      input.SosmedURL,
		KepalaSPPG:     &input.KepalaSPPG,
		NomorTelepon:   input.NomorTelepon,
		Email:          input.Email,
		Latitude:       input.Latitude,
		Longitude:      input.Longitude,
		Kecamatan_ID:   &input.Kecamatan_ID,
		Kelurahan_ID:   &input.Kelurahan_ID,
		KapasitasPorsi: input.KapasitasPorsi,
		StatusAktif:    input.StatusAktif,
	}

	v := validator.New()

	if data.ValidateSPPG(v, sppg); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user := &data.User{
		Name:      input.User.Name,
		Email:     input.User.Email,
		RoleID:    3,
		Activated: true,
	}

	err = user.Password.Set(input.User.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.InsertAndGetID(user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "Alamat email user ini sudah digunakan oleh akun lain")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	sppg.UserID = user.ID

	err = app.models.SPPG.Insert(sppg)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/sppg/%d", sppg.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{
		"sppg": sppg,
	}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listSPPGHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama         string
		Kecamatan_ID int64
		Kelurahan_ID int64
		StatusAktif  *bool
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Kecamatan_ID = app.readInt64(qs, "kecamatan_id", 0, v)
	input.Kelurahan_ID = app.readInt64(qs, "kelurahan_id", 0, v)
	input.StatusAktif = app.readOptionalBool(qs, "status_aktif", v)

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
		"status_aktif",
		"-status_aktif",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	sppg, metadata, err := app.models.SPPG.GetAll(input.Nama, input.Kecamatan_ID, input.Kelurahan_ID, input.StatusAktif, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"sppg": sppg, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getSPPGHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sppg": sppg}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateSPPGHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	if sppg.UserID != user.ID {
		app.notPermittedResponse(w, r)
		return
	}

	// cocokan sppg.user_id dengan current user id, jika tidak maka error tidak diizinkan

	var input struct {
		Nama           *string   `json:"nama"`
		Alamat         *string   `json:"alamat"`
		SosmedURL      *[]string `json:"sosmed_url"`
		KepalaSPPG     *string   `json:"kepala_sppg"`
		NomorTelepon   *string   `json:"nomor_telepon"`
		Email          *string   `json:"email"`
		Latitude       *float64  `json:"latitude"`
		Longitude      *float64  `json:"longitude"`
		Kecamatan_ID   *int64    `json:"kecamatan_id"`
		Kelurahan_ID   *int64    `json:"kelurahan_id"`
		KapasitasPorsi *int      `json:"kapasitas_porsi"`
		StatusAktif    *bool     `json:"status_aktif"`
	}
	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Nama != nil {
		sppg.Nama = input.Nama
	}

	if input.Alamat != nil {
		sppg.Alamat = input.Alamat
	}

	if input.SosmedURL != nil {
		sppg.SosmedURL = *input.SosmedURL
	}

	if input.KepalaSPPG != nil {
		sppg.KepalaSPPG = input.KepalaSPPG
	}

	if input.NomorTelepon != nil {
		sppg.NomorTelepon = *input.NomorTelepon
	}

	if input.Email != nil {
		sppg.Email = *input.Email
	}

	if input.Latitude != nil {
		sppg.Latitude = *input.Latitude
	}

	if input.Longitude != nil {
		sppg.Longitude = *input.Longitude
	}

	if input.Kecamatan_ID != nil {
		sppg.Kecamatan_ID = input.Kecamatan_ID
	}

	if input.Kelurahan_ID != nil {
		sppg.Kelurahan_ID = input.Kelurahan_ID
	}

	if input.KapasitasPorsi != nil {
		sppg.KapasitasPorsi = *input.KapasitasPorsi
	}

	if input.StatusAktif != nil {
		sppg.StatusAktif = *input.StatusAktif
	}
	// Validate the updated movie record, sending the client a 422 Unprocessable Entity
	// response if any checks fail.
	v := validator.New()
	if data.ValidateSPPG(v, sppg); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Pass the updated movie record to our new Update() method.
	err = app.models.SPPG.Update(sppg)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	sppg, err = app.models.SPPG.Get(sppg.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"sppg": sppg}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// New

func (app *application) createSPPGAlokasiHarianHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// cocokan sppg.user_id dengan current user id, jika tidak maka error tidak diizinkan
	if sppg.UserID != user.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		Jumlah  int64  `json:"jumlah"`
		Tanggal string `json:"tanggal"`
	}

	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	alokasiHarian, err := app.models.AlokasiHarian.Get(input.Tanggal, sppg.ID)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			// create
			alokasiHarian = &data.AlokasiHarian{
				SPPGID:  sppg.ID,
				Tanggal: input.Tanggal,
				Jumlah:  input.Jumlah,
			}

			if data.ValidateAlokasiHarian(v, alokasiHarian); !v.Valid() {
				app.failedValidationResponse(w, r, v.Errors)
				return
			}

			err = app.models.AlokasiHarian.Insert(alokasiHarian)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}

		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	} else {
		// update
		alokasiHarian.Jumlah = input.Jumlah

		if data.ValidateAlokasiHarian(v, alokasiHarian); !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		err = app.models.AlokasiHarian.Update(alokasiHarian)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	// ws: untuk publik
	var input2 struct {
		data.Filters
	}

	v2 := validator.New()

	qs := r.URL.Query()
	input2.Filters.Page = app.readInt(qs, "page", 1, v2)
	input2.Filters.PageSize = app.readInt(qs, "page_size", 0, v2)

	input2.Filters.Sort = app.readString(qs, "sort", "id")

	input2.Filters.SortSafelist = []string{
		"id",
		"-id",
		"jumlah",
		"-jumlah",
	}

	if data.ValidateFilters(v2, input2.Filters); !v2.Valid() {
		app.failedValidationResponse(w, r, v2.Errors)
		return
	}
	alokasiHarianAll, _, err := app.models.AlokasiHarian.GetAll(input.Tanggal, input2.Filters)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	pengeluaranHarian, _, err := app.models.PengeluaranHarian.GetAllByTanggal(input.Tanggal, input2.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	payload := map[string]any{
		"type": "keuangan:updated",
		"data": BuildRingkasan(alokasiHarianAll, pengeluaranHarian),
	}
	fmt.Println("keuangan:updated")
	jsonData, _ := json.Marshal(payload)
	app.hub.BroadcastToRoom("open", jsonData)

	err = app.writeJSON(w, http.StatusOK, envelope{"alokasi_harian": alokasiHarian}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getSPPGAlokasiHarianHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// cocokan sppg.user_id dengan current user id, jika tidak maka error tidak diizinkan
	if sppg.UserID != user.ID {
		app.notPermittedResponse(w, r)
		return
	}

	// fmt.Println("lolos cek")

	var input struct {
		Tanggal string
	}
	qs := r.URL.Query()

	input.Tanggal = app.readString(qs, "tanggal", "")

	alokasi := &data.AlokasiHarian{
		Tanggal: input.Tanggal,
	}

	v := validator.New()

	if data.ValidateGetAlokasiHarian(v, alokasi); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	var alokasiHarian *data.AlokasiHarian

	alokasiHarian, err = app.models.AlokasiHarian.Get(alokasi.Tanggal, sppg.ID)
	if err != nil {
		if !errors.Is(err, data.ErrRecordNotFound) {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"alokasi_harian": alokasiHarian,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listSPPGPengeluaranHarianHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// cocokan sppg.user_id dengan current user id, jika tidak maka error tidak diizinkan
	if sppg.UserID != user.ID {
		app.notPermittedResponse(w, r)
		return
	}

	// fmt.Println("lolos cek")

	var input struct {
		Tanggal string
		data.Filters
	}
	v := validator.New()

	qs := r.URL.Query()

	input.Tanggal = app.readString(qs, "tanggal", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"produk",
		"-produk",
		"satuan",
		"-satuan",
		"harga_satuan",
		"-harga_satuan",
		"subtotal",
		"-subtotal",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	alokasi := &data.AlokasiHarian{
		Tanggal: input.Tanggal,
	}

	if data.ValidateGetAlokasiHarian(v, alokasi); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	alokasiHarian, err := app.models.AlokasiHarian.Get(alokasi.Tanggal, sppg.ID)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	pengeluaranHarian, metadata, err := app.models.PengeluaranHarian.GetAll(alokasiHarian.ID, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "pengeluaran_harian": pengeluaranHarian}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createSPPGPengeluaranHarianHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// cocokan sppg.user_id dengan current user id, jika tidak maka error tidak diizinkan
	if sppg.UserID != user.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		AlokasiHarianID      int64   `json:"alokasi_harian_id"`
		Produk               string  `json:"produk"`
		Jumlah               int64   `json:"jumlah"`
		Satuan               string  `json:"satuan"`
		HargaSatuan          int64   `json:"harga_satuan"`
		PedagangLokalID      *int64  `json:"pedagang_lokal_id"`
		NamaPedagangNonLokal *string `json:"nama_pedagang_non_lokal"`
	}

	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// create
	pengeluaranHarian := &data.PengeluaranHarian{
		AlokasiHarianID:      input.AlokasiHarianID,
		Produk:               input.Produk,
		Jumlah:               input.Jumlah,
		Satuan:               input.Satuan,
		HargaSatuan:          input.HargaSatuan,
		PedagangLokalID:      input.PedagangLokalID,
		NamaPedagangNonLokal: input.NamaPedagangNonLokal,
	}

	v := validator.New()

	if data.ValidatePengeluaranHarian(v, pengeluaranHarian); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengeluaranID, tanggal, err := app.models.PengeluaranHarian.Insert(pengeluaranHarian)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// ws: untuk publik
	var input2 struct {
		data.Filters
	}

	v2 := validator.New()

	qs := r.URL.Query()
	input2.Filters.Page = app.readInt(qs, "page", 1, v2)
	input2.Filters.PageSize = app.readInt(qs, "page_size", 0, v2)

	input2.Filters.Sort = app.readString(qs, "sort", "id")

	input2.Filters.SortSafelist = []string{
		"id",
		"-id",
		"jumlah",
		"-jumlah",
	}

	if data.ValidateFilters(v2, input2.Filters); !v2.Valid() {
		app.failedValidationResponse(w, r, v2.Errors)
		return
	}
	alokasiHarianAll, _, err := app.models.AlokasiHarian.GetAll(tanggal, input2.Filters)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	pengeluaranHarianAll, _, err := app.models.PengeluaranHarian.GetAllByTanggal(tanggal, input2.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	payload := map[string]any{
		"type": "keuangan:updated",
		"data": BuildRingkasan(alokasiHarianAll, pengeluaranHarianAll),
	}
	jsonData, _ := json.Marshal(payload)
	app.hub.BroadcastToRoom("open", jsonData)

	if pengeluaranHarian.PedagangLokalID != nil {
		fmt.Println("pengeluaran_lokal:created")
		pengeluaran, err := app.models.PengeluaranHarian.GetLocalWithStoreCoord(pengeluaranID)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.notFoundResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}
		payload := map[string]any{
			"type": "pengeluaran_lokal:created",
			"data": pengeluaran,
		}
		jsonData, _ := json.Marshal(payload)
		app.hub.BroadcastToRoom("open", jsonData)
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pengeluaran_harian": pengeluaranHarian}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteSPPGPengeluaranHarianHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	pengeluaranID, err := app.readInt64Param(r, "pengeluaran_id")
	if err != nil || pengeluaranID < 1 {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// cocokan sppg.user_id dengan current user id, jika tidak maka error tidak diizinkan
	if sppg.UserID != user.ID {
		app.notPermittedResponse(w, r)
		return
	}

	tanggal, err := app.models.PengeluaranHarian.Delete(pengeluaranID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var input2 struct {
		data.Filters
	}

	v2 := validator.New()

	qs := r.URL.Query()
	input2.Filters.Page = app.readInt(qs, "page", 1, v2)
	input2.Filters.PageSize = app.readInt(qs, "page_size", 0, v2)

	input2.Filters.Sort = app.readString(qs, "sort", "id")

	input2.Filters.SortSafelist = []string{
		"id",
		"-id",
		"jumlah",
		"-jumlah",
	}

	if data.ValidateFilters(v2, input2.Filters); !v2.Valid() {
		app.failedValidationResponse(w, r, v2.Errors)
		return
	}
	alokasiHarianAll, _, err := app.models.AlokasiHarian.GetAll(tanggal, input2.Filters)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	pengeluaranHarianAll, _, err := app.models.PengeluaranHarian.GetAllByTanggal(tanggal, input2.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	payload := map[string]any{
		"type": "keuangan:updated",
		"data": BuildRingkasan(alokasiHarianAll, pengeluaranHarianAll),
	}
	fmt.Println("keuangan:updated")
	jsonData, _ := json.Marshal(payload)
	app.hub.BroadcastToRoom("open", jsonData)

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "pengeluaran harian successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createSPPGProduksiHarianHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// cocokan sppg.user_id dengan current user id, jika tidak maka error tidak diizinkan
	if sppg.UserID != user.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		Tanggal              string    `json:"tanggal"`
		WaktuMulai           time.Time `json:"waktu_mulai"`
		EstimasiWaktuSelesai time.Time `json:"estimasi_waktu_selesai"`
	}

	// Read the JSON request body data into the input struct.
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	produksiHarian, err := app.models.ProduksiHarian.Get(input.Tanggal, sppg.ID)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			// create
			produksiHarian = &data.ProduksiHarian{
				SPPGID:               sppg.ID,
				Tanggal:              input.Tanggal,
				WaktuMulai:           input.WaktuMulai,
				EstimasiWaktuSelesai: input.EstimasiWaktuSelesai,
			}

			if data.ValidateProduksiHarian(v, produksiHarian); !v.Valid() {
				app.failedValidationResponse(w, r, v.Errors)
				return
			}

			err = app.models.ProduksiHarian.Insert(produksiHarian)
			if err != nil {
				app.serverErrorResponse(w, r, err)
				return
			}

		default:
			app.serverErrorResponse(w, r, err)
			return
		}
	} else {
		// update
		produksiHarian.WaktuMulai = input.WaktuMulai
		produksiHarian.EstimasiWaktuSelesai = input.EstimasiWaktuSelesai

		if data.ValidateProduksiHarian(v, produksiHarian); !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}

		err = app.models.ProduksiHarian.Update(produksiHarian)
		if err != nil {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	// ws: untuk publik
	var input2 struct {
		data.Filters
	}

	v2 := validator.New()

	qs := r.URL.Query()
	input2.Filters.Page = app.readInt(qs, "page", 1, v2)
	input2.Filters.PageSize = app.readInt(qs, "page_size", 0, v2)

	input2.Filters.Sort = app.readString(qs, "sort", "id")

	input2.Filters.SortSafelist = []string{
		"id",
		"-id",
	}

	if data.ValidateFilters(v2, input2.Filters); !v2.Valid() {
		app.failedValidationResponse(w, r, v2.Errors)
		return
	}
	produksiHarianAll, _, err := app.models.ProduksiHarian.GetAll(input.Tanggal, input2.Filters)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	payload := map[string]any{
		"type": "produksi:updated",
		"data": produksiHarianAll,
	}
	fmt.Println("produksi:updated")
	jsonData, _ := json.Marshal(payload)
	app.hub.BroadcastToRoom("open", jsonData)

	err = app.writeJSON(w, http.StatusOK, envelope{"produksi_harian": produksiHarian}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getSPPGProduksiHarianHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	user := app.contextGetUser(r)
	if user.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// cocokan sppg.user_id dengan current user id, jika tidak maka error tidak diizinkan
	if sppg.UserID != user.ID {
		app.notPermittedResponse(w, r)
		return
	}

	// fmt.Println("lolos cek")

	var input struct {
		Tanggal string
	}
	qs := r.URL.Query()

	input.Tanggal = app.readString(qs, "tanggal", "")

	produksi := &data.ProduksiHarian{
		Tanggal: input.Tanggal,
	}

	v := validator.New()

	if data.ValidateGetProduksiHarian(v, produksi); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	var produksiHarian *data.ProduksiHarian

	produksiHarian, err = app.models.ProduksiHarian.Get(produksi.Tanggal, sppg.ID)
	if err != nil {
		if !errors.Is(err, data.ErrRecordNotFound) {
			app.serverErrorResponse(w, r, err)
			return
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"produksi_harian": produksiHarian,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type RingkasanAlokasi struct {
	SPPGID   int64  `json:"sppg_id"`
	RowID    int    `json:"row_id"`
	SPPGNama string `json:"sppg_nama"`
	Tanggal  string `json:"tanggal"`
	Alokasi  int64  `json:"alokasi"`
	Terpakai int64  `json:"terpakai"`
	Sisa     int64  `json:"sisa"`
}

func BuildRingkasan(
	alokasi []*data.AlokasiHarianOutput,
	pengeluaran []*data.PengeluaranHarian,
) []RingkasanAlokasi {

	terpakaiMap := make(map[int64]int64)

	for _, p := range pengeluaran {
		terpakaiMap[p.AlokasiHarianID] += p.Subtotal
	}

	result := make([]RingkasanAlokasi, 0, len(alokasi))

	for _, a := range alokasi {
		var (
			terpakai      int64
			alokasiJumlah int64
			tanggal       string
		)

		if a.ID != nil {
			terpakai = terpakaiMap[*a.ID]
		}

		if a.Jumlah != nil {
			alokasiJumlah = *a.Jumlah
		}

		if a.Tanggal != nil {
			tanggal = *a.Tanggal
		}

		result = append(result, RingkasanAlokasi{
			SPPGID:   a.SPPGID,
			RowID:    a.RowID,
			SPPGNama: a.SPPGNama,
			Tanggal:  tanggal,
			Alokasi:  alokasiJumlah,
			Terpakai: terpakai,
			Sisa:     alokasiJumlah - terpakai,
		})
	}

	return result
}

func (app *application) listKeuanganHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Tanggal string
		data.Filters
	}
	v := validator.New()

	qs := r.URL.Query()

	input.Tanggal = app.readString(qs, "tanggal", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 0, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"jumlah",
		"-jumlah",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	alokasi := &data.AlokasiHarian{
		Tanggal: input.Tanggal,
	}

	if data.ValidateGetAlokasiHarian(v, alokasi); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	alokasiHarian, metadata, err := app.models.AlokasiHarian.GetAll(alokasi.Tanggal, input.Filters)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	pengeluaranHarian, metadata, err := app.models.PengeluaranHarian.GetAllByTanggal(alokasi.Tanggal, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "keuangan_harian": BuildRingkasan(alokasiHarian, pengeluaranHarian)}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	// err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "keuangan_harian": pengeluaranHarian}, nil)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// }
}

func (app *application) listProduksiHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Tanggal string
		data.Filters
	}
	v := validator.New()

	qs := r.URL.Query()

	input.Tanggal = app.readString(qs, "tanggal", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 0, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	alokasi := &data.AlokasiHarian{
		Tanggal: input.Tanggal,
	}

	if data.ValidateGetAlokasiHarian(v, alokasi); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	produksiHarian, metadata, err := app.models.ProduksiHarian.GetAll(alokasi.Tanggal, input.Filters)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "produksi_harian": produksiHarian}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func GenerateRandomToken() (string, error) {
	b := make([]byte, 12)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (app *application) createSPPGInvitationHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID != 1 {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		NamaSPPG string `json:"nama_sppg"`
	}

	// Read the JSON request body data into the input struct.
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	token, err := GenerateRandomToken()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	invitation := &data.SPPGInvitations{
		Token:    token,
		NamaSPPG: input.NamaSPPG,
		ExpiresAt: sql.NullTime{
			Time:  time.Now().AddDate(0, 0, 7),
			Valid: true,
		},
	}

	v := validator.New()

	data.ValidateSPPGInvitations(v, invitation)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.SPPGInvitations.Insert(invitation)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// // ws: untuk publik
	// var input2 struct {
	// 	data.Filters
	// }

	// v2 := validator.New()

	// qs := r.URL.Query()
	// input2.Filters.Page = app.readInt(qs, "page", 1, v2)
	// input2.Filters.PageSize = app.readInt(qs, "page_size", 0, v2)

	// input2.Filters.Sort = app.readString(qs, "sort", "id")

	// input2.Filters.SortSafelist = []string{
	// 	"id",
	// 	"-id",
	// 	"jumlah",
	// 	"-jumlah",
	// }

	// if data.ValidateFilters(v2, input2.Filters); !v2.Valid() {
	// 	app.failedValidationResponse(w, r, v2.Errors)
	// 	return
	// }
	// alokasiHarianAll, _, err := app.models.AlokasiHarian.GetAll(tanggal, input2.Filters)
	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, data.ErrRecordNotFound):
	// 		app.notFoundResponse(w, r)
	// 	default:
	// 		app.serverErrorResponse(w, r, err)
	// 	}
	// 	return
	// }

	// pengeluaranHarianAll, _, err := app.models.PengeluaranHarian.GetAllByTanggal(tanggal, input2.Filters)
	// if err != nil {
	// 	app.serverErrorResponse(w, r, err)
	// 	return
	// }

	// payload := map[string]any{
	// 	"type": "keuangan:updated",
	// 	"data": BuildRingkasan(alokasiHarianAll, pengeluaranHarianAll),
	// }
	// fmt.Println("keuangan:updated")
	// jsonData, _ := json.Marshal(payload)
	// app.hub.BroadcastToRoom("open", jsonData)

	err = app.writeJSON(w, http.StatusOK, envelope{"invitation": invitation}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listSPPGInvitationHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.RoleID != 1 {
		app.notPermittedResponse(w, r)
		return
	}
	v := validator.New()

	var input struct {
		data.Filters
	}

	qs := r.URL.Query()

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "-created_at")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama_sppg",
		"-nama_sppg",
		"created_at",
		"-created_at",
		"used_at",
		"-used_at",
		"expires_at",
		"-expires_at",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	invitations, metadata, err := app.models.SPPGInvitations.GetAll(input.Filters)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"invitations": invitations,
		"metadata":    metadata,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getSPPGInvitationHandler(w http.ResponseWriter, r *http.Request) {
	token, err := app.readStringParam(r, "token")
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	invitation, err := app.models.SPPGInvitations.GetByToken(token)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"invitation": invitation,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteSPPGInvitationHandler(w http.ResponseWriter, r *http.Request) {
	token, err := app.readStringParam(r, "token")
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.SPPGInvitations.DeleteByToken(token)
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
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "invitation successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type createSPPGByInvitationInput struct {
	Nama           string   `json:"nama"`
	Alamat         string   `json:"alamat"`
	KepalaSPPG     string   `json:"kepala_sppg"`
	NomorTelepon   string   `json:"nomor_telepon"`
	Email          string   `json:"email"`
	KelurahanID    int64    `json:"kelurahan_id"`
	KecamatanID    int64    `json:"kecamatan_id"`
	KapasitasPorsi int      `json:"kapasitas_porsi"`
	Latitude       float64  `json:"latitude"`
	Longitude      float64  `json:"longitude"`
	SosmedURL      []string `json:"sosmed_url"`

	EmailUser string `json:"email_user"`
	Password  string `json:"password"`
}

func (app *application) createSPPGByInvitationHandler(w http.ResponseWriter, r *http.Request) {
	token, err := app.readStringParam(r, "token")
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input createSPPGByInvitationInput

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	sppg := &data.SPPG{
		Nama:           &input.Nama,
		Alamat:         &input.Alamat,
		KepalaSPPG:     &input.KepalaSPPG,
		NomorTelepon:   input.NomorTelepon,
		Email:          input.Email,
		Kecamatan_ID:   &input.KecamatanID,
		Kelurahan_ID:   &input.KelurahanID,
		KapasitasPorsi: input.KapasitasPorsi,
		Latitude:       input.Latitude,
		Longitude:      input.Longitude,
		SosmedURL:      input.SosmedURL,
		StatusAktif:    false,
	}

	v := validator.New()

	data.ValidateCreateSPPG(v, sppg)

	if !v.Valid() {
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

	invitation, err := app.models.SPPGInvitations.GetByTokenTx(ctx, tx, token)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	v = validator.New()

	data.ValidateSPPGInvitations(v, invitation)

	if invitation.UsedAt.Valid {
		v.AddError("token", "token sudah digunakan")
	}

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user := &data.User{
		Name:      *sppg.Nama,
		Email:     input.EmailUser,
		Activated: false,
		RoleID:    3,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.InsertAndGetIDTx(ctx, tx, user)
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

	sppg.UserID = user.ID

	err = app.models.SPPG.InsertTx(ctx, tx, sppg)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.models.SPPGInvitations.MarkAsUsed(ctx, tx, invitation.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{
		"sppg": sppg,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getSPPGKelengkapanDataHandler(w http.ResponseWriter, r *http.Request) {

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
		Tanggal string
	}
	qs := r.URL.Query()

	input.Tanggal = app.readString(qs, "tanggal", "")

	kelengkapan, err := app.models.SPPG.GetKelengkapan(sppg.ID, input.Tanggal)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{
		"kelengkapan_data": kelengkapan,
	}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
