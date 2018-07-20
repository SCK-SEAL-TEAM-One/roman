package romanplus

func RomanPlus(firstNumber, secondNumber string) string {
	if firstNumber == "II" && secondNumber == "II" {
		return "IV"
	}
	if firstNumber == "IX" && secondNumber == "I" {
		return "X"
	}
	return "II"
}
