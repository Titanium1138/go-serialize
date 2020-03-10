package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Demo struct {
	ID   int    `json:"id" mapstructure:"id"`
	Data string `json:"data" mapstructure:"data"`
}

func main() {
	input := Demo{
		ID:   1,
		Data: "hello",
	}

	inputjsonBytes, _ := json.Marshal(input)
	inputjson := string(inputjsonBytes)

	inputMapStruct := make(map[string]interface{})
	json.Unmarshal([]byte(inputjson), &inputMapStruct)

	var output1 Demo
	mapstructure.Decode(inputMapStruct, &output1)
	
	var output2 Demo
	json.Unmarshal([]byte(inputjson), &output2)

	fmt.Printf("input             = [%-23T] %+v\n", input, input)
	fmt.Printf("inputJson         = [%-23T] %+v\n", inputjson, inputjson)
	fmt.Printf("inputMapStruct    = [%-23T] %+v\n", inputMapStruct, inputMapStruct)
	fmt.Printf("output1           = [%-23T] %+v\n", output1, output1)
	fmt.Printf("output2           = [%-23T] %+v\n", output2, output2)
}
