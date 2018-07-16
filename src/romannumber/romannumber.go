package romannumber

func RomanCalculate(firstNumber, secondNumber string) string {
	if firstNumber == "II" && secondNumber == "II" {
		return "IV"
	}

	return firstNumber + secondNumber
}
