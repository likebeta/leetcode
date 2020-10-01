package helper

import (
	"encoding/json"
	"log"
)

func Dump(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Panic("dump failed:", err)
	}
	return string(data)
}
