package main

import (
	"encoding/json"
	"fmt"
)

func mergeJSONData(jsonDataList ...[]byte) ([]byte, error) {
	var mergedData []map[string]interface{}
	for _, jsonData := range jsonDataList {
		var jsonDataMap []map[string]interface{}
		if err := json.Unmarshal(jsonData, &jsonDataMap); err != nil {
			return nil, err
		}
		for _, item := range jsonDataMap {
			mergedData = append(mergedData, item)
		}
	}
	mergedJSON, err := json.Marshal(mergedData)
	if err != nil {
		return nil, err
	}
	return mergedJSON, nil
}

func main() {
	jsonData1 := []byte(`[{"name": "Oleg","class": "9B"},{"name": "Ivan","class": "9A"}]`)
	jsonData2 := []byte(`[{"name": "Maria",	"class": "9B"},{"name": "John","class": "9A"}]`)
	mergedJSON, err := mergeJSONData(jsonData1, jsonData2)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(mergedJSON))
}
