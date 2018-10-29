package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/duxbuse/LogisticsMaps"
)

/*
Page represents all the data that is sent to the dice roller webpage
This is for showing the results as well as inputs.
*/
type Page struct {
	Result    string
	Threshold int
	Forward   bool
	Min       int
	Max       int
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("./../views/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func diceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Dice Page\n")
	threshold, _ := strconv.Atoi(r.FormValue("threshold"))
	forward, _ := strconv.ParseBool(r.FormValue("direction"))
	min, _ := strconv.Atoi(r.FormValue("min"))
	max, _ := strconv.Atoi(r.FormValue("max"))

	result := LogisticsMaps.ChanceOfSuccess(threshold, forward, min, max)
	resultString := strconv.FormatFloat(result*100, 'f', 2, 64)

	p := &Page{Result: resultString, Threshold: threshold, Forward: forward, Min: min, Max: max}
	renderTemplate(w, "dice", p)
}

//Dummy page to use for testing
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Hello World Page\n")

	fmt.Fprintf(w, "Hello World!")
}
func main() {
	port := 9000
	http.HandleFunc("/dice/", diceHandler)
	http.HandleFunc("/", handler)

	fmt.Printf("Listening on Port: %d\n", port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
