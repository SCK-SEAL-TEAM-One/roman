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
