package function

import (
	"net/http"
	"os"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderTemplate(w, "templates/404.html", Error{Errr: 404, Kalma: http.StatusText(404)}, 404)
		return
	}

	artists, err := fetchArtists()
	if err != nil {
		renderTemplate(w, "templates/404.html", Error{Errr: 500, Kalma: http.StatusText(500)}, 500)
		return
	}

	location, err := fetchallLocations()
	if err != nil {
		renderTemplate(w, "templates/404.html", Error{Errr: 500, Kalma: http.StatusText(500)}, 500)
		return
	}

	renderTemplate(w, "templates/index.html", struct {
		Artists  []Artist
		Location LocationResponse
	}{
		Artists:  artists,
		Location: location,
	}, 200)
}

func Static(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		renderTemplate(w, "templates/404.html", Error{Errr: 500, Kalma: http.StatusText(500)}, 500)
		return
	}
	filePath := strings.TrimPrefix(r.URL.Path, "/static/")
	fullPath := "static/" + filePath
	info, err := os.Stat(fullPath)
	if err != nil {
		renderTemplate(w, "templates/404.html", Error{Errr: 500, Kalma: http.StatusText(500)}, 500)
		return
	}
	if info.IsDir() {
		renderTemplate(w, "templates/404.html", Error{Errr: 500, Kalma: http.StatusText(500)}, 500)
		return
	}
	fs := http.Dir("static")
	http.StripPrefix("/static/", http.FileServer(fs)).ServeHTTP(w, r)
}
