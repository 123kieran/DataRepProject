/*
*	Author	: Kieran O'Halloran
*	Data Represenation and Querying Project
 */

package main

import (
	"html/template" //add html/template package
	"net/http"
)

type myMsg struct {
	Input    string
	Output   string
	Previous string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	//serve the homepage.html file
	http.ServeFile(w, r, "elizachat.html")
}

func chatHandler(w http.ResponseWriter, r *http.Request) {

	//create and initialise string
	output := "...........     "
	input := r.FormValue("chat")
	previous := input

	t, _ := template.ParseFiles("elizachat.html")

	//execute template and pass pointer to myMsg 	struct
	t.Execute(w, &myMsg{Input: input, Output: output, Previous: previous})
} //chatHandler

func main() {
	// handles root page
	http.HandleFunc("/", requestHandler)

	//handle /chat page
	http.HandleFunc("/elixachat", chatHandler)
	http.ListenAndServe(":8080", nil)
}
