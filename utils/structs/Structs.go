package structs

type WOutter struct {
	Stations [] WaterLevelStation `json:"stations,omitempty"`
}

type COutter struct {
	Stations [] CurrentStation `json:"stations,omitempty"`
}

type CPOutter struct {
	Stations [] CurrentPredictionStation `json:"stations,omitempty"`
}

type TPOutter struct {
	Stations [] TidePredictionStation `json:"stations,omitempty"`
}

type WaterLevelStation struct {
	Tidal      bool   `json:"tidal,omitempty"`
	Greatlakes bool   `json:"greatlakes,omitempty"`
	Shefcode   string `json:"shefcode,omitempty"`
	Details    struct {
		Self string `json:"self,omitempty"`
	} `json:"details,omitempty"`
	Sensors struct {
		Self string `json:"self,omitempty"`
	} `json:"sensors,omitempty"`
	Floodlevels struct {
		Self string `json:"self,omitempty"`
	} `json:"floodlevels,omitempty"`
	Datums struct {
		Self string `json:"self,omitempty"`
	} `json:"datums,omitempty"`
	Supersededdatums struct {
		Self string `json:"self,omitempty"`
	} `json:"supersededdatums,omitempty"`
	HarmonicConstituents struct {
		Self string `json:"self,omitempty"`
	} `json:"harmonicConstituents,omitempty"`
	TidePredOffsets struct {
		Self string `json:"self,omitempty"`
	} `json:"tidePredOffsets,omitempty"`
	State        string  `json:"state,omitempty"`
	Timezone     string  `json:"timezone,omitempty"`
	Timezonecorr int     `json:"timezonecorr,omitempty"`
	Observedst   bool    `json:"observedst,omitempty"`
	Stormsurge   bool    `json:"stormsurge,omitempty"`
	ID           string  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Lat          float64 `json:"lat,omitempty"`
	Lng          float64 `json:"lng,omitempty"`
	Affiliations string  `json:"affiliations,omitempty"`
	Portscode    *string `json:"portscode,omitempty"`
	Self         string  `json:"self,omitempty"`
	TideType     string  `json:"tideType,omitempty"`
}

type CurrentStation struct {
	Project        string `json:"project,omitempty"`
	Deployed       string `json:"deployed,omitempty"`
	Retrieved      string `json:"retrieved,omitempty"`
	TimezoneOffset string `json:"timezone_offset,omitempty"`
	Observedst     bool   `json:"observedst,omitempty"`
	ProjectType    string `json:"project_type,omitempty"`
	Noaachart      int    `json:"noaachart,omitempty"`
	Deployments    struct {
		Self string `json:"self,omitempty"`
	} `json:"deployments,omitempty"`
	Bins struct {
		Self string `json:"self,omitempty"`
	} `json:"bins,omitempty"`
	HeightFromBottom float64 `json:"height_from_bottom,omitempty"`
	CenterBin1Dist   float64 `json:"center_bin_1_dist,omitempty"`
	ID               string  `json:"id,omitempty"`
	Name             string  `json:"name,omitempty"`
	Lat              float64 `json:"lat,omitempty"`
	Lng              float64 `json:"lng,omitempty"`
	Affiliations     string  `json:"affiliations,omitempty"`
	Portscode        string  `json:"portscode,omitempty"`
	Self             string  `json:"self,omitempty"`
	Expand           string  `json:"expand,omitempty"`
	TideType         string  `json:"tideType,omitempty"`
}

type CurrentPredictionStation struct {
	Currentpredictionoffsets struct {
		CPSelf string `json:"self,omitempty"`
	} `json:"currentpredictionoffsets,omitempty" gorm:"embedded"`
	Currbin      int      `json:"currbin,omitempty"`
	Type         string   `json:"type,omitempty"`
	Depth        *float64 `json:"depth,omitempty"`
	DepthType    string   `json:"depthType,omitempty"`
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Lat          float64  `json:"lat,omitempty"`
	Lng          float64  `json:"lng,omitempty"`
	Affiliations string   `json:"affiliations,omitempty"`
	Portscode    string   `json:"portscode,omitempty"`
	Self         string   `json:"self,omitempty"`
	TideType     string   `json:"tideType,omitempty"`
}

type TidePredictionStation struct {
	State           string `json:"state,omitempty"`
	Tidepredoffsets struct {
		Self string `json:"self,omitempty"`
	} `json:"tidepredoffsets,omitempty"`
	Type         string  `json:"type,omitempty"`
	Timemeridian int     `json:"timemeridian,omitempty"`
	ReferenceID  string  `json:"reference_id,omitempty"`
	Timezonecorr int     `json:"timezonecorr,omitempty"`
	ID           string  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Lat          float64 `json:"lat,omitempty"`
	Lng          float64 `json:"lng,omitempty"`
	Affiliations string  `json:"affiliations,omitempty"`
	Portscode    string  `json:"portscode,omitempty"`
	TideType     string  `json:"tideType,omitempty"`
}
