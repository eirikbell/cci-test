package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println(helloWorld())
		}
	}
}

func helloWorld() string {
	return "Hello World"
}
