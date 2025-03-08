package main

import (
	"fmt"
	"log"
	"net/http"
)

func main_page_handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This is main page!")
    log.Println("Somebody visited main page!")
}

func handle_left_arrow(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This will handle left arrow click in the future!")
    log.Println("Left arrow clicked!")
}

func handle_right_arrow(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "This will handle right arrow click in the future!")
    log.Println("Right arrow clicked!")
}

func main() {
	fmt.Println("hello from the main function. I am new here.")
    http.HandleFunc("/", main_page_handler)
    http.HandleFunc("/left", handle_left_arrow)
    http.HandleFunc("/right", handle_right_arrow)
	log.Println("Starting http server")
	http.ListenAndServe("localhost:10000", nil)
}
