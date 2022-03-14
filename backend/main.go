package main

import (
	"api/teamdev/controllers"
	"api/teamdev/utils"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// @Title TeamDev API
// @Version 0.1
// @Description API for culinary recipes (BMSTU Team development project)
// @securityDefinitions.basic BasicAuth
func main()  {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 1 {
		utils.InitConfig()
	} else {
		utils.InitConfig(os.Args[1])
	}

	utils.InitLogger()
	defer utils.CloseLogger()

	r := controllers.InitRouter(nil)
	controllers.RunSwagger(r);

	r.HandleFunc("/test", getTest).Methods("GET")

	utils.Logger.Print("Server started")
	fmt.Printf("Server is running on http://localhost:%d\n", utils.Config.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", utils.Config.Port),  r)
	code := controllers.RunRouter(r, utils.Config.Port)

	utils.Logger.Printf("Server ended with code %s", code)
}

// @Tags Categories
// @Router /categories [get]
// @Summary Retrieves all categories
// @Produce json
func getTest(w http.ResponseWriter, r *http.Request) {
	var data = [...]int{1, 2, 3, 4}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}
