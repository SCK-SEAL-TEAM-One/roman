package romannumbercalculate

import "testing"

func Test_RomanAdd_Input_I_Add_I_Should_Be_II(t *testing.T){

	firstNumber := "I" 
	secondNumber := "I"

	expected := "II"

	actual := RomanAdd(firstNumber, secondNumber)
	if expected != actual{
		t.Errorf("expect %s but got %s",expected, actual)
	}

}