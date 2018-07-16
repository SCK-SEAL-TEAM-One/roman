package roman_aunpukkad_test

import (
	"roman_aunpukkad"
	"testing"
)

func Test_Input_XX_and_II_Should_Be_XXII(t *testing.T) {
	expected := "XXII"
	firstNumber := "XX"
	secondNumber := "II"

	actual := roman_aunpukkad.Calculate(firstNumber, secondNumber)

	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}

func Test_Input_II_and_II_Should_Be_IV(t *testing.T) {
	expected := "IV"
	firstNumber := "II"
	secondNumber := "II"

	actual := roman_aunpukkad.Calculate(firstNumber, secondNumber)

	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
