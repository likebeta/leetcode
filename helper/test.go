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
