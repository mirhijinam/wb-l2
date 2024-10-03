package net

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/mirhijinam/wb-l2/develop/dev11/internal/models"
)

func parseEvent(r *http.Request) (models.Event, error) {
	err := r.ParseForm()
	if err != nil {
		return models.Event{}, err
	}

	id, _ := strconv.Atoi(r.Form.Get("id"))
	name := r.Form.Get("name")
	data := r.Form.Get("data")
	dateStr := r.Form.Get("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return models.Event{}, errors.New("invalid date format")
	}

	return models.Event{
		ID:   id,
		Name: name,
		Data: data,
		Date: date,
	}, nil
}

func parseID(r *http.Request) (int, error) {
	err := r.ParseForm()
	if err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(r.Form.Get("id"))
	if err != nil {
		return 0, errors.New("invalid ID format")
	}

	return id, nil
}

func (s *Server) errorResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func (s *Server) successResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"result": message})
}

func (s *Server) jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func parseEventWithoutID(r *http.Request) (models.Event, error) {
	err := r.ParseForm()
	if err != nil {
		return models.Event{}, err
	}

	name := r.Form.Get("name")
	data := r.Form.Get("data")
	dateStr := r.Form.Get("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return models.Event{}, errors.New("invalid date format")
	}

	return models.Event{
		Name: name,
		Data: data,
		Date: date,
	}, nil
}
