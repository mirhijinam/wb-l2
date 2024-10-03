package utils

import (
	"encoding/json"

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
