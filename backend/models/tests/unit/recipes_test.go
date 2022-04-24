package models_test

import (
	"api/teamdev/errors"
	"api/teamdev/mocks"
	"api/teamdev/models"
	"api/teamdev/models/tests"
	"api/teamdev/objects"
	"api/teamdev/objects/builder"
	"api/teamdev/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Get all recipes
func TestGetAllRecipes(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.RecipeMother{}.All()

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewRecipe(mockRep, nil)

		resArr := model.GetAll("")

		tests.CompareRecipes(t, resArr, objArr)
	})

	t.Run("empty recipes", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := []objects.Recipe{}

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewRecipe(mockRep, nil)

		resArr := model.GetAll("")

		tests.CompareRecipes(t, resArr, objArr)
	})
}

// Get recipe's author
func TestGetAuthor(t *testing.T) {
	t.Run("author's account exists", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		objRcp := dbuilder.RecipeMother{}.Obj0()
		objAcc := &objAccArr[0]

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		resAuthor, err := allM.Recipes.GetAuthor(objRcp.Id)

		assert.Nil(t, err, "GetAuthor has unexpected error")
		assert.Equal(t, resAuthor, objAcc, "GetAuthor has unexpected error")
	})

	t.Run("recipe doesn't exist", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		// objRcp := dbuilder.RecipeMother{}.Obj0()
		var objAcc *objects.Account = nil

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		resAuthor, err := allM.Recipes.GetAuthor(-1)

		assert.Equal(t, errors.UnknownRecipe, err, "GetAuthor have unexpected error")
		assert.Equal(t, resAuthor, objAcc, "GetAuthor has unexpected error")
	})

	t.Run("author account doesn't exist", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		// objAccArr := dbuilder.AccountMother{}.All()
		objRcp := dbuilder.RecipeMother{}.Obj0()
		var objAcc *objects.Account = nil

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		// err = mockAcc.CreateList(objAccArr)
		// if err != nil { panic(err) }

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		resAuthor, err := allM.Recipes.GetAuthor(objRcp.Id)

		assert.Equal(t, errors.UnknownAccount, err, "GetAuthor have unexpected error")
		assert.Equal(t, resAuthor, objAcc, "GetAuthor has unexpected error")
	})
}

// Find recipes by login
func TestFindRecipesByLogin(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		objAcc := dbuilder.AccountMother{}.Obj0()
		aimArr := []objects.Recipe{objRcpArr[0], objRcpArr[1]}

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		resArr, _ := allM.Recipes.FindByLogin(objAcc.Login)

		tests.CompareRecipes(t, resArr, aimArr)
	})

	t.Run("successful operation, but empty result table", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		// objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		objAcc := dbuilder.AccountMother{}.Obj0()
		aimArr := []objects.Recipe{}

		mockRep := repository.NewRecipesRep(db)
		// err = mockRep.CreateList(objRcpArr)
		// if err != nil { panic(err) }

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		resArr, _ := allM.Recipes.FindByLogin(objAcc.Login)

		tests.CompareRecipes(t, resArr, aimArr)
	})

	t.Run("author account doesn't exist", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		// objAccArr := dbuilder.AccountMother{}.All()
		objAcc := dbuilder.AccountMother{}.Obj0()
		aimArr := []objects.Recipe{}

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		// err = mockAcc.CreateList(objAccArr)
		// if err != nil { panic(err) }

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		resArr, err := allM.Recipes.FindByLogin(objAcc.Login)

		assert.Equal(t, errors.UnknownAccount, err, "Find recipes by login have unexpected error")
		tests.CompareRecipes(t, resArr, aimArr)
	})
}

// Find recipes by its id
func TestFindRecipesById(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		aimRcp := dbuilder.RecipeMother{}.Obj0()

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		model := models.NewRecipe(mockRep, nil)

		resRcp, err := model.FindById(aimRcp.Id)

		assert.Nil(t, err, "Find recipes by its id has unexpected error")
		tests.CompareRecipe(t, *aimRcp, *resRcp)
	})

	t.Run("recipe was not found", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		// aimRcp := dbuilder.RecipeMother{}.Obj0()

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		model := models.NewRecipe(mockRep, nil)

		res, err := model.FindById(-1)

		var nil_ptr *objects.Recipe = nil
		assert.Equal(t, errors.UnknownRecipe, err, "Find recipes by its id has unexpected error")
		assert.Equal(t, nil_ptr, res, "Find recipes by its id has unexpected result")
	})
}

// Get amount of grades
func TestGetAmountGrades(t *testing.T) {
	t.Run("result: 2 likes", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		objRcp := dbuilder.RecipeMother{}.Obj0()
		objAcc1 := dbuilder.AccountMother{}.Obj1()
		objAcc2 := dbuilder.AccountMother{}.Obj2()
		aimNum := 2

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		_ = allM.Recipes.AddGrade(objRcp.Id, objAcc1.Login)
		_ = allM.Recipes.AddGrade(objRcp.Id, objAcc2.Login)

		resNum, err := allM.Recipes.GetAmountGrades(objRcp.Id)

		assert.Nil(t, err, "GetAmountGrades has unexpected error")
		assert.Equal(t, aimNum, resNum, "GetAmountGrades returns wrong answer")
	})

	t.Run("result: 0 likes", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		objRcp := dbuilder.RecipeMother{}.Obj0()
		aimNum := 0

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)
		resNum, err := allM.Recipes.GetAmountGrades(objRcp.Id)

		assert.Nil(t, err, "GetAmountGrades has unexpected error")
		assert.Equal(t, aimNum, resNum, "GetAmountGrades returns wrong answer")
	})

	t.Run("recipe doesn't exist", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		aimNum := 0

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)
		resNum, err := allM.Recipes.GetAmountGrades(-1)

		assert.Equal(t, errors.UnknownRecipe, err, "GetAmountGrades has unexpected error")
		assert.Equal(t, aimNum, resNum, "GetAmountGrades returns wrong answer")
	})
}

// Get liked by login
func TestGetLiked(t *testing.T) {
	t.Run("account exist, result: 1 liked recipe", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		objAcc := dbuilder.AccountMother{}.Obj1()

		aimArr := []objects.Recipe{objRcpArr[0]}

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		for i := 0; i < len(aimArr); i++ {
			_ = allM.Recipes.AddGrade(aimArr[i].Id, objAcc.Login)
		}

		resArr, err := allM.Recipes.GetLikedByLogin(objAcc.Login)

		tests.CompareRecipes(t, resArr, aimArr)
	})

	t.Run("account exist, result: 0 liked recipe", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		objAcc := dbuilder.AccountMother{}.Obj1()

		aimArr := []objects.Recipe{}

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		for i := 0; i < len(aimArr); i++ {
			_ = allM.Recipes.AddGrade(aimArr[i].Id, objAcc.Login)
		}

		resArr, err := allM.Recipes.GetLikedByLogin(objAcc.Login)

		tests.CompareRecipes(t, resArr, aimArr)
	})

	t.Run("account doesn't exist", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objRcpArr := dbuilder.RecipeMother{}.All()
		objAccArr := dbuilder.AccountMother{}.All()
		objAcc := dbuilder.AccountMother{}.Obj1()

		aimArr := []objects.Recipe{}

		mockRep := repository.NewRecipesRep(db)
		err = mockRep.CreateList(objRcpArr)
		if err != nil {
			panic(err)
		}

		mockAcc := repository.NewAccountsRep(db)
		err = mockAcc.CreateList(objAccArr)
		if err != nil {
			panic(err)
		}

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRep, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		for i := 0; i < len(aimArr); i++ {
			_ = allM.Recipes.AddGrade(aimArr[i].Id, objAcc.Login)
		}

		resArr, err := allM.Recipes.GetLikedByLogin(nWord)

		assert.Equal(t, errors.UnknownAccount, err, "GetLikedByLogin has unexpected error")
		tests.CompareRecipes(t, resArr, aimArr)
	})
}

// Create recipe
func TestCreate(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)
		obj := dbuilder.RecipeMother{}.Obj0()
		author := dbuilder.AccountMother{}.Obj0()

		mockAcc.On("Find", obj.Author).Return(author, nil).Once()
		mockRec.On("Create", obj).Return(nil).Once()

		err := allM.Recipes.AddRecipe(obj)

		assert.Nil(t, err, "Create recipe has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})

	t.Run("author doesn't exist", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)
		obj := dbuilder.RecipeMother{}.Obj0()

		mockAcc.On("Find", obj.Author).Return(nil, errors.RecordNotFound).Once()
		// mockRec.On("Create", obj).Return(nil).Once()

		err := allM.Recipes.AddRecipe(obj)

		assert.Equal(t, errors.UnknownAccount, err, "Create recipe have unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})
}

// Add grade
func TestAddGrade(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		obj := dbuilder.RecipeMother{}.Obj0()
		acc := dbuilder.AccountMother{}.Obj0()

		mockRec.On("FindById", obj.Id).Return(obj, nil).Once()
		mockAcc.On("Find", acc.Login).Return(acc, nil).Once()
		mockRec.On("AddGrade", obj.Id, acc.Login).Return(nil).Once()

		err := allM.Recipes.AddGrade(obj.Id, acc.Login)

		assert.Nil(t, err, "Add grade has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})

	t.Run("recipe doesn't exist", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		obj := dbuilder.RecipeMother{}.Obj0()
		acc := dbuilder.AccountMother{}.Obj0()

		mockRec.On("FindById", obj.Id).Return(nil, errors.RecordNotFound).Once()

		err := allM.Recipes.AddGrade(obj.Id, acc.Login)

		assert.Equal(t, errors.UnknownRecipe, err, "Add grade has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})

	t.Run("account doesn't exist", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		obj := dbuilder.RecipeMother{}.Obj0()
		acc := dbuilder.AccountMother{}.Obj0()

		mockRec.On("FindById", obj.Id).Return(obj, nil).Once()
		mockAcc.On("Find", acc.Login).Return(nil, errors.RecordNotFound).Once()

		err := allM.Recipes.AddGrade(obj.Id, acc.Login)

		assert.Equal(t, errors.UnknownAccount, err, "Add grade has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})
}

// Delete grade
func TestDelGradeRecipe(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		objRcp := dbuilder.RecipeMother{}.Obj0()
		objAcc := dbuilder.AccountMother{}.Obj0()

		mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
		mockAcc.On("Find", objAcc.Login).Return(objAcc, nil).Once()
		mockRec.On("DeleteGrade", objRcp.Id, objAcc.Login).Return(nil).Once()

		err := allM.Recipes.DeleteGrade(objRcp.Id, objAcc.Login)

		assert.Nil(t, err, "Deletion a grade from recipe has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})

	t.Run("recipe doesn't exist", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		objRcp := dbuilder.RecipeMother{}.Obj0()
		objAcc := dbuilder.AccountMother{}.Obj0()

		mockRec.On("FindById", objRcp.Id).Return(nil, errors.RecordNotFound).Once()

		err := allM.Recipes.DeleteGrade(objRcp.Id, objAcc.Login)

		assert.Equal(t, errors.UnknownRecipe, err, "Deletion a grade from recipe has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})

	t.Run("account doesn't exist", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		objRcp := dbuilder.RecipeMother{}.Obj0()
		objAcc := dbuilder.AccountMother{}.Obj0()

		mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
		mockAcc.On("Find", objAcc.Login).Return(nil, errors.RecordNotFound).Once()

		err := allM.Recipes.DeleteGrade(objRcp.Id, objAcc.Login)

		assert.Equal(t, errors.UnknownAccount, err, "Deletion a grade from recipe has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})
}

// Delete recipe
func TestDelRecipe(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)

		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)

		objRcp := dbuilder.RecipeMother{}.Obj2()
		objAdmin := dbuilder.AccountMother{}.Obj0()
		objAuthor := dbuilder.AccountMother{}.Obj1()

		mockAcc.On("Find", objAdmin.Login).Return(objAdmin, nil).Once()
		mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
		mockAcc.On("Find", objRcp.Author).Return(objAuthor, nil).Once()
		mockRec.On("Delete", objRcp.Id).Return(nil).Once()

		err := allM.Recipes.DeleteRecipe(objRcp.Id, objAdmin.Login)

		assert.Nil(t, err, "Deletion a recipe has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})

	t.Run("recipe doesn't exist", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)
	
		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)
	
		objRcp := dbuilder.RecipeMother{}.Obj2()
		objAdmin := dbuilder.AccountMother{}.Obj0()
		// objAuthor := dbuilder.AccountMother{}.Obj1()
	
		mockAcc.On("Find", objAdmin.Login).Return(objAdmin, nil).Once()
		mockRec.On("FindById", objRcp.Id).Return(nil, errors.RecordNotFound).Once()
	
		// mockAcc.On("Find", objRcp.Author).Return(objAuthor, nil).Once()
		// mockRec.On("Delete", objRcp.Id).Return(nil).Once()
	
		err := allM.Recipes.DeleteRecipe(objRcp.Id, objAdmin.Login)
	
		assert.Equal(t, errors.UnknownRecipe, err, "Deletion a recipe has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})

	t.Run("account doesn't exist", func(t *testing.T) {
		mockRec := new(mocks.RecipesRep)
		mockAcc := new(mocks.AccountsRep)
	
		allM := new(models.Models)
		allM.Recipes = models.NewRecipe(mockRec, allM)
		allM.Accounts = models.NewAccount(mockAcc, allM)
	
		objRcp := dbuilder.RecipeMother{}.Obj2()
		objAdmin := dbuilder.AccountMother{}.Obj0()
	
		mockAcc.On("Find", objAdmin.Login).Return(nil, errors.RecordNotFound).Once()
	
		err := allM.Recipes.DeleteRecipe(objRcp.Id, objAdmin.Login)
	
		assert.Equal(t, errors.RecordNotFound, err, "Deletion a recipe has unexpected error")
		mockRec.AssertExpectations(t)
		mockAcc.AssertExpectations(t)
	})
}
