package models

import (
	"api/teamdev/repository"

	"github.com/jinzhu/gorm"
)

type Models struct {
	Accounts *AccountM
	Recipes *RecipeM
	Steps *StepM
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)

	models.Accounts = NewAccount(repository.NewAccountsRep(db), models)
	models.Recipes = NewRecipe(repository.NewRecipesRep(db), models)
	models.Steps = NewStep(repository.NewStepsRep(db), models)

	return models
}