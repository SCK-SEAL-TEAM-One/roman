package main

import (
	"romannumbercalculate"
	"net/http"
	"encoding/json"
)

type roman struct{
	RomanResult string `json:"romanResult"`
}
func main(){
	http.HandleFunc("/roman/add", RomanAddHandler)
	http.ListenAndServe(":3000", nil)
}

func RomanAddHandler(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	firstNumber := query.Get("firstNumber")
	secondNumber := query.Get("secondNumber")
	romanNumber := romannumbercalculate.RomanAdd(firstNumber, secondNumber)

	roman := roman{
		RomanResult : romanNumber,
	}
	romanResponse, _ := json.Marshal(roman)
	w.Write(romanResponse)
		
}