package function

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchLocations(artistID int) (LocationResponse, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", artistID)
	resp, err := http.Get(url)
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return LocationResponse{},err
	}

	var locationResponse LocationResponse
	err = json.NewDecoder(resp.Body).Decode(&locationResponse)
	if err != nil {
		return LocationResponse{}, err
	}
	return locationResponse, nil
}

func fetchallLocations() (LocationResponse, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return LocationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return LocationResponse{}, err
	}

	var locationResponse LocationResponse
	err = json.NewDecoder(resp.Body).Decode(&locationResponse)
	if err != nil {
		return LocationResponse{}, err
	}
	return locationResponse, nil
}
