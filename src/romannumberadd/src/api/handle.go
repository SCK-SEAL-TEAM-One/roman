package api

import (
	"encoding/json"
	"net/http"

	"romannumber"
)

type RomanResponse struct {
	RomanResult string `json:"romanNumber"`
}

func RomanNumberAddHandler(writer http.ResponseWriter, request *http.Request) {
	firstnumber := request.URL.Query().Get("firstnumber")
	secondnumber := request.URL.Query().Get("secondnumber")

	roman := romannumber.RomanNumberAdd(firstnumber, secondnumber)
	romanJson, _ := json.Marshal(RomanResponse{
		RomanResult: roman,
	})
	writer.Write(romanJson)
}
