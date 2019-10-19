package utils

import "encoding/json"

func JsonStringConvert(keyVal map[string]string) string {

	jsonString, _ := json.Marshal(keyVal)
	result := string(jsonString)

	return result
}
