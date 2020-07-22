package main

import (
	"fmt"
	"net/http"
)

//USData handler for / renders the usdata.html
func USData(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Loading Data...")

	usData := UsData{}
	getJSON("https://covidtracking.com/api/us", &usData)

	pageVars := PageVars{
		Title:      "U.S. Data",
		USJsonData: usData,
	}
	render(w, "usdata.html", pageVars)

}
