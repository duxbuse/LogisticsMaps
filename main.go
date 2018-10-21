package main

import (
	"strconv"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Result string
	Threshold int
	Forward bool
	Min int
	Max int
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles("views/" + tmpl + ".html")
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
	threshold, _ := strconv.Atoi(r.FormValue("threshold"))
	forward, _ := strconv.ParseBool(r.FormValue("direction"))
	min, _ := strconv.Atoi(r.FormValue("min"))
	max, _ := strconv.Atoi(r.FormValue("max"))
	

	result := chanceOfSuccess(threshold,forward, min, max)
	resultString := strconv.FormatFloat(result*100, 'f', 2, 64)

	p := &Page{Result: resultString, Threshold: threshold, Forward: forward, Min: min, Max: max}
	renderTemplate(w, "dice", p)
}

func main() {
	http.HandleFunc("/dice/", diceHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
