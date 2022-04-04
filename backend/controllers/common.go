package controllers

import (
	auth "api/teamdev/controllers/token"
	"api/teamdev/utils"
	"api/teamdev/models"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/jinzhu/gorm"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initControllers(r *mux.Router, m *models.Models) {
	r.Use(utils.LogHandler)
	r.Use(auth.JwtAuthentication)

	InitAccount(r, m.Accounts)
	InitRecipes(r, m.Recipes)
	InitIngredients(r, m.Ingredients)
	InitLikes(r, m.Recipes, m.Accounts)
	InitSteps(r, m.Steps)
}

func InitRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	models := models.InitModels(db)

	initControllers(router, models)
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