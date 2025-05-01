package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan string)

    go func() {
        ch <- "goroutine!"
    }()

	time.Sleep(3 * time.Second) 
    fmt.Println(<-ch)
}