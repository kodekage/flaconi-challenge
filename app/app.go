package main

import (
	"github.com/kodekage/flaconi-challenge/utils"
	"os"
)

func main() {
	args := os.Args[1:]

	input := utils.GenerateJsonFromStdin()
	dataToProcess := utils.ProcessJsonToMap(input)
	result := utils.InputJsonParser(dataToProcess, args)

	utils.OutputResult(result)
}
