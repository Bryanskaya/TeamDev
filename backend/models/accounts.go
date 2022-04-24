package models

import (
	"api/teamdev/errors"
	"api/teamdev/objects"
	"api/teamdev/repository"
	"api/teamdev/utils"
)

const (
	AdminRole  string = "admin"
	UserRole   string = "user"
	UnauthRole string = "unauthorized"
)

type AccountM struct {
	rep    repository.AccountsRep
	models *Models
}

func NewAccount(rep repository.AccountsRep, models *Models) *AccountM {
	return &AccountM{rep, models}
}

func (model *AccountM) Create(obj *objects.Account) error {
	if model.IsExists(obj.Login) {
		return errors.AccountExists
	}

	obj.Role = UserRole
	obj.Salt, obj.HashedPassword = utils.HashPassword(obj.HashedPassword)

	err := model.rep.Create(obj)
	if err != nil {
		return errors.DBAdditionError
	}

	return err
}

func (model *AccountM) UpdateRole(cur_login, login, role string) error {
	cur_role, err := model.GetRole(cur_login)
	if err != nil {
		return errors.UnknownAccount
	}

	if cur_role != AdminRole {
		return errors.AccessDenied
	}

	if !model.IsExists(login) {
		return errors.UnknownAccount
	}

	if role != AdminRole && role != UserRole {
		return errors.UnknownRole
	}

	return model.rep.UpdateRole(login, role)
}

func (model *AccountM) Find(login string) (*objects.Account, error) {
	return model.rep.Find(login)
}

func (model *AccountM) FindLikedRecipe(id_rcp int) ([]objects.Account, error) {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return model.rep.FindLikedRecipe(id_rcp)
}

func (model *AccountM) IsExists(login string) bool {
	_, err := model.Find(login)

	return err == nil
}

func (model *AccountM) GetRole(login string) (role string, err error) {
	acc, err := model.Find(login)

	if err != nil {
		return "", err
	}

	return acc.Role, err
}

func (this *AccountM) GetAll() []objects.Account {
	temp := this.rep.GetAll()
	return temp
}

func (this *AccountM) LogIn(login string, password string) (acc *objects.Account, err error) {
	if acc, err = this.Find(login); err != nil {
		return nil, err
	}
	if !utils.CmpPassword(password, acc.Salt, acc.HashedPassword) {
		return nil, errors.WrongPassword
	}

	return acc, err
}
