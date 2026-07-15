package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/kecamatan", app.kecamatanHandler)
	router.HandlerFunc(http.MethodGet, "/v1/kelurahan/:id", app.kelurahanHandler)

	// SPPG routes
	router.HandlerFunc(http.MethodGet, "/v1/sppg", app.listSPPGHandler)
	router.HandlerFunc(http.MethodPost, "/v1/sppg", app.requirePermission("akunsppg:write", app.createSPPGHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sppg/:id", app.requirePermission("sppg:read", app.getSPPGHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/sppg/:id", app.requirePermission("sppg:write", app.updateSPPGHandler))
	router.HandlerFunc(http.MethodPost, "/v1/sppg/:id/alokasiharian", app.requirePermission("sppg:write", app.createSPPGAlokasiHarianHandler))
	router.HandlerFunc(http.MethodPost, "/v1/sppg/:id/pengeluaranharian", app.requirePermission("sppg:write", app.createSPPGPengeluaranHarianHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sppg/:id/alokasiharian", app.requirePermission("sppg:write", app.getSPPGAlokasiHarianHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sppg/:id/pengeluaranharian", app.requirePermission("sppg:write", app.listSPPGPengeluaranHarianHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/sppg/:id/pengeluaranharian/:pengeluaran_id", app.requirePermission("sppg:write", app.deleteSPPGPengeluaranHarianHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sppg/:id/produksiharian", app.requirePermission("sppg:write", app.getSPPGProduksiHarianHandler))
	router.HandlerFunc(http.MethodPost, "/v1/sppg/:id/produksiharian", app.requirePermission("sppg:write", app.createSPPGProduksiHarianHandler))

	// Sekolah routes
	router.HandlerFunc(http.MethodGet, "/v1/sekolah", app.listSekolahHandler)
	router.HandlerFunc(http.MethodPost, "/v1/sekolah", app.requirePermission("sekolah:write", app.createSekolahHandler))
	router.HandlerFunc(http.MethodGet, "/v1/sekolah/:id", app.requirePermission("sekolah:read", app.getSekolahHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/sekolah/:id", app.requirePermission("sekolah:write", app.updateSekolahHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/sekolah/:id", app.requirePermission("sekolah:write", app.deleteSekolahHandler))

	// Posyandu routes
	router.HandlerFunc(http.MethodGet, "/v1/posyandu", app.listPosyanduHandler)
	router.HandlerFunc(http.MethodPost, "/v1/posyandu", app.requirePermission("posyandu:write", app.createPosyanduHandler))
	router.HandlerFunc(http.MethodGet, "/v1/posyandu/:id", app.requirePermission("posyandu:read", app.getPosyanduHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/posyandu/:id", app.requirePermission("posyandu:write", app.updatePosyanduHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/posyandu/:id", app.requirePermission("posyandu:write", app.deletePosyanduHandler))

	// Driver routes
	router.HandlerFunc(http.MethodGet, "/v1/drivers", app.requirePermission("driver:read", app.listDriverHandler))
	router.HandlerFunc(http.MethodPost, "/v1/drivers", app.requirePermission("driver:write", app.createDriverHandler))
	router.HandlerFunc(http.MethodGet, "/v1/driver/current", app.requirePermission("tracking:write", app.getDriverMeHandler))
	router.HandlerFunc(http.MethodGet, "/v1/drivers/:id", app.requirePermission("driver:read", app.getDriverHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/drivers/:id", app.requirePermission("driver:write", app.updateDriverHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/drivers/:id", app.requirePermission("driver:write", app.deleteDriverHandler))

	// Pengiriman routes
	router.HandlerFunc(http.MethodGet, "/v1/pengiriman", app.requirePermission("pengiriman:read", app.listPengirimanHandler))
	router.HandlerFunc(http.MethodPost, "/v1/pengiriman", app.requirePermission("pengiriman:write", app.createPengirimanHandler))
	router.HandlerFunc(http.MethodGet, "/v1/pengiriman-aktif/driver", app.requirePermission("tracking:write", app.getPengirimanAktifByDriverHandler))
	router.HandlerFunc(http.MethodGet, "/v1/pengiriman/:id", app.requirePermission("pengiriman:read", app.getPengirimanHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/pengiriman/:id", app.requirePermission("pengiriman:write", app.updatePengirimanHandler))
	// router.HandlerFunc(http.MethodDelete, "/v1/pengiriman/:id", app.requirePermission("pengiriman:write", app.deletePengirimanHandler))

	// TRACKING routes
	router.HandlerFunc(http.MethodGet, "/v1/tracking", app.requirePermission("tracking:read", app.getTrackingHandler))
	router.HandlerFunc(http.MethodPost, "/v1/pengiriman/:id/tracking", app.requirePermission("tracking:write", app.createTrackingHandler))
	// router.HandlerFunc(http.MethodPatch, "/v1/tracking/:id", app.requirePermission("tracking:write", app.updateTrackingHandler))

	// Admin Routes
	router.HandlerFunc(http.MethodGet, "/v1/admin/akun", app.requirePermission("akunsppg:write", app.getAkunHandler))
	router.HandlerFunc(http.MethodGet, "/v1/admin/akun/summary", app.requirePermission("akunsppg:write", app.getAkunSummaryHandler))
	router.HandlerFunc(http.MethodPost, "/v1/admin/akun", app.requirePermission("akunsppg:write", app.createAkunHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/admin/akun", app.requirePermission("akunsppg:write", app.updateAkunHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/admin/akun", app.requirePermission("akunsppg:write", app.deleteAkunHandler))
	router.HandlerFunc(http.MethodPost, "/v1/admin/invitation/new", app.requirePermission("akunsppg:write", app.createSPPGInvitationHandler))
	router.HandlerFunc(http.MethodGet, "/v1/admin/invitation", app.requirePermission("akunsppg:write", app.listSPPGInvitationHandler))
	router.HandlerFunc(http.MethodGet, "/v1/invitation/:token", app.getSPPGInvitationHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/invitation/:token", app.deleteSPPGInvitationHandler)
	router.HandlerFunc(http.MethodPost, "/v1/register/:token", app.createSPPGByInvitationHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/admin/akun", app.requirePermission("akunsppg:write", app.geAkunHandler))
	// router.HandlerFunc(http.MethodGet, "/v1/admin/akun", app.requirePermission("akunsppg:write", app.geAkunHandler))
	// router.HandlerFunc(http.MethodGet, "/v1/admin/akun", app.requirePermission("akunsppg:write", app.geAkunHandler))

	// Keuangan routes
	router.HandlerFunc(http.MethodGet, "/v1/keuanganharian", app.listKeuanganHandler)
	router.HandlerFunc(http.MethodGet, "/v1/produksiharian", app.listProduksiHandler)

	// Pedagang Lokal Routes
	router.HandlerFunc(http.MethodGet, "/v1/pedaganglokal", app.listPedagangLokalHandler)
	router.HandlerFunc(http.MethodPost, "/v1/pedaganglokal", app.requirePermission("sekolah:write", app.createPedagangLokalHandler))
	router.HandlerFunc(http.MethodPatch, "/v1/pedaganglokal/:id", app.requirePermission("sekolah:write", app.updatePedagangLokalHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/pedaganglokal/:id", app.requirePermission("sekolah:write", app.deletePedagangLokalHandler))

	// Auth
	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)
	router.HandlerFunc(http.MethodGet, "/v1/users/authme", app.requireAuthenticatedUser(app.authUserHandler))
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)
	// router.HandlerFunc(http.MethodDelete, "/v1/tokens/authentication", app.requireAuthenticatedUser(app.deleteAuthenticationTokenHandler))
	router.HandlerFunc(http.MethodDelete, "/v1/tokens/authentication", app.requireAuthenticatedUser(app.deleteCurrentAuthenticationTokenHandler))

	// Register a new GET /debug/vars endpoint pointing to the expvar handler.
	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	// WebSocket endpoint
	router.HandlerFunc(http.MethodGet, "/ws/open", app.openWSHandler)
	router.HandlerFunc(http.MethodGet, "/ws/sppg/:id", app.sppgWSHandler)

	return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
