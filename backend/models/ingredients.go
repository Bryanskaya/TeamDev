package models

import (
	"api/teamdev/errors"
	"api/teamdev/objects"
	"api/teamdev/repository"
)

type IngredientM struct {
	rep    repository.IngredientsRep
	models *Models
}

func NewIngredient(rep repository.IngredientsRep, models *Models) *IngredientM {
	return &IngredientM{rep, models}
}

func (model *IngredientM) IsExists(ctg string) bool {
	_, err := model.Get(ctg)

	return err == nil
}

func (model *IngredientM) Create(obj *objects.Ingredient) error {
	if model.IsExists(obj.Title) {
		return errors.CategoryExists
	}

	err := model.rep.Create(obj)
	if err != nil {
		return errors.DBAdditionError
	}

	return err
}

func (model *IngredientM) Get(title string) (*objects.Ingredient, error) {
	data, err := model.rep.Get(title)
	if err != nil {
		return nil, errors.RecordNotFound
	}
	return data, err
}

func (model *IngredientM) GetByRecipe(idRecipe int) ([]objects.RecipeIngredient, error) {
	_, err := model.models.Recipes.FindById(idRecipe)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return model.rep.FindByRecipe(idRecipe)
}

func (model *IngredientM) PostToRecipe(id_rcp int, ingArr *[]objects.RecipeIngredient) error {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	for _, obj := range *ingArr {
		err = model.Create(&objects.Ingredient{Title: obj.Ingredient_id})
		if err != nil && err != errors.CategoryExists {
			return err
		}
	}

	return model.rep.ReplaceInRecipe(id_rcp, *ingArr)
}

func (model *IngredientM) AddToRecipe(id_rcp int, obj *objects.RecipeIngredient) error {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	err = model.Create(&objects.Ingredient{Title: obj.Ingredient_id})
	if err != nil && err != errors.CategoryExists {
		return err
	}

	return model.rep.AddToRecipe(obj)
}

func (model *IngredientM) DelFromRecipe(id_rcp int, obj *objects.RecipeIngredient) error {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	if !model.IsExists(obj.Ingredient_id) {
		return errors.UnknownCategory
	}

	return model.rep.DelFromRecipe(id_rcp, obj.Ingredient_id)
}
