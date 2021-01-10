package handler

import (
	"encoding/json"
	"net/http"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

const (
	Error   = "error"
	Message = "message"
)

func newResponse(messageType string, messsage string, data interface{}) response {
	return response{
		messageType,
		messsage,
		data,
	}
}

// responseJSON funcion que se encarga de responder al cliente
func responseJSON(w http.ResponseWriter, statusCode int, resp response) {
	w.Header().Set("Content-Type", "applicaction/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
