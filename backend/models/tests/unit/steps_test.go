package models_test

import (
	"api/teamdev/mocks"
	"api/teamdev/models"
	"api/teamdev/models/tests"
	"api/teamdev/objects"
	"api/teamdev/objects/builder"
	"api/teamdev/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Get all recipe's steps
func TestGetSteps(t *testing.T) {
	t.Run("recipe exists + 2 steps", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objStpArr := dbuilder.StepMother{}.All()
		objRcpArr := dbuilder.RecipeMother{}.All()
		objRcp := dbuilder.RecipeMother{}.Obj0()
		aimArr := []objects.Step{objStpArr[0], objStpArr[1]}

		mockStp := repository.NewStepsRep(db)
		err = mockStp.CreateList(objStpArr)
		if err != nil {
			panic(err)
		}

		mockRcp := repository.NewRecipesRep(db)
		err = mockRcp.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRcp, allM)
		allM.Steps = models.NewStep(mockStp, allM)

		resArr, err := allM.Steps.GetSteps(objRcp.Id)

		assert.Nil(t, err, "GetSteps has unexpected error")
		assert.ElementsMatch(t, aimArr, resArr, "GetSteps has unexpected error")
	})
}

// Get step by its number and recipe's id
func TestGetStepByNum(t *testing.T) {
	t.Run("recipe exists + 2 steps (in test #2)", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objStpArr := dbuilder.StepMother{}.All()
		objRcpArr := dbuilder.RecipeMother{}.All()
		objRcp := dbuilder.RecipeMother{}.Obj0()
		aimStp := dbuilder.StepMother{}.Obj1()

		mockStp := repository.NewStepsRep(db)
		err = mockStp.CreateList(objStpArr)
		if err != nil {
			panic(err)
		}

		mockRcp := repository.NewRecipesRep(db)
		err = mockRcp.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRcp, allM)
		allM.Steps = models.NewStep(mockStp, allM)

		resStp, err := allM.Steps.GetStepByNum(objRcp.Id, aimStp.Num)

		assert.Nil(t, err, "GetStepByNum has unexpected error")
		assert.Equal(t, aimStp, resStp, "GetStepByNum has unexpected error")
	})
}

// Get last step
func TestGetStepLast(t *testing.T) {
	t.Run("sucсess", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.StepMother{}.All()
		objRcp := dbuilder.RecipeMother{}.Obj0()
		aimStp := dbuilder.StepMother{}.Obj1()

		mockStp := repository.NewStepsRep(db)
		err = mockStp.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		mockRcp := repository.NewRecipesRep(db)
		err = mockRcp.Create(objRcp)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRcp, allM)
		allM.Steps = models.NewStep(mockStp, allM)

		resArr, err := allM.Steps.GetStepLast(objRcp.Id)

		assert.Nil(t, err, "GetStepLast has unexpected error")
		assert.Equal(t, aimStp, resArr, "GetStepLast has unexpected error")
	})
}

// Create step
func TestCreateStep(t *testing.T) {
	t.Run("by admin - successful operation", func(t *testing.T) {
		mockStp := new(mocks.StepsRep)
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Steps = models.NewStep(mockStp, allM)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		objAcc := dbuilder.AccountMother{}.Obj0()
		objRec := dbuilder.RecipeMother{}.Obj0()
		obj := dbuilder.StepMother{}.Obj0()

		mockAcc.On("Find", objAcc.Login).Return(objAcc, nil).Once()
		mockRec.On("FindById", objRec.Id).Return(objRec, nil).Once()
		mockStp.On("Create", obj).Return(nil).Once()
		mockStp.On("FindStepLast", objRec.Id).Return(*obj).Once()

		err := allM.Steps.AddStep(objRec.Id, obj, objAcc.Login)

		assert.Nil(t, err, "Create step have unexpected error")
		mockRec.AssertExpectations(t)
	})

	t.Run("by recipe author - successful operation", func(t *testing.T) {
		mockStp := new(mocks.StepsRep)
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Steps = models.NewStep(mockStp, allM)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		objAcc := dbuilder.AccountMother{}.Obj1()
		objRec := dbuilder.RecipeMother{}.Obj2()
		obj := dbuilder.StepMother{}.Obj2()

		mockAcc.On("Find", objAcc.Login).Return(objAcc, nil).Twice()
		mockRec.On("FindById", objRec.Id).Return(objRec, nil).Twice()
		mockStp.On("Create", obj).Return(nil).Once()
		mockStp.On("FindStepLast", objRec.Id).Return(*obj).Once()

		err := allM.Steps.AddStep(objRec.Id, obj, objAcc.Login)

		assert.Nil(t, err, "Create step have unexpected error")
		mockRec.AssertExpectations(t)
	})
}

// Delete step
func TestDeleteStep(t *testing.T) {
	t.Run("by admin - successful operation", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)
		mockStp := new(mocks.StepsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)
		allM.Steps = models.NewStep(mockStp, allM)

		objRcp := dbuilder.RecipeMother{}.Obj2()
		objAdmin := dbuilder.AccountMother{}.Obj0()
		objAuthor := dbuilder.AccountMother{}.Obj1()
		objStep := dbuilder.StepMother{}.Obj2()

		mockAcc.On("Find", objAdmin.Login).Return(objAdmin, nil).Once()

		mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
		mockAcc.On("Find", objRcp.Author).Return(objAuthor, nil).Once()

		mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
		mockStp.On("FindStepByNum", objRcp.Id, objStep.Num).Return(*objStep, nil).Once()

		mockStp.On("Delete", objRcp.Id, objStep.Num).Return(nil).Once()

		err := allM.Steps.DeleteStep(objRcp.Id, objStep.Num, objAdmin.Login)

		assert.Nil(t, err, "Deletion a step has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
		mockStp.AssertExpectations(t)
	})
}

// Update step
func TestUpdateStep(t *testing.T) {
	t.Run("by admin - successful operation", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)
		mockStp := new(mocks.StepsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)
		allM.Steps = models.NewStep(mockStp, allM)

		objRcp := dbuilder.RecipeMother{}.Obj2()
		objAdmin := dbuilder.AccountMother{}.Obj0()
		objAuthor := dbuilder.AccountMother{}.Obj1()
		objStep := dbuilder.StepMother{}.Obj2()

		mockAcc.On("Find", objAdmin.Login).Return(objAdmin, nil).Twice()

		mockAcc.On("Find", objRcp.Author).Return(objAuthor, nil).Once()
		mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

		mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
		mockStp.On("FindStepByNum", objRcp.Id, objStep.Num).Return(*objStep, nil).Once()

		mockStp.On("UpdateStep", objRcp.Id, objStep.Num, objStep).Return(nil).Once()

		err := allM.Steps.UpdateStep(objAdmin.Login, objRcp.Id, objStep.Num, objStep)

		assert.Nil(t, err, "Update step has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
		mockStp.AssertExpectations(t)
	})
}
