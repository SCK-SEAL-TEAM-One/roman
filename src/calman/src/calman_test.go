package calman

import "testing"

func Test_PlusRoman_Input_XX_And_II_Should_Be_XXII(t *testing.T) {
	firstNumber := "XX"
	secondNumber := "II"
	expected := "XXII"

	actual := PlusRoman(firstNumber, secondNumber)

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}
}
