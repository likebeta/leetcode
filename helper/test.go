package helper

import (
	"fmt"
	"log"
)

func Assert(b bool) {
	if !b {
		log.Panic()
	}
}

func Log(a ...interface{}) {
	if n, err := fmt.Println(a...); err != nil {
		log.Panic(n, err)
	}
}

func LogF(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	if n, err := fmt.Println(s); err != nil {
		log.Panic(n, err)
	}
}

func Print(a ...interface{}) {
	if n, err := fmt.Print(a...); err != nil {
		log.Panic(n, err)
	}
}
