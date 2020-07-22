package main

import (
	"net/http"
)

//Home handler for / renders the home.html
func Home(w http.ResponseWriter, req *http.Request) {
	pageVars := PageVars{
		Title: "Home",
	}
	render(w, "home.html", pageVars)
}
