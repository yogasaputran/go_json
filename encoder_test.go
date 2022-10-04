package gojson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Yoga",
		MiddleName: "Kay",
		LastName:   "Saputra",
	}

	encoder.Encode(customer)
}
