package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Yoga",
		MiddleName: "Sa",
		LastName:   "putra",
		Hobbies:    []string{"Gaming", "Coding", "Swimming"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Yoga","MiddleName":"Sa","LastName":"putra","Age":10,"Married":false,"Hobbies":["Gaming","Coding","Swimming"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Yoga",
		Address: []Address{
			{
				Street:     "Jalan Belum Ada",
				Country:    "Indonesia",
				PostalCode: "9999",
			},
			{
				Street:     "Jalan Lagi Dibangun",
				Country:    "Indonesia",
				PostalCode: "8888",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Yoga","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Address":[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"8888"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Belum Ada","Country":"Indonesia","PostalCode":"9999"},{"Street":"Jalan Lagi Dibangun","Country":"Indonesia","PostalCode":"8888"}]`

	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Belum Ada",
			Country:    "Indonesia",
			PostalCode: "9999",
		},
		{
			Street:     "Jalan Lagi Dibangun",
			Country:    "Indonesia",
			PostalCode: "8888",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
