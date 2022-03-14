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

func ParseSlice[T any](in string) []T {
	var arr []T
	if err := json.Unmarshal([]byte(in), &arr); err != nil {
		log.Panic("parse failed:", err)
	}
	return arr
}

func DumpArray[T any](arr []T) string {
	if len(arr) == 0 {
		return "[]"
	}
	data, err := json.Marshal(arr)
	if err != nil {
		log.Panic("dump failed:", err)
	}
	return string(data)
}
