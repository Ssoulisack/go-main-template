package utilities

import "encoding/json"

func JsonToString(data interface{}) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
