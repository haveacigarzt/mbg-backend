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

func (app *application) createSekolahHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 3 {
		app.notPermittedResponse(w, r)
		return
	}

	sppg, err := app.models.SPPG.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get sppg")
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
		Kategori     string  `json:"kategori"`
		JumlahSiswa  int     `json:"jumlah_siswa"`
		Kecamatan_ID int64   `json:"kecamatan_id"`
		Kelurahan_ID int64   `json:"kelurahan_id"`
		Latitude     float64 `json:"latitude"`
		Longitude    float64 `json:"longitude"`
		User         struct {
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
	sekolah := &data.Sekolah{
		Nama:         input.Nama,
		Alamat:       input.Alamat,
		Kategori:     input.Kategori,
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

	user := &data.User{
		Name:      input.User.Name,
		Email:     input.User.Email,
		RoleID:    6,
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

	sekolah.UserID = &user.ID

	err = app.models.Sekolah.InsertTx(ctx, tx, sekolah)
	if err != nil {
		fmt.Println("error sekolah inserttx")
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
		Kategori     *string  `json:"Kategori"`
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

	if input.Kategori != nil {
		sekolah.Kategori = *input.Kategori
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

func (app *application) createPesertaDidikHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 6 {
		fmt.Println("error role id != 6")
		app.notPermittedResponse(w, r)
		return
	}

	sekolah, err := app.models.Sekolah.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get sekolah")
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

	if id != sekolah.ID {
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
		PesertaDidik struct {
			NISN   string `json:"nisn"`
			Kelas  string `json:"kelas"`
			Rombel string `json:"rombel"`
		} `json:"peserta_didik"`
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
		Kategori:     data.KategoriPesertaDidik,
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

	pesertaDidik := &data.PesertaDidik{
		SekolahID:  sekolah.ID,
		PendudukID: penduduk.ID,
		NISN:       input.PesertaDidik.NISN,
		Kelas:      input.PesertaDidik.Kelas,
		Rombel:     input.PesertaDidik.Rombel,
	}

	// Initialize a new Validator instance.
	pesertaDidikValidator := validator.New()

	if data.ValidatePesertaDidik(pesertaDidikValidator, pesertaDidik); !pesertaDidikValidator.Valid() {
		app.failedValidationResponse(w, r, pesertaDidikValidator.Errors)
		return
	}

	err = app.models.PesertaDidik.InsertTx(ctx, tx, pesertaDidik)
	if err != nil {
		fmt.Println("error peserta didik inserttx")
		switch {
		case errors.Is(err, data.ErrDuplicateNISN):
			v := validator.New()
			v.AddError("nisn", "NISN sudah terdaftar")
			app.failedValidationResponse(w, r, v.Errors)

		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("error commit")
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/sekolah/%d", sekolah.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"peserta_didik": pesertaDidik}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPesertaDidikHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	peserta_didik, metadata, err := app.models.PesertaDidik.GetAll(sekolah.ID, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"peserta_didik": peserta_didik, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createPengukuranPesertaDidikHandler(w http.ResponseWriter, r *http.Request) {
	currUser := app.contextGetUser(r)
	if currUser.RoleID != 6 {
		fmt.Println("error role id != 6")
		app.notPermittedResponse(w, r)
		return
	}

	sekolah, err := app.models.Sekolah.GetByUserID(currUser.ID)
	if err != nil {
		fmt.Println("error get sekolah")
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

	if id != sekolah.ID {
		app.notPermittedResponse(w, r)
		return
	}

	var input struct {
		PendudukID  int64   `json:"penduduk_id"`
		Tanggal     string  `json:"tanggal"`
		UmurBulan   int16   `json:"umur_bulan"`
		BeratBadan  float64 `json:"berat_badan"`
		TinggiBadan float64 `json:"tinggi_badan"`
		Catatan     string  `json:"catatan"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	err = app.models.PesertaDidik.ValidatePesertaDidikInclude(sekolah.ID, input.PendudukID)
	if err != nil {
		v.AddError("penduduk_id", "penduduk yang dipilih tidak terdaftar sebagai peserta didik di sekolah ini")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	tanggal, err := time.Parse("2006-01-02", input.Tanggal)
	if err != nil {
		v.AddError("tanggal", "format tanggal harus YYYY-MM-DD")
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengukuran_anak := &data.PengukuranAnak{
		PendudukID:  input.PendudukID,
		Tanggal:     tanggal,
		UmurBulan:   input.UmurBulan,
		BeratBadan:  input.BeratBadan,
		TinggiBadan: input.TinggiBadan,
		Catatan:     input.Catatan,
	}

	if data.ValidatePengukuranAnak(v, pengukuran_anak); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.PengukuranAnak.Insert(pengukuran_anak)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/sekolah/%d", sekolah.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"pengukuran_anak": pengukuran_anak}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listPengukuranPesertaDidikHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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
		"tinggi_badan",
		"-tinggi_badan",
		"berat_badan",
		"-berat_badan",
		"umur_bulan",
		"-umur_bulan",
		"kelas",
		"-kelas",
		"rombel",
		"-rombel",
	}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pengukuran_peserta_didik, metadata, err := app.models.PengukuranAnak.GetAllPSD(sekolah.ID, input.Bulan, input.Nama, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Send a JSON response containing the movie data.
	err = app.writeJSON(w, http.StatusOK, envelope{"pengukuran_peserta_didik": pengukuran_peserta_didik, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
