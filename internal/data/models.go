package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	DB                *sql.DB
	Sekolah           SekolahModel
	Tokens            TokenModel
	Users             UserModel
	Kecamatan         KecamatanModel
	Kelurahan         KelurahanModel
	Permissions       PermissionModel
	SPPG              SPPGModel
	Posyandu          PosyanduModel
	Drivers           DriverModel
	Pengiriman        PengirimanModel
	Tracking          TrackingModel
	AlokasiHarian     AlokasiHarianModel
	PengeluaranHarian PengeluaranHarianModel
	ProduksiHarian    ProduksiHarianModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		DB:                db,
		Sekolah:           SekolahModel{DB: db},
		Tokens:            TokenModel{DB: db},
		Users:             UserModel{DB: db},
		Kecamatan:         KecamatanModel{DB: db},
		Kelurahan:         KelurahanModel{DB: db},
		Permissions:       PermissionModel{DB: db},
		SPPG:              SPPGModel{DB: db},
		Posyandu:          PosyanduModel{DB: db},
		Drivers:           DriverModel{DB: db},
		Pengiriman:        PengirimanModel{DB: db},
		Tracking:          TrackingModel{DB: db},
		AlokasiHarian:     AlokasiHarianModel{DB: db},
		PengeluaranHarian: PengeluaranHarianModel{DB: db},
		ProduksiHarian:    ProduksiHarianModel{DB: db},
	}
}
