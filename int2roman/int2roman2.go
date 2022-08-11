package int2roman

import (
	"strings"
)

func IntToRomanMap(num int) string {
	var numeral string // Roman Numeral to return
	k := [15]int{1000,
		900,
		500,
		400,
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}
	m := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		3:    "III",
		2:    "II",
		1:    "I",
	}
	for _, n := range k {
		// fmt.Printf("%d | %s\n", i, n)
		if n > 0 {
			div := num / n
			if div >= 1 {
				numeral = numeral + strings.Repeat(m[n], div)
				num = num - n*div
			}
		}
	}

	return numeral

}
