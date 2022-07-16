package main

import (
	"fmt"

	"gitlab.com/jhax/goplayground/twosum"
)

func main() {
	fmt.Println("Starting")
	// out, _ := maptest.MapPractice()
	var target int = 6
	var numList = []int{2, 3, 6, 3}

	out := twosum.TwoSum_HM(numList, target)
	fmt.Printf("Output below:\n%v\n", out)
}
