package roman_number_calculator

import "testing"

func Test_RomanCalculate_Input_FirstNumber_XX_SecondNumber_II_Should_Be_XXII(t *testing.T) {
	firstNumber := "XX"
	secondNumber := "II"
	expected := "XXII"

	actual := RomanCalculate(firstNumber, secondNumber)
	if expected != actual {
		t.Errorf("expected is %s but it got %s", expected, actual)
	}

}
