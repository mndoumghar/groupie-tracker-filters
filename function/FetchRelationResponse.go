package function

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchRelations(artistID int) (RelationReponse, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", artistID)
	resp, err := http.Get(url)
	if err != nil {
		return RelationReponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return RelationReponse{}, fmt.Errorf("failed to fetch relation, status code: %d", resp.StatusCode)
	}

	var relationResponse RelationReponse
	err = json.NewDecoder(resp.Body).Decode(&relationResponse)
	if err != nil {
		return RelationReponse{}, err
	}
	return relationResponse, nil
}
