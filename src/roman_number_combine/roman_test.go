package roman_number_combine

import "testing"

func Test_CombinRomandNumber_Input_XX_And_II_Should_Be_XXII(t *testing.T) {
	firstNumber := "XX"
	secordNumber := "II"
	expected := "XXII"

	actual := CombinRomandNumber(firstNumber, secordNumber)

	if expected != actual {
		t.Errorf("expect %s but got %s", expected, actual)
	}
}
