package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

// 1. ECHO

func EchoHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%v", http.StatusText(404))))
		return
	}

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

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%v", http.StatusText(404))))
		return
	}

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

// 3. FLATTEN
func FlattenHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%v", http.StatusText(404))))
		return
	}

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

	response := Flatten(records)
	fmt.Fprint(w, "\n"+response+"\n")
}

// 4. SUM
func SumHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%v", http.StatusText(404))))
		return
	}

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

	response, err := Sum(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror computing sum: %s\n", err.Error())))
		return
	}
	fmt.Fprint(w, "\n"+response+"\n")
}

// 5. MULTIPLY
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("%v", http.StatusText(404))))
		return
	}

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

	response, err := Multiply(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror computing multiply: %s\n", err.Error())))
		return
	}
	fmt.Fprint(w, "\n"+response+"\n")
}

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.HandleFunc("/invert", InvertHandler)
	http.HandleFunc("/flatten", FlattenHandler)
	http.HandleFunc("/sum", SumHandler)
	http.HandleFunc("/multiply", MultiplyHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Run server error: ", err)
	}

}
