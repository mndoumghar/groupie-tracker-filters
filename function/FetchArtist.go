package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func fetchArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Error: Received non-200 status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch data from API, status code: %d", resp.StatusCode)
	}

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func fetchArtistDetails(artistID int) (Artist, LocationResponse, DateResponse, RelationReponse, error) {
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", artistID)
	resp, err := http.Get(artistURL)
	if err != nil {
		return Artist{}, LocationResponse{}, DateResponse{}, RelationReponse{}, err
	}
	defer resp.Body.Close()



	var artist Artist
	err = json.NewDecoder(resp.Body).Decode(&artist)
	if err != nil {
		return Artist{}, LocationResponse{}, DateResponse{}, RelationReponse{}, err
	}

	locations, err := fetchLocations(artist.ID)
	if err != nil {
		return Artist{}, LocationResponse{}, DateResponse{}, RelationReponse{}, err
	}

	dates, err := fetchDates(artist.ID)
	if err != nil {
		return Artist{}, LocationResponse{}, DateResponse{}, RelationReponse{}, err
	}

	relation, err := fetchRelations(artist.ID)
	if err != nil {
		return Artist{}, LocationResponse{}, DateResponse{}, RelationReponse{}, err
	}

	return artist, locations, dates, relation, nil
}
