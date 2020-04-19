package main

import (
	"dep_check/src"
	"fmt"
	"log"
	"time"
)

func main() {
	duration := time.Now()
	err := src.DepCheck()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Executed in", time.Since(duration))
}