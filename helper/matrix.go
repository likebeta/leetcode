package helper

import (
	"encoding/json"
	"log"
)

func ParseMatrix[T any](in string) [][]T {
	var matrix [][]T
	if err := json.Unmarshal([]byte(in), &matrix); err != nil {
		log.Panic("parse failed:", err)
	}
	return matrix
}

func DumpMatrix[T any](matrix [][]T) string {
	if len(matrix) == 0 {
		return "[]"
	}
	data, err := json.Marshal(matrix)
	if err != nil {
		log.Panic("dump failed:", err)
	}
	return string(data)
}
