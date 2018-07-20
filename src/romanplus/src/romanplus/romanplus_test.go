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
