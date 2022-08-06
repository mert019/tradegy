package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func DecodeJSONFromBody[T any](r *http.Request, model T) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&model)
	if err != nil {
		log.Printf("Error on decoding json from body: %v", err)
		return err
	}
	return nil
}
