package stations

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strings"
	"tideWatchAPI/utils"
	"tideWatchAPI/utils/db"
	"tideWatchAPI/utils/structs"
)

var d = db.CreateConnection()

// GET: Retrieves a list of stations based on an entered ID
func GetStationsById(w http.ResponseWriter, r *http.Request) {
	t := strings.ToUpper(chi.URLParam(r, "type"))
	id := chi.URLParam(r, "id")

	switch t {
	case "W":
		var stations [] structs.WaterLevelStation
		db.GetAllStationsById(d, &stations, id)

		render.JSON(w, r, stations)
	case "C":
		var stations [] structs.CurrentStation
		db.GetAllStationsById(d, &stations, id)

		render.JSON(w, r, stations)
	case "CP":
		var stations [] structs.CurrentPredictionStation
		db.GetAllStationsById(d, &stations, id)

		render.JSON(w, r, stations)
	case "TP":
		var stations [] structs.TidePredictionStation
		db.GetAllStationsById(d, &stations, id)

		render.JSON(w, r, stations)
	default:
		utils.ErrorResponse(w, r, `The type "`+t+`" does not exist.`)
	}
}

// GET: Retrieves a list of stations by type
func GetStationsByType(w http.ResponseWriter, r *http.Request) {
	t := strings.ToUpper(chi.URLParam(r, "type"))

	switch t {
	case "W":
		var stations [] structs.WaterLevelStation
		db.GetAllStationsByType(d, &stations)

		render.JSON(w, r, stations)
	case "C":
		var stations [] structs.CurrentStation
		db.GetAllStationsByType(d, &stations)

		render.JSON(w, r, stations)
	case "CP":
		var stations [] structs.CurrentPredictionStation
		db.GetAllStationsByType(d, &stations)

		render.JSON(w, r, stations)
	case "TP":
		var stations [] structs.TidePredictionStation
		db.GetAllStationsByType(d, &stations)

		render.JSON(w, r, stations)
	default:
		utils.ErrorResponse(w, r, `The type "`+t+`" does not exist.`)
	}
}
