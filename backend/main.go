package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	port := 9000

	r := mux.NewRouter()
	r.HandleFunc("/test", getTest).Methods("GET")
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port),  r)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	var data = [...]int{1, 2, 3, 4}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}
