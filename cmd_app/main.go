package main

import (
	"bufio"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"log"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	args := os.Args[1:]

	input := generateJsonFromStdin()
	dataToProcess := processJsonToMap(input)
	result := inputJsonParser(dataToProcess, args)

	outputResult(result)
}

func processJsonToMap(input string) []map[string]interface{} {
	var result []map[string]interface{}

	err := json.Unmarshal([]byte(input), &result)

	if err != nil {
		log.Fatal("ProcessJsonToMap: ", err)
	}

	return result
}

func outputResult(data map[interface{}]interface{}) {
	body, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	writeError := ioutil.WriteFile("output.json", body, 0644)

	if writeError != nil {
		log.Fatal(err)
	}
}

func generateJsonFromStdin() string {
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

func inputJsonParser(inputJson []map[string]interface{}, nestKeys []string) map[interface{}]interface{} {
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
