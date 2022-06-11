package main

import (
	"github.com/kodekage/flaconi-challenge/cmd-app/app"
	"os"
)

func main() {
	args := os.Args[1:]

	input := app.GenerateJsonFromStdin()
	dataToProcess := app.ProcessJsonToMap(input)
	result := app.InputJsonParser(dataToProcess, args)

	app.OutputResult(result)
}
