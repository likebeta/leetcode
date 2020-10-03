package helper

import (
	"encoding/json"
	"log"
)

func ParseArray(in string) []int {
	var arr []int
	if err := json.Unmarshal([]byte(in), &arr); err != nil {
		log.Panic("parse failed:", err)
	}
	return arr
}

func DumpArray(arr []int) string {
	if len(arr) == 0 {
		return "[]"
	}
	data, err := json.Marshal(arr)
	if err != nil {
		log.Panic("dump failed:", err)
	}
	return string(data)
}
