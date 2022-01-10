package server

import (
	"dev11/internal/cache"
	"log"
	"net/http"
	"net/url"
)

const (
	queryUserID      = "user_id"
	queryEventID     = "event_id"
	queryDescription = "description"
	queryDate        = "date"
)

func ParsePostQueryStringToEvent(form url.Values) (cache.Event, error) {
	UID := form.Get(queryUserID)
	EID := form.Get(queryEventID)
	description := form.Get(queryDescription)
	date := form.Get(queryDate)

	event, err := cache.NewEvent(UID, EID, description, date)
	if err != nil {
		return event, err
	}

	return event, err
}

func ParseGetQueryString(form url.Values) (string, cache.DateInfo, error) {
	UID := form.Get(queryUserID)
	date := form.Get(queryDate)

	UID, dateParsed, err := cache.CheckUIDAndDate(UID, date)
	if err != nil {
		return "", cache.DateInfo{}, err
	}

	return UID, dateParsed, err
}

func LogRequest(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Request: %s %s", req.Method, req.RequestURI)
		handler.ServeHTTP(w, req)
	}
}