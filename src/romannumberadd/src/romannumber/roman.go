package romannumber

func RomanNumberAdd(firstnumber, secondnumber string) string {
	numberFromAdd := RomannumberToNumber(firstnumber) + RomannumberToNumber(secondnumber)
	return ConvertToRomanNumber(numberFromAdd)
}

func RomannumberToNumber(firstnumber string) int {

	roman := make(map[string]int)
	roman["I"] = 1
	roman["IV"] = 4
	roman["V"] = 5
	roman["IX"] = 9
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	var number int
	for _, romanvalue := range firstnumber {
		if roman[string(romanvalue)] > number {
			number = roman[string(romanvalue)] - number
		} else {
			number += roman[string(romanvalue)]
		}

	}

	return number
}

func ConvertToRomanNumber(number int) string {
	valueRomanNumber := []int{100, 90, 50, 40, 30, 10, 9, 5, 4, 1}
	roman := make(map[int]string)
	roman[1] = "I"
	roman[4] = "IV"
	roman[5] = "V"
	roman[9] = "IX"
	roman[10] = "X"
	roman[30] = "XXX"
	roman[40] = "XL"
	roman[50] = "L"
	roman[90] = "XC"
	roman[100] = "C"
	romanNumber := ""

	for _, arabicNumber := range valueRomanNumber {
		for number >= arabicNumber {
			number -= arabicNumber
			romanNumber += roman[arabicNumber]
		}
	}

	return romanNumber
}
