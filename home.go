package main

import (
	"fmt"
	"log"
	"net/http"

	"./eliza"
	//	"strings"
	//	"bytes"
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// this is code that runs when a request is made to the /ask resource.
	userInput := r.URL.Query().Get("user-input")
	reply := eliza.AskEliza(userInput)
	fmt.Fprintf(w, reply)

} //chatHandler

func main() {

	//serve the files from the /static folder
	dir := http.Dir("./webApp")
	fileServer := http.FileServer(dir)

	//handle requests to /
	http.Handle("/", fileServer)
	//handle request to /chat
	http.HandleFunc("/chat", chatHandler)

	log.Println("Listening....")
	http.ListenAndServe(":8080", nil)
}
