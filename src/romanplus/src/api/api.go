package api

import (
	"encoding/json"
	"net/http"
	"romanplus"
)

type RomanResult struct {
	Result string `json:"RomanResult"`
}

func RomanPlusApi(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	firstNumber := query.Get("firstNumber")
	secondNumber := query.Get("secondNumber")
	plusResult := romanplus.RomanPlus(firstNumber, secondNumber)
	newRomanResult := RomanResult{Result: plusResult}
	romanResponseString, _ := json.Marshal(newRomanResult)
	w.Write(romanResponseString)
}
