package romanaddnumber

import (
	"testing"
)

func Test_AddNumber_Input_XX_And_II_Should_Be_XXII(t *testing.T) {
	expected := "XXII"

	actual := AddNumber("XX", "II")

	if expected != actual {
		t.Errorf("Should be %s but got %s", expected, actual)
	}

}
