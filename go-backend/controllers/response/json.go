package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Status  int
	Payload interface{}
	Message string
}

func JSON(w http.ResponseWriter, status int, message string, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Status:  status,
		Message: message,
		Payload: payload,
	}

	enc, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return
	}

	_, err = w.Write(enc)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
