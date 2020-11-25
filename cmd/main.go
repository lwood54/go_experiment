package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/lwood54/go_experiment/math"
)

func main() {
	numAsString := strconv.Itoa(math.Add(1, 5))

	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, numAsString)
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
