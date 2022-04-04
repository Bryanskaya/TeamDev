package models

import (
	"api/teamdev/errors"
	"api/teamdev/objects"
	"api/teamdev/repository"
)

type IngredientM struct {
	rep repository.IngredientsRep
	models *Models
}

func NewIngredient(rep repository.IngredientsRep, models *Models) *IngredientM {
	return &IngredientM{rep, models}
}

func (this *IngredientM) IsExists(ctg string) bool {
	_, err := this.Get(ctg)

	return err == nil
}

func (this *IngredientM) Create(obj *objects.Ingredient) error {
	if this.IsExists(obj.Title) {
		return errors.CategoryExists
	}

	err := this.rep.Create(obj)
	if err != nil {
		return errors.DBAdditionError
	}

	return err
}

func (this *IngredientM) Get(title string) (*objects.Ingredient, error) {
	data, err := this.rep.Get(title)
	if err != nil {
		return nil, errors.RecordNotFound
	}
	return data, err
}

func (this *IngredientM) GetByRecipe(idRecipe int) ([]objects.RecipeIngredient, error) {
	_, err := this.models.Recipes.FindById(idRecipe)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return this.rep.FindByRecipe(idRecipe)
}

func (this *IngredientM) PostToRecipe(id_rcp int, ingArr *[]objects.RecipeIngredient) error {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	for _, obj := range *ingArr {
		err = this.Create(&objects.Ingredient{Title: obj.Ingredient_id})
		if err != nil && err != errors.CategoryExists {
			return err
		}
	}
	
	return this.rep.ReplaceInRecipe(id_rcp, *ingArr)
}

func (this *IngredientM) AddToRecipe(id_rcp int, obj *objects.RecipeIngredient) error {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	err = this.Create(&objects.Ingredient{Title: obj.Ingredient_id})
	if err != nil && err != errors.CategoryExists {
		return err
	}

	return this.rep.AddToRecipe(obj)
}

func (this *IngredientM) DelFromRecipe(id_rcp int, obj *objects.RecipeIngredient) error {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	if this.IsExists(obj.Ingredient_id) == false {
		return errors.UnknownCategory
	}

	return this.rep.DelFromRecipe(id_rcp, obj.Ingredient_id)
}
