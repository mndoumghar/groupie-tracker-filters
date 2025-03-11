package function

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchDates(artistID int) (DateResponse, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", artistID)
	resp, err := http.Get(url)
	if err != nil {
		return DateResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return DateResponse{}, fmt.Errorf("failed to fetch dates, status code: %d", resp.StatusCode)
	}

	var dateResponse DateResponse
	err = json.NewDecoder(resp.Body).Decode(&dateResponse)
	if err != nil {
		return DateResponse{}, err
	}
	return dateResponse, nil
}
