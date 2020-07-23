package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//USData handler for / renders the usdata.html
func USData(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Loading Data (USData)...")

	////displayData := UsDataDisplay{}
	usData := UsData{}
	getJSON("https://covidtracking.com/api/us", &usData)

	////displayData.Cases = formatCommas(usData[0].Positive)

	pageVars := PageVars{
		Title:      "U.S. Data",
		USJsonData: NewDisplayData(usData),
	}
	render(w, "usdata.html", pageVars)
}

// NewDisplayData creates a new UsDataDisplay
func NewDisplayData(item UsData) UsDataDisplay {

	displayData := UsDataDisplay{}
	cases := formatCommas(item[0].Positive)
	death := formatCommas(item[0].Death)
	recovered := formatCommas(item[0].Recovered)
	totalTests := formatCommas(item[0].TotalTestResults)
	hospCurrent := formatCommas(item[0].HospitalizedCurrently)
	hospCumulative := formatCommas(item[0].HospitalizedCumulative)

	outcomesCalc := item[0].Death + item[0].Recovered
	outcomes := formatCommas(outcomesCalc)
	////var mortalityCalc float32
	mortalityCalc := float64(item[0].Death) / 330724728
	mortality := strconv.FormatFloat(mortalityCalc, 'f', 5, 64)

	mortalityPer100kCalc := mortalityCalc * 100000
	mortalityPer100k := strconv.FormatFloat(mortalityPer100kCalc, 'f', 0, 64)

	////formatCommas(math.Round(mortalityCalc))
	////mortality := math.Round(mortalityCalc)

	displayData.Cases = cases
	displayData.Death = death
	displayData.Outcomes = outcomes
	displayData.Recovered = recovered
	displayData.TotalTestResults = totalTests
	displayData.HospitalizedCurrently = hospCurrent
	displayData.HospitalizedCumulative = hospCumulative
	displayData.MortalityRate = mortality
	displayData.MortalityPer100k = mortalityPer100k

	return displayData
}
