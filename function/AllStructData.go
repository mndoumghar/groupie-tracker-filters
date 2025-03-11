package function


type Artist struct {
	ID int `json:"id"`

	Name       string   `json:"name"`
	Image      string   `json:"image"`
	Year       int      `json:"creationDate"`
	FirstAlbum string   `json:"firstAlbum"`
	Members    []string `json:"members"`
}

type LocationResponse struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type DateResponse struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationReponse struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Error struct {
	Errr  int
	Kalma string
}