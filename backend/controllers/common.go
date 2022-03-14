package controllers

import (
	"api/teamdev/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/jinzhu/gorm"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initControllers(r *mux.Router) {
	r.Use(utils.LogHandler)
}

func InitRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	initControllers(router)
	return router
}

func RunSwagger(r *mux.Router) {
	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}

func RunRouter(r *mux.Router, port uint16) error {
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "PUT"},
    })
	
	handler := c.Handler(r)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}