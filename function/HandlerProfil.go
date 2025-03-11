package function

import (
	"net/http"
	"strconv"
)

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		renderTemplate(w, "templates/404.html", Error{Errr: 400, Kalma: http.StatusText(400)}, 400)
		return
	}

	artistID, err := strconv.Atoi(idStr)
	if err != nil {
		renderTemplate(w, "templates/404.html", Error{Errr: 400, Kalma: http.StatusText(400)}, 400)
		return
	}

	artist, locations, dates, relation, err := fetchArtistDetails(artistID)
	if err != nil {
		renderTemplate(w, "templates/404.html", Error{Errr: 404, Kalma: http.StatusText(404)}, 404)
		return
	}

	renderTemplate(w, "templates/artist.html", struct {
		Artist    Artist
		Locations LocationResponse
		Dates     DateResponse
		Relation  RelationReponse
	}{
		Artist:    artist,
		Locations: locations,
		Dates:     dates,
		Relation:  relation,
	}, 200)
}
