package main

import (
	"fmt"
	"time"
)

func main() {
	for range time.Tick(5 * time.Second) {
		fmt.Println("tick")
	}
}
