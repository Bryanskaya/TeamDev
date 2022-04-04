package models

import (
	"api/teamdev/repository"

	"github.com/jinzhu/gorm"
)

type Models struct {
	Accounts *AccountM
	Recipes *RecipeM
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)

	models.Accounts = NewAccount(repository.NewAccountsRep(db), models)

	return models
}