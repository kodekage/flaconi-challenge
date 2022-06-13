package utils

import (
	"bufio"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

/*
	This function processes JSON string into a map slice
*/
func ProcessJsonToMap(input string) []map[string]interface{} {
	var result []map[string]interface{}

	err := json.Unmarshal([]byte(input), &result)

	if err != nil {
		log.Fatal("ProcessJsonToMap: ", err)
	}

	return result
}

/*
	This function writes the data input into an output.json file
*/
func OutputResult(data map[interface{}]interface{}) {
	body, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	writeError := ioutil.WriteFile("output.json", body, 0644)

	if writeError != nil {
		log.Fatal(err)
	}
}

/*
	This function receives stdin values and parses into a JSON string
*/
func GenerateJsonFromStdin() string {
	scanner := bufio.NewScanner(os.Stdin)
	var jsonString string

	for scanner.Scan() {
		jsonString += scanner.Text()
	}

	if scanner.Err() != nil {
		log.Fatal("GenerateJsonFromStdin: ", scanner.Err())
	}

	return jsonString
}

/*
	This function contains the logic that transforms a slice of a flat map
	and transforms it into a slice of arbitrary nested maps
*/
func InputJsonParser(inputJson []map[string]interface{}, nestKeys []string) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	output := result

	for _, data := range inputJson {
		counter := 0
		for _, key := range nestKeys {
			if counter+1 < len(nestKeys) {
				if val, ok := data[key]; ok {
					if output[key] != val {
						output[data[key]] = make(map[interface{}]interface{})
					}
					output = output[data[key]].(map[interface{}]interface{})
					delete(data, key)
				}
			} else {
				if val, ok := data[key]; ok {
					output[val] = []interface{}{data}
					delete(data, key)
				}
			}
			counter += 1
		}
		output = result
	}

	return result
}
