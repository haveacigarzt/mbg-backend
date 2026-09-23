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

// func (app *application) updatePesertaDidikHandler(w http.ResponseWriter, r *http.Request) {

// 	nisn, err := app.readStringParam(r, "nisn")
// 	if err != nil || len(nisn) != 10 {
// 		app.notFoundResponse(w, r)
// 		return
// 	}

// 	user := app.contextGetUser(r)
// 	if user.RoleID != 6 {
// 		app.notPermittedResponse(w, r)
// 		return
// 	}

// 	sekolah, err := app.models.Sekolah.GetByUserID(user.ID)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			app.notFoundResponse(w, r)
// 		default:
// 			app.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}

// 	peserta_didik, err := app.models.PesertaDidik.GetByNISN(nisn)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			app.notFoundResponse(w, r)
// 		default:
// 			app.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}

// 	if sekolah.ID != peserta_didik.SekolahID {
// 		app.notPermittedResponse(w, r)
// 		return
// 	}

// 	var input struct {
// 		// Data penduduk
// 		Penduduk struct {
// 			NIK          *string `json:"nik"`
// 			Nama         *string `json:"nama"`
// 			JenisKelamin *string `json:"jenis_kelamin"`
// 			TanggalLahir *string `json:"tanggal_lahir"`
// 			KelurahanID  *int64  `json:"kelurahan_id"`
// 			Alamat       *string `json:"alamat"`
// 			NoHP         *string `json:"no_hp"`
// 		} `json:"penduduk"`

// 		// Data peserta didik
// 		PesertaDidik struct {
// 			NISN   *string `json:"nisn"`
// 			Kelas  *string `json:"kelas"`
// 			Rombel *string `json:"rombel"`
// 		} `json:"peserta_didik"`
// 	}

// 	// Read the JSON request body data into the input struct.
// 	err = app.readJSON(w, r, &input)
// 	if err != nil {
// 		app.badRequestResponse(w, r, err)
// 		return
// 	}

// 	if input.Penduduk.Nama != nil {
// 		peserta_didik. = *input.Nama
// 	}

// 	if input.Alamat != nil {
// 		posyandu.Alamat = *input.Alamat
// 	}

// 	if input.Kecamatan_ID != nil {
// 		posyandu.Kecamatan_ID = *input.Kecamatan_ID
// 	}

// 	if input.Kelurahan_ID != nil {
// 		posyandu.Kelurahan_ID = *input.Kelurahan_ID
// 	}

// 	if input.Latitude != nil {
// 		posyandu.Latitude = *input.Latitude
// 	}

// 	if input.Longitude != nil {
// 		posyandu.Longitude = *input.Longitude
// 	}

// 	peserta_didik, err := app.models.PesertaDidik.GetByNISN(nisn)
// 	if err != nil {
// 		switch {
// 		case errors.Is(err, data.ErrRecordNotFound):
// 			err = app.writeJSON(w, http.StatusOK, envelope{"peserta_didik": nil}, nil)
// 			if err != nil {
// 				app.serverErrorResponse(w, r, err)
// 			}
// 		default:
// 			app.serverErrorResponse(w, r, err)
// 		}
// 		return
// 	}

// 	err = app.writeJSON(w, http.StatusOK, envelope{"peserta_didik": peserta_didik}, nil)
// 	if err != nil {
// 		app.serverErrorResponse(w, r, err)
// 	}
// }
