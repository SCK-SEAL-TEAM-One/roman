package main

import (
	"api"
	"net/http"
)

func main() {
	http.HandleFunc("/roman/add", api.RomanPlusApi)
	http.ListenAndServe(":3000", nil)
}
