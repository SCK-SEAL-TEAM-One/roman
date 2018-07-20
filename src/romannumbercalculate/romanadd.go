package romannumbercalculate

func RomanAdd(firstRoman,secondRoman string) string{

	romanNumber := make(map[string]int)
	romanNumber["I"] = 1

	romancharacter := make(map[int]string)
	romancharacter[2] = "II"
	firstNumber := romanNumber[firstRoman]
	secondNumber := romanNumber[secondRoman]
	arabicNumber := firstNumber + secondNumber
	resultRoman := romancharacter[arabicNumber]

	return resultRoman
	

}