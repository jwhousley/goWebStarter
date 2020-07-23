package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"
)

//PageVars variables for pages
type PageVars struct {
	Title      string
	USJsonData UsDataDisplay
}

// UsDataDisplay type for JSON
type UsDataDisplay struct {
	Date                   string
	Cases                  string
	HospitalizedCurrently  string
	HospitalizedCumulative string
	Outcomes               string
	MortalityRate          string
	MortalityPer100k       string
	Death                  string
	Recovered              string
	TotalTestResults       string
}

// UsData type for JSON
type UsData []struct {
	Date                     int       `json:"date"`
	States                   int       `json:"states"`
	Positive                 int       `json:"positive"`
	Negative                 int       `json:"negative"`
	Pending                  int       `json:"pending"`
	HospitalizedCurrently    int       `json:"hospitalizedCurrently"`
	HospitalizedCumulative   int       `json:"hospitalizedCumulative"`
	InIcuCurrently           int       `json:"inIcuCurrently"`
	InIcuCumulative          int       `json:"inIcuCumulative"`
	OnVentilatorCurrently    int       `json:"onVentilatorCurrently"`
	OnVentilatorCumulative   int       `json:"onVentilatorCumulative"`
	Recovered                int       `json:"recovered"`
	DateChecked              time.Time `json:"dateChecked"`
	Death                    int       `json:"death"`
	Hospitalized             int       `json:"hospitalized"`
	LastModified             time.Time `json:"lastModified"`
	Total                    int       `json:"total"`
	TotalTestResults         int       `json:"totalTestResults"`
	PosNeg                   int       `json:"posNeg"`
	DeathIncrease            int       `json:"deathIncrease"`
	HospitalizedIncrease     int       `json:"hospitalizedIncrease"`
	NegativeIncrease         int       `json:"negativeIncrease"`
	PositiveIncrease         int       `json:"positiveIncrease"`
	TotalTestResultsIncrease int       `json:"totalTestResultsIncrease"`
	Hash                     string    `json:"hash"`
}

func main() {

	// serve everything in the css folder
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/usdata", USData)
	http.ListenAndServe(getPort(), nil)
}

func formatCommas(num int) string {
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	//println(r)
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

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
