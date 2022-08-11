package main

import (
	"fmt"
	"math/rand"
	"time"

	"gitlab.com/jhax/goplayground/azonmissingarray"
)

const alen = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	var a []int

	a = azonmissingarray.GenerateArray(alen)

	fmt.Printf("Created array: %v\n", a)
}
