package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

// 1. ECHO

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror %s\n", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror reading CSV: %s\n", err.Error())))
		return
	}

	response := Echo(records)
	fmt.Fprint(w, "\n"+response+"\n")
}

// 2. INVERT

func InvertHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror %s\n", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror reading CSV: %s\n", err.Error())))
		return
	}

	response := Invert(records)
	fmt.Fprint(w, "\n"+response+"\n")
}

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.HandleFunc("/invert", InvertHandler)

	fmt.Println("Server listening!")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
