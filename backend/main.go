package main

import (
	"api/teamdev/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main()  {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		utils.InitConfig()
	} else {
		utils.InitConfig(os.Args[1])
	}


	r := mux.NewRouter()
	r.HandleFunc("/test", getTest).Methods("GET")
	fmt.Printf("Server is running on http://localhost:%d\n", utils.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", utils.Config.Port),  r)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	var data = [...]int{1, 2, 3, 4}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}
