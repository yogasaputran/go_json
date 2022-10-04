package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "P0001",
		Name:     "Apple Macbook Pro",
		ImageUrl: "https://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Macbook Pro","image_url":"https://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	err := json.Unmarshal(jsonBytes, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)

}
