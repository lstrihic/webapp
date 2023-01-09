package utils

import (
	"encoding/json"
)

type Error struct {
	Code    int `json:"code"`
	Message any `json:"error"`
}

func (err *Error) Error() string {
	marshal, e := json.Marshal(err)
	if e != nil {
		return ""
	}
	return string(marshal)
}
