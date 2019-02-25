package routes

import (
	"github.com/go-chi/chi"
	"tideWatchAPI/stations"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/stations/{type}", stations.GetStationsByType)
	router.Get("/stations/{type}/{id}", stations.GetStationsById)

	return router
}
