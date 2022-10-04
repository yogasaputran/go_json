package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	jsonString := `{"id":"P0001","name":"Apple Macbook Pro","image_url":"https://example.com/image.png"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":        "P0001",
		"name":      "Apple Macbook Pro",
		"image_url": "https://example.com/image.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
