package utils

import (
	"encoding/json"
)

func ToJson(v interface{}) string{
	jn,err :=json.Marshal(v);
	if err != nil {
		return ""
	}
	return string(jn)
}
