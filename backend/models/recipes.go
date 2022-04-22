package models

import (
	"api/teamdev/errors"
	"api/teamdev/objects"
	"api/teamdev/repository"
)

type RecipeM struct {
	rep    repository.RecipesRep
	models *Models
}

func NewRecipe(rep repository.RecipesRep, models *Models) *RecipeM {
	return &RecipeM{rep, models}
}

func (model *RecipeM) GetAll(title string) []objects.Recipe {
	temp, _ := model.rep.FindByTitle(title)
	return temp
}

func (model *RecipeM) GetAllCategories() []objects.Category {
	temp, _ := model.rep.GetAllCategories()
	return temp
}

func (model *RecipeM) GetAuthor(id int) (*objects.Account, error) {
	rcp, err := model.FindById(id)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	login := rcp.Author
	acc, err := model.models.Accounts.Find(login)
	if err != nil {
		err = errors.UnknownAccount
	}
	return acc, err
}

func (model *RecipeM) FindByLogin(login string) ([]objects.Recipe, error) {
	if !model.models.Accounts.IsExists(login) {
		return nil, errors.UnknownAccount
	}

	return model.rep.FindByLogin(login)
}

func (model *RecipeM) FindByCategory(title string) ([]objects.Recipe, error) {
	return model.rep.FindByCategory(title)
}

func (model *RecipeM) FindById(id int) (*objects.Recipe, error) {
	return model.rep.FindById(id)
}

func (model *RecipeM) AddRecipe(obj *objects.Recipe) error {
	_, err := model.models.Accounts.Find(obj.Author)
	if err != nil {
		return errors.UnknownAccount
	}

	// TODO: validate other data
	err = model.rep.Create(obj)
	return err
}

func (obj *RecipeM) AddGrade(id int, login string) error {
	_, err := obj.models.Recipes.FindById(id)
	if err != nil {
		return errors.UnknownRecipe
	}

	if !obj.models.Accounts.IsExists(login) {
		return errors.UnknownAccount
	}

	return obj.rep.AddGrade(id, login)
}

func (model *RecipeM) DeleteGrade(id_rcp int, login string) error {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	if !model.models.Accounts.IsExists(login) {
		return errors.UnknownAccount
	}

	return model.rep.DeleteGrade(id_rcp, login)
}

func (model *RecipeM) GetAmountGrades(id_rcp int) (int, error) {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return 0, errors.UnknownRecipe
	}

	return model.rep.GetAmountGrades(id_rcp), nil
}

func (model *RecipeM) GetLikedByLogin(login string) ([]objects.Recipe, error) {
	if !model.models.Accounts.IsExists(login) {
		return nil, errors.UnknownAccount
	}

	return model.rep.GetLikedByLogin(login)
}

func (model *RecipeM) DeleteRecipe(id int, login string) (err error) {
	userRole, err := model.models.Accounts.GetRole(login)
	if err != nil {
		return err
	}

	author, err := model.GetAuthor(id)
	if err != nil {
		return err
	}

	if userRole == AdminRole || login == author.Login {
		err = model.rep.Delete(id)
	} else {
		err = errors.AccessDeleteDenied
	}

	return err
}

func (model *RecipeM) IsLiked(id_rcp int, login string) (res bool, err error) {
	res = false

	if !model.models.Accounts.IsExists(login) {
		return res, errors.UnknownAccount
	}

	_, err = model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return res, errors.UnknownRecipe
	}

	res = model.rep.IsLiked(id_rcp, login)

	return res, nil
}
