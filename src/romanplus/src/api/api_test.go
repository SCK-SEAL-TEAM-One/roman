package api_test

import (
	"api"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_RomanPlus_Input_I_And_I_Should_Be_II(t *testing.T) {
	req := httptest.NewRequest("GET", "/roman/add?firstNumber=I&secondNumber=I", nil)
	w := httptest.NewRecorder()
	api.RomanPlusApi(w, req)
	expected := api.RomanResult{Result: "II"}

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	actual := api.RomanResult{}
	json.Unmarshal(body, &actual)

	if expected != actual {
		t.Errorf("should be %v but is %v", expected, actual)
	}
}
