package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJson(t *testing.T) {
	jsonString := `{
		"FirstName":"Yoga",
		"MiddleName":"Sa",
		"LastName":"putra",
		"Age":26,"Married":true
	}`

	jsonBytes := []byte(jsonString)
	customer := &Customer{}

	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
}
