package controllers

import (
	auth "api/teamdev/controllers/token"
	"api/teamdev/models"
	"api/teamdev/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
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
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin", "Content-Type","content-type","Origin", "Accept"},
    })
	
	handler := c.Handler(r)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
}
