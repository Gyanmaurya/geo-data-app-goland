
package services

import (
	"encoding/json"
	"errors"
)

func ProcessFile(fileData []byte) ([]byte, error) {
	var geoJSON map[string]interface{}
	err := json.Unmarshal(fileData, &geoJSON)
	if err != nil {
		return nil, errors.New("Invalid GeoJSON")
	}
	return json.Marshal(geoJSON)
}
