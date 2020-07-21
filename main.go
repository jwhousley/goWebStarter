package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//PageVars variables for pages
type PageVars struct {
	Title string
}

func main() {

	// serve everything in the css folder
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", Home)
	http.ListenAndServe(getPort(), nil)
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	io.WriteString(w, `Hello World`)
// }

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func render(w http.ResponseWriter, tmpl string, pageVars PageVars) {

	tmpl = fmt.Sprintf("templates/%s", tmpl) // prefix the name passed in with templates/
	t, err := template.ParseFiles(tmpl)      //parse the template file held in the templates folder

	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}

	err = t.Execute(w, pageVars) //execute the template and pass in the variables to fill the gaps

	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
