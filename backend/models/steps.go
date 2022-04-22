package models

import (
	"api/teamdev/errors"
	"api/teamdev/objects"
	"api/teamdev/repository"
)

type StepM struct {
	rep    repository.StepsRep
	models *Models
}

func NewStep(rep repository.StepsRep, models *Models) *StepM {
	return &StepM{rep, models}
}

func (model *StepM) GetSteps(id_rcp int) ([]objects.Step, error) {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return model.rep.FindSteps(id_rcp)
}

func (model *StepM) GetStepByNum(id_rcp, step int) (*objects.Step, error) {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	data, err := model.rep.FindStepByNum(id_rcp, step)
	if err != nil {
		err = errors.UnknownStep
	}

	return &data, err
}

func (model *StepM) GetStepLast(id_rcp int) (*objects.Step, error) {
	_, err := model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	data := model.rep.FindStepLast(id_rcp)

	return &data, err
}

func (model *StepM) AddStep(id_rcp int, obj *objects.Step, login string) error {
	cur_acc, err := model.models.Accounts.Find(login)
	if err != nil {
		return errors.UnknownAccount
	}

	if cur_acc.Role != AdminRole {
		auth_acc, err := model.models.Recipes.GetAuthor(id_rcp)
		if err != nil {
			return err
		}

		if auth_acc.Login != login {
			return errors.AccessDenied
		}
	}

	_, err = model.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	last := model.rep.FindStepLast(id_rcp)

	obj.Recipe = id_rcp
	obj.Num = last.Num + 1

	return model.rep.Create(obj)
}

func (model *StepM) DeleteStep(id_rcp, step int, login string) error {
	userRole, err := model.models.Accounts.GetRole(login)
	if err != nil {
		return err
	}

	author, err := model.models.Recipes.GetAuthor(id_rcp)
	if err != nil {
		return err
	}

	if userRole == AdminRole || login == author.Login {
		_, err = model.GetStepByNum(id_rcp, step)
		if err != nil {
			return errors.UnknownStep
		}

		err = model.rep.Delete(id_rcp, step)
	} else {
		err = errors.AccessDeleteDenied
	}

	return err
}

func (model *StepM) UpdateStep(cur_login string, id_rcp int, step int, obj *objects.Step) error {
	if !model.models.Accounts.IsExists(cur_login) {
		return errors.UnknownAccount
	}

	cur_role, err := model.models.Accounts.GetRole(cur_login)
	if err != nil {
		return errors.UnknownAccount
	}

	author, err := model.models.Recipes.GetAuthor(id_rcp)
	if err != nil {
		return err
	}

	if cur_role == AdminRole || cur_login == author.Login {
		_, err = model.GetStepByNum(id_rcp, step)
		if err != nil {
			return errors.UnknownStep
		}

		err = model.rep.UpdateStep(id_rcp, step, obj)
	} else {
		err = errors.AccessDeleteDenied
	}

	return err
}
