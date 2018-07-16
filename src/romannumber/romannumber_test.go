package romannumber

import "testing"

func Test_RomanCalculate_Input_FirstNumber_XX_And_SecondNumber_II_Should_Be_XXII(t *testing.T) {
	firstNumber := "XX"
	secondNumber := "II"

	expected := "XXII"

	actual := RomanCalculate(firstNumber, secondNumber)

	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func Test_RomanCalculate_Input_FirstNumber_II_And_SecondNumber_II_Should_Be_IV(t *testing.T) {
	firstNumber := "II"
	secondNumber := "II"

	expected := "IV"

	actual := RomanCalculate(firstNumber, secondNumber)

	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
