package utils

import "encoding/json"

func PrintJSON(v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		LogError(err)
		return
	}
	println(string(buf))
}
