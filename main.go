package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello, World!")
		time.Sleep(1 * time.Second)
	}
}
