package helper

import (
	"encoding/json"
	"log"
)

func parseArray[T any](in string) []T {
	var arr []T
	if err := json.Unmarshal([]byte(in), &arr); err != nil {
		log.Panic("parse failed:", err)
	}
	return arr
}

func dumpArray[T any](arr []T) string {
	if len(arr) == 0 {
		return "[]"
	}
	data, err := json.Marshal(arr)
	if err != nil {
		log.Panic("dump failed:", err)
	}
	return string(data)
}

var ParseIntArray = parseArray[int]
var DumpIntArray = dumpArray[int]

var ParseStrArray = parseArray[string]

var ParseFloatArray = parseArray[float64]

var ParseArray = ParseIntArray
var DumpArray = DumpIntArray
