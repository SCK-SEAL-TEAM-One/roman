package romanplus_test

import (
	"romanplus"
	"testing"
)

func Test_RomanPlus_Input_I_And_I_Should_Be_II(t *testing.T) {
	firstNumber := "I"
	secondNumber := "I"
	expected := "II"

	actual := romanplus.RomanPlus(firstNumber, secondNumber)

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}
}

func Test_RomanPlus_Input_II_And_II_Should_Be_IV(t *testing.T) {
	firstNumber := "II"
	secondNumber := "II"
	expected := "IV"

	actual := romanplus.RomanPlus(firstNumber, secondNumber)

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}
}
func Test_RomanPlus_Input_IX_And_I_Should_Be_X(t *testing.T) {
	firstNumber := "IX"
	secondNumber := "I"
	expected := "X"

	actual := romanplus.RomanPlus(firstNumber, secondNumber)

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}
}
