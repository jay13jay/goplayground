package azonmissingarray

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate an array of n length with 1 random value discarded. len will be n
func GenerateArray(n int) []int {
	rand.Seed(time.Now().UnixNano())

	var a []int
	mInt := rand.Intn(n)

	for i := 0; i <= n; i++ {
		if i != mInt {
			a = append(a, i)
		}
	}
	return a
}

func FindMissing(a []int) int {
	var missingNumber int

	for _, n := range a {
		i := 0

		// check if
		if n == 1 {
			missingNumber = a[i] + 1
			return missingNumber
		}

	}
	fmt.Println("Finding missing number")

	return missingNumber
}
