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

func (this *StepM) GetSteps(id_rcp int) ([]objects.Step, error) {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return this.rep.FindSteps(id_rcp)
}

func (this *StepM) GetStepByNum(id_rcp, step int) (*objects.Step, error) {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	data, err := this.rep.FindStepByNum(id_rcp, step)
	if err != nil {
		err = errors.UnknownStep
	}

	return &data, err
}

func (this *StepM) GetStepLast(id_rcp int) (*objects.Step, error) {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	data := this.rep.FindStepLast(id_rcp)

	return &data, err
}

func (this *StepM) AddStep(id_rcp int, obj *objects.Step, login string) error {
	cur_acc, err := this.models.Accounts.Find(login)
	if err != nil {
		return errors.UnknownAccount
	}

	if cur_acc.Role != AdminRole {
		auth_acc, err := this.models.Recipes.GetAuthor(id_rcp)
		if err != nil {
			return err
		}

		if auth_acc.Login != login {
			return errors.AccessDenied
		}
	}

	_, err = this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	last := this.rep.FindStepLast(id_rcp)

	obj.Recipe = id_rcp
	obj.Num = last.Num + 1

	return this.rep.Create(obj)
}

func (this *StepM) DeleteStep(id_rcp, step int, login string) error {
	userRole, err := this.models.Accounts.GetRole(login)
	if err != nil {
		return err
	}

	author, err := this.models.Recipes.GetAuthor(id_rcp)
	if err != nil {
		return err
	}

	if userRole == AdminRole || login == author.Login {
		_, err = this.GetStepByNum(id_rcp, step)
		if err != nil {
			return errors.UnknownStep
		}

		err = this.rep.Delete(id_rcp, step)
	} else {
		err = errors.AccessDeleteDenied
	}

	return err
}

func (this *StepM) UpdateStep(cur_login string, id_rcp int, step int, obj *objects.Step) error {
	if this.models.Accounts.IsExists(cur_login) == false {
		return errors.UnknownAccount
	}

	cur_role, err := this.models.Accounts.GetRole(cur_login)
	if err != nil {
		return errors.UnknownAccount
	}

	author, err := this.models.Recipes.GetAuthor(id_rcp)
	if err != nil {
		return err
	}

	if cur_role == AdminRole || cur_login == author.Login {
		_, err = this.GetStepByNum(id_rcp, step)
		if err != nil {
			return errors.UnknownStep
		}

		err = this.rep.UpdateStep(id_rcp, step, obj)
	} else {
		err = errors.AccessDeleteDenied
	}

	return err
}
