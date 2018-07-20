package romannumber_test

import (
	. "romannumber"
	"testing"
)

func Test_RomanNumberAdd_Input_FirstNumber_I_SecondNumber_I_Should_Be_II(t *testing.T) {
	firstnumber := "I"
	secondnumber := "I"
	expectedromannumberadd := "II"

	actualroman := RomanNumberAdd(firstnumber, secondnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %s but got %s", expectedromannumberadd, actualroman)
	}

}

func Test_RomanNumberAdd_Input_FirstNumber_II_SecondNumber_II_Should_Be_IV(t *testing.T) {
	firstnumber := "II"
	secondnumber := "II"
	expectedromannumberadd := "IV"

	actualroman := RomanNumberAdd(firstnumber, secondnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %s but got %s", expectedromannumberadd, actualroman)
	}

}

func Test_RomanNumberAdd_Input_FirstNumber_IX_SecondNumber_I_Should_Be_X(t *testing.T) {
	firstnumber := "IX"
	secondnumber := "I"
	expectedromannumberadd := "X"

	actualroman := RomanNumberAdd(firstnumber, secondnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %s but got %s", expectedromannumberadd, actualroman)
	}

}

func Test_RomanNumberAdd_Input_FirstNumber_XXV_SecondNumber_XXV_Should_Be_L(t *testing.T) {
	firstnumber := "XXV"
	secondnumber := "XXV"
	expectedromannumberadd := "L"

	actualroman := RomanNumberAdd(firstnumber, secondnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %s but got %s", expectedromannumberadd, actualroman)
	}

}
func Test_RomanNumberAdd_Input_FirstNumber_L_SecondNumber_L_Should_Be_C(t *testing.T) {
	firstnumber := "L"
	secondnumber := "L"
	expectedromannumberadd := "C"

	actualroman := RomanNumberAdd(firstnumber, secondnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %s but got %s", expectedromannumberadd, actualroman)
	}

}

func Test_RomannumberToNumber_Input_FirstNumber_I_Should_Be_1(t *testing.T) {
	firstnumber := "I"
	expectedromannumberadd := 1

	actualroman := RomannumberToNumber(firstnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %d but got %d", expectedromannumberadd, actualroman)
	}

}

func Test_RomannumberToNumber_Input_FirstNumber_II_Should_Be_2(t *testing.T) {
	firstnumber := "II"
	expectedromannumberadd := 2

	actualroman := RomannumberToNumber(firstnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %d but got %d", expectedromannumberadd, actualroman)
	}

}

func Test_RomannumberToNumber_Input_FirstNumber_IX_Should_Be_9(t *testing.T) {
	firstnumber := "IX"
	expectedromannumberadd := 9

	actualroman := RomannumberToNumber(firstnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %d but got %d", expectedromannumberadd, actualroman)
	}

}

func Test_RomannumberToNumber_Input_FirstNumber_XXV_Should_Be_25(t *testing.T) {
	firstnumber := "XXV"
	expectedromannumberadd := 25

	actualroman := RomannumberToNumber(firstnumber)

	if expectedromannumberadd != actualroman {
		t.Errorf("expected %d but got %d", expectedromannumberadd, actualroman)
	}

}
