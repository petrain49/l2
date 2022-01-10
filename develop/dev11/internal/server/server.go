package server

import (
	"context"
	"dev11/internal/cache"
	"errors"
	"log"
	"net/http"
)

type Service struct {
	server http.Server
	cache  *cache.Cache
}

func NewService(addr string, cache *cache.Cache) *Service {
	service := new(Service)
	service.cache = cache
	service.server.Addr = addr

	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", LogRequest(http.HandlerFunc(service.HandleCreateEvent)))
	mux.HandleFunc("/update_event", LogRequest(http.HandlerFunc(service.HandleUpdateEvent)))
	mux.HandleFunc("/delete_event", LogRequest(http.HandlerFunc(service.HandleDeleteEvent)))
	mux.HandleFunc("/events_for_day", LogRequest(http.HandlerFunc(service.HandleGetEventsForDay)))
	mux.HandleFunc("/events_for_week", LogRequest(http.HandlerFunc(service.HandleGetEventsForWeek)))
	mux.HandleFunc("/events_for_month", LogRequest(http.HandlerFunc(service.HandleGetEventsForMonth)))

	service.server.Handler = mux

	return service
}

func (s *Service) RunServer() error {
	err := s.server.ListenAndServe()
	if err == http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Service) StopServer() {
	s.server.Shutdown(context.TODO())
}

// POST ----------------------------------------------

func (s *Service) HandleCreateEvent(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	event, err := ParsePostQueryStringToEvent(req.Form)
	if err != nil {
		log.Println("Error parse query string:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	s.cache.CreateEvent(event)

	ResponseWithResult(w, []cache.Event{})
}

func (s *Service) HandleUpdateEvent(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	event, err := ParsePostQueryStringToEvent(req.Form)
	if err != nil {
		log.Println("Error parse query string:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	err = s.cache.UpdateEvent(event)
	if err != nil {
		log.Println("Error parse query string:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	ResponseWithResult(w, []cache.Event{})
}

func (s *Service) HandleDeleteEvent(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	EID := req.Form.Get(queryEventID)
	if EID == "" {
		log.Println("Incomplete query string")

		queryError := errors.New("incomplete query string")
		ResponseWithError(w, queryError, http.StatusBadRequest)
		return
	}

	s.cache.DeleteEvent(EID)

	ResponseWithResult(w, []cache.Event{})
}

// GET -----------------------------------------------

func (s *Service) HandleGetEventsForDay(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	UID, date, err := ParseGetQueryString(req.Form)
	if err != nil {
		log.Println("Wrong format:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
	}

	events, err := s.cache.GetEventsForDay(UID, date)
	if err != nil {
		log.Println("No events:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	ResponseWithResult(w, events)
}

func (s *Service) HandleGetEventsForMonth(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	UID, date, err := ParseGetQueryString(req.Form)
	if err != nil {
		log.Println("Wrong format:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
	}

	events, err := s.cache.GetEventsForMonth(UID, date)
	if err != nil {
		log.Println("No events:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	ResponseWithResult(w, events)
}

func (s *Service) HandleGetEventsForWeek(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		log.Println("Error parsing form:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	UID, date, err := ParseGetQueryString(req.Form)
	if err != nil {
		log.Println("Wrong format:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
	}

	events, err := s.cache.GetEventsForWeek(UID, date)
	if err != nil {
		log.Println("No events:", err)

		ResponseWithError(w, err, http.StatusBadRequest)
		return
	}

	ResponseWithResult(w, events)
}
