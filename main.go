package main

import (
	"fmt"

	"gitlab.com/jhax/goplayground/maptest"
)

func main() {
	fmt.Println("Starting")
	out, err := maptest.MapPractice()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Output below:\n%s", out)
}
