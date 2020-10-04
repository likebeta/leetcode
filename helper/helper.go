package helper

import (
	"encoding/json"
	"log"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func Dump(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		log.Panic("dump failed:", err)
	}
	return string(data)
}
