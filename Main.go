package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"tidewatchBackend/stations"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api", stations.Routes())
	})

	return router
}

func main() {
	router := Routes()

	log.Fatal(http.ListenAndServe(":3000", router))
}