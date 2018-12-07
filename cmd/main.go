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
	Reroll    bool
	Min       int
	Max       int
	// Dave - factorial
	Input     int
	Factorial int
	// Dave - Vdrop
	Voltage_input float64
	Current       float64
	Length        float64
	Area          float64
	//Voltage_drop  float64
	Vdrop         float64
	Vout          float64
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
	// Ensure threshold is bounded by 2d6
	if threshold < 2 {
		threshold = 2
	} else if threshold > 12 {
		threshold = 12
	}

	forward, _ := strconv.ParseBool(r.FormValue("direction"))
	min, _ := strconv.Atoi(r.FormValue("min"))
	max, _ := strconv.Atoi(r.FormValue("max"))
	reroll, _ := strconv.ParseBool(r.FormValue("reroll"))

	result := logisticsmaps.ChanceOfSuccess(threshold, forward, reroll, min, max)
	resultString := strconv.FormatFloat(result*100, 'f', 2, 64)

	p := &Page{Result: resultString, Threshold: threshold, Forward: forward, Min: min, Max: max, Reroll: reroll}
	renderTemplate(w, "dice", p)
}

func factorialHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Factorial Page\n")
	Input, _ := strconv.Atoi(r.FormValue("input"))

	Output := logisticsmaps.Factorializer(Input)

	p := &Page{Factorial: Output}     // assign a value to the block
	renderTemplate(w, "factorial", p) // push the block into the hole
}

func dropHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Voltage Drop Page\n")
	voltage_input, _ := strconv.ParseFloat(r.FormValue("voltage_input"), 64)
	current, _ := strconv.ParseFloat(r.FormValue("current"), 64)
	length, _ := strconv.ParseFloat(r.FormValue("length"), 64)
	area, _ := strconv.ParseFloat(r.FormValue("area"), 64)

	Voltage_drop := logisticsmaps.Drop(current, length, area)
	//VFinal := logisticsmaps.Vfinal(voltage_input, Voltage_drop)
	VFinal := voltage_input - Voltage_drop


	p := &Page{Vdrop: Voltage_drop, Vout: VFinal} // assign a value to the block
	renderTemplate(w, "vdrop", p)                  // push the block into the hole
}

//Dummy page to use for testing
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Serving Hello World Page\n")

	fmt.Fprintf(w, "Hello World!\n")
	fmt.Fprintf(w, "Landing page for various testing")
	
}
func main() {
	port := 9000
	http.HandleFunc("/dice/", diceHandler)
	http.HandleFunc("/", handler)

	// Dave

	http.HandleFunc("/factorial/", factorialHandler)
	http.HandleFunc("/vdrop/", dropHandler)

	fmt.Printf("Listening on Port: %d\n", port)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
