package net

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mirhijinam/wb-l2/develop/dev11/internal/models"
	"github.com/mirhijinam/wb-l2/develop/dev11/internal/service"
)

type servc interface {
	CreateEvent(models.Event) (models.Event, error)
	UpdateEvent(int, models.Event) error
	DeleteEvent(int) error
	GetDailyEvents() ([]models.Event, error)
	GetWeeklyEvents() ([]models.Event, error)
	GetMonthlyEvents() ([]models.Event, error)
}
type Server struct {
	service servc
	mux     *http.ServeMux
}

func New(s *service.Service) *Server {
	server := &Server{
		service: s,
		mux:     http.NewServeMux(),
	}

	server.routes()
	return server
}

func (s *Server) routes() {
	s.mux.HandleFunc("/create_event", s.logMiddleware(s.createEventHandler))
	s.mux.HandleFunc("/update_event", s.logMiddleware(s.updateEventHandler))
	s.mux.HandleFunc("/delete_event", s.logMiddleware(s.deleteEventHandler))
	s.mux.HandleFunc("/events_for_day", s.logMiddleware(s.eventsForDayHandler))
	s.mux.HandleFunc("/events_for_week", s.logMiddleware(s.eventsForWeekHandler))
	s.mux.HandleFunc("/events_for_month", s.logMiddleware(s.eventsForMonthHandler))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *Server) createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.errorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseEventWithoutID(r)
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdEvent, err := s.service.CreateEvent(event)
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	s.successResponse(w, fmt.Sprintf("event created successfully with ID: %d", createdEvent.ID))
}

func (s *Server) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.errorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	event, err := parseEvent(r)
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.service.UpdateEvent(event.ID, event)
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	s.successResponse(w, "event updated successfully")
}

func (s *Server) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		s.errorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := parseID(r)
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.service.DeleteEvent(id)
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	s.successResponse(w, "event is deleted successfully")
}

func (s *Server) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	events, err := s.service.GetDailyEvents()
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	s.jsonResponse(w, events)
}

func (s *Server) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	events, err := s.service.GetWeeklyEvents()
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	s.jsonResponse(w, events)
}

func (s *Server) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		s.errorResponse(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	events, err := s.service.GetMonthlyEvents()
	if err != nil {
		s.errorResponse(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	s.jsonResponse(w, events)
}

func (s *Server) logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	}
}
