package main

import (
	"context"
	"errors"
	"fmt"
	"mbg/internal/data"
	"mbg/internal/validator"
	"net/http"
	"time"
)

func (app *application) createPosyanduHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.GetByUserID(currUser.ID)
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
		User           struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		} `json:"user"`
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

	user := &data.User{
		Name:      input.User.Name,
		Email:     input.User.Email,
		RoleID:    5,
		Activated: true,
	}

	err = user.Password.Set(input.User.Password)
	if err != nil {
		fmt.Println("error set password")
		app.serverErrorResponse(w, r, err)
		return
	}

	userValidator := validator.New()

	if data.ValidateUser(userValidator, user); !userValidator.Valid() {
		app.failedValidationResponse(w, r, userValidator.Errors)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	tx, err := app.models.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("error begintx")
		app.serverErrorResponse(w, r, err)
		return
	}
	defer tx.Rollback()

	err = app.models.Users.InsertTx(ctx, tx, user)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			userValidator.AddError("email", "Email user yang dikirimkan sudah digunakan oleh akun lain")
			app.failedValidationResponse(w, r, userValidator.Errors)
		default:
			fmt.Println("error user inserttx")
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	posyandu.UserID = &user.ID

	err = app.models.Posyandu.InsertTx(ctx, tx, posyandu)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error commit")
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

func (app *application) createBusuiHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 5 {
		fmt.Println("error role id != 5")
		app.notPermittedResponse(w, r)
		return
	}

	posyandu, err := app.models.Posyandu.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get posyandu")
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if id != posyandu.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		// Data penduduk
		Penduduk struct {
			NIK          string `json:"nik"`
			Nama         string `json:"nama"`
			JenisKelamin string `json:"jenis_kelamin"`
			TanggalLahir string `json:"tanggal_lahir"`
			KelurahanID  int64  `json:"kelurahan_id"`
			Alamat       string `json:"alamat"`
			NoHP         string `json:"no_hp"`
		} `json:"penduduk"`

		// Data busui
		Busui struct {
			TanggalPersalinan string `json:"tanggal_persalinan"`
			AnakKe            int8   `json:"anak_ke"`
			AsiEksklusif      bool   `json:"asi_eksklusif"`
		} `json:"busui"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	tanggalLahir, err := time.Parse("2006-01-02", input.Penduduk.TanggalLahir)
	if err != nil {
		v.AddError("tanggal_lahir", "format tanggal harus YYYY-MM-DD")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	penduduk := &data.Penduduk{
		NIK:          input.Penduduk.NIK,
		Nama:         input.Penduduk.Nama,
		JenisKelamin: input.Penduduk.JenisKelamin,
		TanggalLahir: tanggalLahir,
		KelurahanID:  input.Penduduk.KelurahanID,
		Alamat:       input.Penduduk.Alamat,
		NoHP:         input.Penduduk.NoHP,
		Kategori:     data.KategoriBUSUI,
	}

	data.ValidatePendudukInput(v, &data.PendudukInput{
		NIK:          penduduk.NIK,
		Nama:         penduduk.Nama,
		JenisKelamin: penduduk.JenisKelamin,
		TanggalLahir: input.Penduduk.TanggalLahir,
		KelurahanID:  penduduk.KelurahanID,
		Alamat:       penduduk.Alamat,
		NoHP:         penduduk.NoHP,
		Kategori:     penduduk.Kategori,
	})

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	tx, err := app.models.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("error begintx")
		app.serverErrorResponse(w, r, err)
		return
	}

	defer tx.Rollback()

	err = app.models.Penduduk.InsertTx(ctx, tx, penduduk)
	if err != nil {
		fmt.Println("error penduduk inserttx")
		switch {
		case errors.Is(err, data.ErrDuplicateNIK):
			v := validator.New()
			v.AddError("nik", "NIK sudah terdaftar")
			app.failedValidationResponse(w, r, v.Errors)

		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	busuiValidator := validator.New()

	tanggalPersalinan, err := time.Parse("2006-01-02", input.Busui.TanggalPersalinan)
	if err != nil {
		v.AddError("tanggal_lahir", "format tanggal harus YYYY-MM-DD")
		app.failedValidationResponse(w, r, busuiValidator.Errors)
		return
	}

	busui := &data.Busui{
		PendudukID:        penduduk.ID,
		PosyanduID:        posyandu.ID,
		AnakKe:            input.Busui.AnakKe,
		AsiEksklusif:      input.Busui.AsiEksklusif,
		TanggalPersalinan: tanggalPersalinan,
	}

	if data.ValidateBusui(busuiValidator, busui); !busuiValidator.Valid() {
		app.failedValidationResponse(w, r, busuiValidator.Errors)
		return
	}

	err = app.models.Busui.InsertTx(ctx, tx, busui)
	if err != nil {
		fmt.Println("error busui inserttx")
		app.serverErrorResponse(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error commit")
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posyandu/%d", posyandu.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"busui": busui}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createBalitaHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 5 {
		fmt.Println("error role id != 5")
		app.notPermittedResponse(w, r)
		return
	}

	posyandu, err := app.models.Posyandu.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get posyandu")
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if id != posyandu.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		// Data penduduk
		Penduduk struct {
			NIK          string `json:"nik"`
			Nama         string `json:"nama"`
			JenisKelamin string `json:"jenis_kelamin"`
			TanggalLahir string `json:"tanggal_lahir"`
			KelurahanID  int64  `json:"kelurahan_id"`
			Alamat       string `json:"alamat"`
			NoHP         string `json:"no_hp"`
		} `json:"penduduk"`

		// Data peserta didik
		Balita struct {
			IbuID        int64 `json:"ibu_id"`
			AnakKe       int8  `json:"anak_ke"`
			BeratLahir   int   `json:"berat_lahir"`
			PanjangLahir int   `json:"panjang_lahir"`
		} `json:"balita"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	tanggalLahir, err := time.Parse("2006-01-02", input.Penduduk.TanggalLahir)
	if err != nil {
		v.AddError("tanggal_lahir", "format tanggal harus YYYY-MM-DD")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	penduduk := &data.Penduduk{
		NIK:          input.Penduduk.NIK,
		Nama:         input.Penduduk.Nama,
		JenisKelamin: input.Penduduk.JenisKelamin,
		TanggalLahir: tanggalLahir,
		KelurahanID:  input.Penduduk.KelurahanID,
		Alamat:       input.Penduduk.Alamat,
		NoHP:         input.Penduduk.NoHP,
		Kategori:     data.KategoriBALITA,
	}

	data.ValidatePendudukInput(v, &data.PendudukInput{
		NIK:          penduduk.NIK,
		Nama:         penduduk.Nama,
		JenisKelamin: penduduk.JenisKelamin,
		TanggalLahir: input.Penduduk.TanggalLahir,
		KelurahanID:  penduduk.KelurahanID,
		Alamat:       penduduk.Alamat,
		NoHP:         penduduk.NoHP,
		Kategori:     penduduk.Kategori,
	})

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	tx, err := app.models.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("error begintx")
		app.serverErrorResponse(w, r, err)
		return
	}

	defer tx.Rollback()

	err = app.models.Penduduk.InsertTx(ctx, tx, penduduk)
	if err != nil {
		fmt.Println("error penduduk inserttx")
		switch {
		case errors.Is(err, data.ErrDuplicateNIK):
			v := validator.New()
			v.AddError("nik", "NIK sudah terdaftar")
			app.failedValidationResponse(w, r, v.Errors)

		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	balita := &data.Balita{
		PendudukID:   penduduk.ID,
		PosyanduID:   posyandu.ID,
		IbuID:        input.Balita.IbuID,
		AnakKe:       input.Balita.AnakKe,
		BeratLahir:   input.Balita.BeratLahir,
		PanjangLahir: input.Balita.PanjangLahir,
	}

	// Initialize a new Validator instance.
	balitaValidator := validator.New()

	if data.ValidateBalita(balitaValidator, balita); !balitaValidator.Valid() {
		app.failedValidationResponse(w, r, balitaValidator.Errors)
		return
	}

	err = app.models.Balita.InsertTx(ctx, tx, balita)
	if err != nil {
		fmt.Println("error balita inserttx")
		app.serverErrorResponse(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error commit")
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posyandu/%d", posyandu.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"balita": balita}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listBalitaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	var input struct {
		Nama string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	balita, metadata, err := app.models.Balita.GetAll(posyandu.ID, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"balita": balita, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listBumilHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	var input struct {
		Nama string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	bumil, metadata, err := app.models.Bumil.GetAll(posyandu.ID, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"bumil": bumil, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listBusuiHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	var input struct {
		Nama string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	busui, metadata, err := app.models.Busui.GetAll(posyandu.ID, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"busui": busui, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createBumilHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 5 {
		fmt.Println("error role id != 5")
		app.notPermittedResponse(w, r)
		return
	}

	posyandu, err := app.models.Posyandu.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get posyandu")
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if id != posyandu.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		// Data penduduk
		Penduduk struct {
			NIK          string `json:"nik"`
			Nama         string `json:"nama"`
			JenisKelamin string `json:"jenis_kelamin"`
			TanggalLahir string `json:"tanggal_lahir"`
			KelurahanID  int64  `json:"kelurahan_id"`
			Alamat       string `json:"alamat"`
			NoHP         string `json:"no_hp"`
		} `json:"penduduk"`

		// Data peserta didik
		Bumil struct {
			HPHT time.Time `json:"hpht"`
			HPL  time.Time `json:"hpl"`

			Gravida int `json:"gravida"`
			Para    int `json:"para"`
			Abortus int `json:"abortus"`
		} `json:"bumil"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	tanggalLahir, err := time.Parse("2006-01-02", input.Penduduk.TanggalLahir)
	if err != nil {
		v.AddError("tanggal_lahir", "format tanggal harus YYYY-MM-DD")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	penduduk := &data.Penduduk{
		NIK:          input.Penduduk.NIK,
		Nama:         input.Penduduk.Nama,
		JenisKelamin: input.Penduduk.JenisKelamin,
		TanggalLahir: tanggalLahir,
		KelurahanID:  input.Penduduk.KelurahanID,
		Alamat:       input.Penduduk.Alamat,
		NoHP:         input.Penduduk.NoHP,
		Kategori:     data.KategoriBUMIL,
	}

	data.ValidatePendudukInput(v, &data.PendudukInput{
		NIK:          penduduk.NIK,
		Nama:         penduduk.Nama,
		JenisKelamin: penduduk.JenisKelamin,
		TanggalLahir: input.Penduduk.TanggalLahir,
		KelurahanID:  penduduk.KelurahanID,
		Alamat:       penduduk.Alamat,
		NoHP:         penduduk.NoHP,
		Kategori:     penduduk.Kategori,
	})

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	tx, err := app.models.DB.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("error begintx")
		app.serverErrorResponse(w, r, err)
		return
	}

	defer tx.Rollback()

	err = app.models.Penduduk.InsertTx(ctx, tx, penduduk)
	if err != nil {
		fmt.Println("error penduduk inserttx")
		switch {
		case errors.Is(err, data.ErrDuplicateNIK):
			v := validator.New()
			v.AddError("nik", "NIK sudah terdaftar")
			app.failedValidationResponse(w, r, v.Errors)

		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	bumil := &data.Bumil{
		PendudukID: penduduk.ID,
		PosyanduID: posyandu.ID,
		HPHT:       input.Bumil.HPHT,
		HPL:        input.Bumil.HPL,
		Gravida:    input.Bumil.Gravida,
		Para:       input.Bumil.Para,
		Abortus:    input.Bumil.Abortus,
	}

	// Initialize a new Validator instance.
	bumilValidator := validator.New()

	if data.ValidateBumil(bumilValidator, bumil); !bumilValidator.Valid() {
		app.failedValidationResponse(w, r, bumilValidator.Errors)
		return
	}

	err = app.models.Bumil.InsertTx(ctx, tx, bumil)
	if err != nil {
		fmt.Println("error bumil inserttx")
		app.serverErrorResponse(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error commit")
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posyandu/%d", posyandu.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"bumil": bumil}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listIbuHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	var input struct {
		Nama string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	ibu, metadata, err := app.models.Penduduk.GetAllIbu(posyandu.ID, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"ibu": ibu, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createPengukuranBalitaHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 5 {
		fmt.Println("error role id != 5")
		app.notPermittedResponse(w, r)
		return
	}

	posyandu, err := app.models.Posyandu.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get posyandu")
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if id != posyandu.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		PendudukID    int64   `json:"penduduk_id"`
		Tanggal       string  `json:"tanggal"`
		UmurBulan     int16   `json:"umur_bulan"`
		BeratBadan    float64 `json:"berat_badan"`
		TinggiBadan   float64 `json:"tinggi_badan"`
		LingkarKepala float64 `json:"lingkar_kepala"`
		Lila          float64 `json:"lila"`
		Catatan       string  `json:"catatan"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	err = app.models.Balita.ValidateBalitaInclude(posyandu.ID, input.PendudukID)
	if err != nil {
		v.AddError("penduduk_id", "penduduk yang dipilih tidak terdaftar sebagai balita di posyandu ini")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	tanggal, err := time.Parse("2006-01-02", input.Tanggal)
	if err != nil {
		v.AddError("tanggal", "format tanggal harus YYYY-MM-DD")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengukuran_balita := &data.PengukuranBalita{
		PendudukID:    input.PendudukID,
		Tanggal:       tanggal,
		UmurBulan:     input.UmurBulan,
		BeratBadan:    input.BeratBadan,
		TinggiBadan:   input.TinggiBadan,
		LingkarKepala: input.LingkarKepala,
		Lila:          input.Lila,
		Catatan:       input.Catatan,
	}

	if data.ValidatePengukuranBalita(v, pengukuran_balita); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.PengukuranBalita.Insert(pengukuran_balita)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posyandu/%d", posyandu.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"pengukuran_balita": pengukuran_balita}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createPengukuranBusuiHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 5 {
		fmt.Println("error role id != 5")
		app.notPermittedResponse(w, r)
		return
	}

	posyandu, err := app.models.Posyandu.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get posyandu")
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if id != posyandu.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		PendudukID  int64   `json:"penduduk_id"`
		Tanggal     string  `json:"tanggal"`
		BeratBadan  float64 `json:"berat_badan"`
		TinggiBadan float64 `json:"tinggi_badan"`
		Hemoglobin  float64 `json:"hemoglobin"`
		Lila        float64 `json:"lila"`
		Catatan     string  `json:"catatan"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	err = app.models.Busui.ValidateBusuiInclude(posyandu.ID, input.PendudukID)
	if err != nil {
		v.AddError("penduduk_id", "penduduk yang dipilih tidak terdaftar sebagai busui di posyandu ini")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	tanggal, err := time.Parse("2006-01-02", input.Tanggal)
	if err != nil {
		v.AddError("tanggal", "format tanggal harus YYYY-MM-DD")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengukuran_busui := &data.PengukuranBusui{
		PendudukID:  input.PendudukID,
		Tanggal:     tanggal,
		BeratBadan:  input.BeratBadan,
		TinggiBadan: input.TinggiBadan,
		Hemoglobin:  input.Hemoglobin,
		Lila:        input.Lila,
		Catatan:     input.Catatan,
	}

	if data.ValidatePengukuranBusui(v, pengukuran_busui); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.PengukuranBusui.Insert(pengukuran_busui)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posyandu/%d", posyandu.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"pengukuran_busui": pengukuran_busui}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createPengukuranBumilHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 5 {
		fmt.Println("error role id != 5")
		app.notPermittedResponse(w, r)
		return
	}

	posyandu, err := app.models.Posyandu.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get posyandu")
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	if id != posyandu.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		PendudukID            int64   `json:"penduduk_id"`
		Tanggal               string  `json:"tanggal"`
		BeratBadan            float64 `json:"berat_badan"`
		TinggiBadan           float64 `json:"tinggi_badan"`
		Hemoglobin            float64 `json:"hemoglobin"`
		Lila                  float64 `json:"lila"`
		UsiaKehamilanMinggu   int     `json:"usia_kehamilan_minggu"`
		TekananDarahSistolik  int     `json:"tekanan_darah_sistolik"`
		TekananDarahDiastolik int     `json:"tekanan_darah_diastolik"`
		Catatan               string  `json:"catatan"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	err = app.models.Bumil.ValidateBumilInclude(posyandu.ID, input.PendudukID)
	if err != nil {
		v.AddError("penduduk_id", "penduduk yang dipilih tidak terdaftar sebagai bumil di posyandu ini")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	tanggal, err := time.Parse("2006-01-02", input.Tanggal)
	if err != nil {
		v.AddError("tanggal", "format tanggal harus YYYY-MM-DD")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengukuran_bumil := &data.PengukuranBumil{
		PendudukID:            input.PendudukID,
		Tanggal:               tanggal,
		BeratBadan:            input.BeratBadan,
		TinggiBadan:           input.TinggiBadan,
		Hemoglobin:            input.Hemoglobin,
		Lila:                  input.Lila,
		UsiaKehamilanMinggu:   input.UsiaKehamilanMinggu,
		TekananDarahSistolik:  input.TekananDarahSistolik,
		TekananDarahDiastolik: input.TekananDarahDiastolik,
		Catatan:               input.Catatan,
	}

	if data.ValidatePengukuranBumil(v, pengukuran_bumil); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.PengukuranBumil.Insert(pengukuran_bumil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/posyandu/%d", posyandu.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"pengukuran_bumil": pengukuran_bumil}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPengukuranBalitaHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	var input struct {
		Nama  string
		Bulan string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Bulan = app.readString(qs, "bulan", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
		"ibu_nama",
		"-ibu_nama",
		"tinggi_badan",
		"-tinggi_badan",
		"berat_badan",
		"-berat_badan",
		"umur_bulan",
		"-umur_bulan",
		"lingkar_kepala",
		"-lingkar_kepala",
		"lila",
		"-lila",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengukuran_balita, metadata, err := app.models.PengukuranBalita.GetAll(posyandu.ID, input.Bulan, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"pengukuran_balita": pengukuran_balita, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPengukuranBumilHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	var input struct {
		Nama  string
		Bulan string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Bulan = app.readString(qs, "bulan", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
		"berat_badan",
		"-berat_badan",
		"tinggi_badan",
		"-tinggi_badan",
		"hemoglobin",
		"-hemoglobin",
		"lila",
		"-lila",
		"usia_kehamilan_minggu",
		"-usia_kehamilan_minggu",
		"tekanan_darah_sistolik",
		"-tekanan_darah_sistolik",
		"tekanan_darah_diastolik",
		"-tekanan_darah_diastolik",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengukuran_bumil, metadata, err := app.models.PengukuranBumil.GetAll(posyandu.ID, input.Bulan, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"pengukuran_bumil": pengukuran_bumil, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPengukuranBusuiHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	var input struct {
		Nama  string
		Bulan string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	input.Nama = app.readString(qs, "nama", "")
	input.Bulan = app.readString(qs, "bulan", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)

	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{
		"id",
		"-id",
		"nama",
		"-nama",
		"berat_badan",
		"-berat_badan",
		"tinggi_badan",
		"-tinggi_badan",
		"hemoglobin",
		"-hemoglobin",
		"lila",
		"-lila",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengukuran_busui, metadata, err := app.models.PengukuranBusui.GetAll(posyandu.ID, input.Bulan, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"pengukuran_busui": pengukuran_busui, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
