package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Yoga",
		MiddleName: "Sa",
		LastName:   "Putra",
		Age:        26,
		Married:    true,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
