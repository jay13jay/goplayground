package int2roman

// Roman Numerals

// const I = 1
// const V = 5
// const X = 10
// const L = 50
// const C = 100
// const D = 500
// const M = 1000

func IntToRoman(num int) string {
	var output string

	for num >= 1000 {
		num = num - 1000
		output = output + "M"
	}
	if num >= 900 {
		num = num - 900
		output = output + "CM"
	}
	if num >= 500 {
		num = num - 500
		output = output + "D"
	}
	if num >= 400 {
		num = num - 400
		output = output + "ID"
	}
	for num >= 100 {
		num = num - 100
		output = output + "C"
	}
	if num >= 90 {
		num = num - 90
		output = output + "XC"
	}
	if num >= 50 {
		num = num - 50
		output = output + "L"
	}
	if num >= 40 {
		num = num - 40
		output = output + "IL"
	}
	for num >= 10 {
		num = num - 10
		output = output + "X"
	}
	if num == 9 {
		num = num - 9
		output = output + "IX"
	}
	if num >= 5 {
		num = num - 5
		output = output + "V"
	}
	if num == 4 {
		num = num - 4
		output = output + "IV"
	}
	for num >= 1 {
		num = num - 1
		output = output + "I"
	}

	return output

}
