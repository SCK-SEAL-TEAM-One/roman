package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_API_RomanNumberAddHandler_Input_1_Should_be_RomanNumberAdd(t *testing.T) {
	url := "http://localhost:3000/roman/add?firstnumber=1&secondnumber=1"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expectedResponse := RomanResponse{
		RomanResult: "II",
	}
	RomanNumberAddHandler(responseRecorder, request)
	response := responseRecorder.Result()
	body, _ := ioutil.ReadAll(response.Body)

	var actual RomanResponse
	json.Unmarshal(body, &actual)

	if expectedResponse != actual {
		t.Errorf("Should Be %v but it got %v", expectedResponse, actual)
	}
}
