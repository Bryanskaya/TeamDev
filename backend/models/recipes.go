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

func (this *RecipeM) GetAll(title string) []objects.Recipe {
	temp, _ := this.rep.FindByTitle(title)
	return temp
}

func (this *RecipeM) GetAllCategories() []objects.Category {
	temp, _ := this.rep.GetAllCategories()
	return temp
}

func (this *RecipeM) GetAuthor(id int) (*objects.Account, error) {
	rcp, err := this.FindById(id)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	login := rcp.Author
	acc, err := this.models.Accounts.Find(login)
	if err != nil {
		err = errors.UnknownAccount
	}
	return acc, err
}

func (this *RecipeM) FindByLogin(login string) ([]objects.Recipe, error) {
	if !this.models.Accounts.IsExists(login) {
		return nil, errors.UnknownAccount
	}

	return this.rep.FindByLogin(login)
}

func (this *RecipeM) FindByCategory(title string) ([]objects.Recipe, error) {
	return this.rep.FindByCategory(title)
}

func (this *RecipeM) FindById(id int) (*objects.Recipe, error) {
	return this.rep.FindById(id)
}

func (this *RecipeM) AddRecipe(obj *objects.Recipe) error {
	_, err := this.models.Accounts.Find(obj.Author)
	if err != nil {
		return errors.UnknownAccount
	}

	// TODO: validate other data
	err = this.rep.Create(obj)
	return err
}

func (obj *RecipeM) AddGrade(id int, login string) error {
	_, err := obj.models.Recipes.FindById(id)
	if err != nil {
		return errors.UnknownRecipe
	}

	if obj.models.Accounts.IsExists(login) == false {
		return errors.UnknownAccount
	}

	return obj.rep.AddGrade(id, login)
}

func (this *RecipeM) DeleteGrade(id_rcp int, login string) error {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	if this.models.Accounts.IsExists(login) == false {
		return errors.UnknownAccount
	}

	return this.rep.DeleteGrade(id_rcp, login)
}

func (this *RecipeM) GetAmountGrades(id_rcp int) (int, error) {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return 0, errors.UnknownRecipe
	}

	return this.rep.GetAmountGrades(id_rcp), nil
}

func (this *RecipeM) GetLikedByLogin(login string) ([]objects.Recipe, error) {
	if this.models.Accounts.IsExists(login) == false {
		return nil, errors.UnknownAccount
	}

	return this.rep.GetLikedByLogin(login)
}

func (this *RecipeM) DeleteRecipe(id int, login string) (err error) {
	userRole, err := this.models.Accounts.GetRole(login)
	if err != nil {
		return err
	}

	author, err := this.GetAuthor(id)
	if err != nil {
		return err
	}

	if userRole == AdminRole || login == author.Login {
		err = this.rep.Delete(id)
	} else {
		err = errors.AccessDeleteDenied
	}

	return err
}

func (this *RecipeM) IsLiked(id_rcp int, login string) (res bool, err error) {
	res = false

	if this.models.Accounts.IsExists(login) == false {
		return res, errors.UnknownAccount
	}

	_, err = this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return res, errors.UnknownRecipe
	}

	res = this.rep.IsLiked(id_rcp, login)

	return res, nil
}
