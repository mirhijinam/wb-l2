package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/mirhijinam/wb-l2/develop/dev11/internal/models"
)

type serializedEvents struct {
	SerializedEvents []models.Event `json:"result"`
}

func SerializeEvents(
	events []models.Event,
) ([]byte, error) {
	data := serializedEvents{
		events,
	}

	result, err := json.Marshal(data)

	return result, err
}

func ValidateEvent(e models.Event) error {
	if e.Name == "" {
		return errors.New("name is empty")
	}

	if len(e.Name) > 100 {
		return errors.New("name is too long")
	}

	if e.Date.IsZero() {
		return errors.New("date is not specified")
	}

	if e.Date.Before(time.Now()) {
		return errors.New("date is in the past")
	}

	if len(e.Data) > 1000 {
		return errors.New("description is too long")
	}

	return nil
}
