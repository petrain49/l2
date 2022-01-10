package server

import (
	"dev11/internal/cache"
	"encoding/json"
	"net/http"
)

type ResultResponse struct {
	Events []cache.Event `json:"result"`
}

func ResponseWithResult(w http.ResponseWriter, events []cache.Event) error {
	resp, err := json.Marshal(ResultResponse{events})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}

type ErrorResponse struct {
	Err string `json:"error"`
}

func ResponseWithError(w http.ResponseWriter, err error, code int) error {
	resp, marshalError := json.Marshal(ErrorResponse{err.Error()})
	if marshalError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return marshalError
	}

	http.Error(w, string(resp), code)
	return nil
}
