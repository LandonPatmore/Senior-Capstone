package stations

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/landonp1203/goUtils/networking"
	"net/http"
	"tidewatchBackend/utils"
)

type Stations struct {
	Count    int         `json:"count"`
	Units    interface{} `json:"units"`
	Stations []struct {
		Tidal      bool   `json:"tidal"`
		Greatlakes bool   `json:"greatlakes"`
		Shefcode   string `json:"shefcode"`
		Details    struct {
			Self string `json:"self"`
		} `json:"details"`
		Sensors struct {
			Self string `json:"self"`
		} `json:"sensors"`
		Floodlevels struct {
			Self string `json:"self"`
		} `json:"floodlevels"`
		Datums struct {
			Self string `json:"self"`
		} `json:"datums"`
		Supersededdatums struct {
			Self string `json:"self"`
		} `json:"supersededdatums"`
		HarmonicConstituents struct {
			Self string `json:"self"`
		} `json:"harmonicConstituents"`
		Benchmarks struct {
			Self string `json:"self"`
		} `json:"benchmarks"`
		TidePredOffsets struct {
			Self string `json:"self"`
		} `json:"tidePredOffsets"`
		State        string `json:"state"`
		Timezone     string `json:"timezone"`
		Timezonecorr int    `json:"timezonecorr"`
		Observedst   bool   `json:"observedst"`
		Stormsurge   bool   `json:"stormsurge"`
		Nearby       struct {
			Self string `json:"self"`
		} `json:"nearby"`
		ID           string      `json:"id"`
		Name         string      `json:"name"`
		Lat          float64     `json:"lat"`
		Lng          float64     `json:"lng"`
		Affiliations string      `json:"affiliations"`
		Portscode    interface{} `json:"portscode"`
		Products     struct {
			Self string `json:"self"`
		} `json:"products"`
		Disclaimers struct {
			Self string `json:"self"`
		} `json:"disclaimers"`
		Notices struct {
			Self string `json:"self"`
		} `json:"notices"`
		Self     string `json:"self"`
		Expand   string `json:"expand"`
		TideType string `json:"tideType"`
	} `json:"stations"`
	Self interface{} `json:"self"`
}

// TODO: Need to update this to work with the DB so we are not always hitting NOAA for no reason if we already have this data saved in a DB

// Generates the routes
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/stations", getAllStations)
	router.Get("/stations/{type}", getAllStationsByType)
	router.Get("/station/{id}", getSingleStation)

	return router
}

// GET: Retrieves all stations from NOAA
func getAllStations(w http.ResponseWriter, r *http.Request) {
	stations := Stations{}

	body, _ := networking.Get("https://tidesandcurrents.noaa.gov/mdapi/v1.0/webapi/stations.json")
	utils.ErrorCheck(json.Unmarshal(body, &stations))

	render.JSON(w, r, stations)
}

// GET: Retrieves all stations from NOAA, based on type
func getAllStationsByType(w http.ResponseWriter, r *http.Request) {
	t := chi.URLParam(r, "type")

	stations := Stations{}
	body, _ := networking.Get("https://tidesandcurrents.noaa.gov/mdapi/v1.0/webapi/stations.json?type=" + t)
	utils.ErrorCheck(json.Unmarshal(body, &stations))

	render.JSON(w, r, stations)
}

// GET: Retrieves a single station from NOAA
func getSingleStation(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	stations := Stations{}
	body, _ := networking.Get("https://tidesandcurrents.noaa.gov/mdapi/v1.0/webapi/stations/" + id + ".json")
	utils.ErrorCheck(json.Unmarshal(body, &stations))

	render.JSON(w, r, stations)
}
