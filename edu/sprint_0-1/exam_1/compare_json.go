package main

import (
	"encoding/json"
	"reflect"
)

func CompareJSON(json1, json2 []byte) (bool, error) {
	var data1, data2 map[string]interface{}

	err := json.Unmarshal(json1, &data1)
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(json2, &data2)
	if err != nil {
		return false, err
	}

	if len(data1) != len(data2) {
		return false, nil
	}

	for key, value := range data1 {
		if val, ok := data2[key]; !ok || !reflect.DeepEqual(value, val) {
			return false, nil
		}
	}

	return true, nil
}
