package main

import (
	"fmt"
	"github.com/kodekage/flaconi-challenge/utils"
	"reflect"
	"testing"
)

var testJson string = `
[
  {
    "country": "US",
    "city": "Boston",
    "currency": "USD",
    "amount": 100
  }
]`

func TestProcessJsonToMap(t *testing.T) {
	// Arrange
	expected := make([]map[string]interface{}, 1)
	expected[0] = map[string]interface{}{
		"country":  "US",
		"city":     "Boston",
		"currency": "USD",
		"amount":   float64(100),
	}

	// Act
	output := utils.ProcessJsonToMap(testJson)
	eq := reflect.DeepEqual(output, expected)

	// Asset
	if !eq {
		t.Errorf("Output %q not equal to expected %q", output, expected)
	}
}

func TestInputJsonParser(t *testing.T) {
	// Arrange
	args := []string{"currency", "country", "city"}
	expected := map[interface{}]interface{}{
		"USD": map[interface{}]interface{}{
			"US": map[interface{}]interface{}{
				"Boston": []map[interface{}]interface{}{
					0: map[interface{}]interface{}{
						"amount": float64(100),
					},
				},
			},
		},
	}

	// Act
	input := utils.ProcessJsonToMap(testJson)
	parsedJson := utils.InputJsonParser(input, args)

	// Assert
	if fmt.Sprint(expected) != fmt.Sprint(parsedJson) {
		t.Errorf("Output %q not equal to expected %q", parsedJson, expected)
	}
}
