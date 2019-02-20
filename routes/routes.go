package routes

import (
	"github.com/go-chi/chi"
	"tidewatchBackend/stations"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/stations", stations.GetAllStations)
	router.Get("/stations/{type}", stations.GetAllStationsByType)
	router.Get("/station/{id}", stations.GetSingleStation)

	return router
}