package models_test

import (
	"api/teamdev/errors"
	"api/teamdev/mocks"
	"api/teamdev/models"
	"api/teamdev/models/tests"
	"api/teamdev/objects"
	"api/teamdev/objects/builder"
	"api/teamdev/repository"
	err "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)


// Find account
func TestFind(t *testing.T) {
	t.Run("account exists", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.AccountMother{}.All()
		objAcc := &objArr[1]

		mockRep := repository.NewAccountsRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewAccount(mockRep, nil)
		res, err := model.Find(objAcc.Login)

		assert.Nil(t, err, "Find has unexpected error")
		assert.Equal(t, objAcc, res, "Find has unexpected result")
	})

	t.Run("account doesn't exists", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.AccountMother{}.All()

		mockRep := repository.NewAccountsRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewAccount(mockRep, nil)
		res, err := model.Find(nWord)

		var nil_ptr *objects.Account = nil
		assert.Equal(t, errors.RecordNotFound, err, "Find has unexpected error")
		assert.Equal(t, nil_ptr, res, "Find has unexpected result")
	})
}

// Get role
func TestGetRole(t *testing.T) {
	t.Run("account exists", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.AccountMother{}.All()
		objAcc := dbuilder.AccountMother{}.Obj0()

		mockRep := repository.NewAccountsRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewAccount(mockRep, nil)
		resRole, err := model.GetRole(objAcc.Login)

		assert.Nil(t, err, "GetRole has unexpected error")
		assert.Equal(t, resRole, objAcc.Role, "GetRole has unexpected error")
	})

	t.Run("account doesn't exists", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.AccountMother{}.All()

		mockRep := repository.NewAccountsRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewAccount(mockRep, nil)
		resRole, err := model.GetRole(nWord)

		assert.Equal(t, errors.RecordNotFound, err, "GetRole has unexpected error")
		assert.Equal(t, "", resRole, "GetRole has unexpected result")
	})
}


func TestLogin(t *testing.T) {
	t.Run("account exists, correct password", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.AccountMother{}.All()
		objAcc := &objArr[0] // not dbuilder.AccountMother{}.Obj0(), hash generated differently every time

		mockRep := repository.NewAccountsRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewAccount(mockRep, nil)
		res, err := model.LogIn(objAcc.Login, objAcc.Login)

		assert.Nil(t, err, "Login has unexpected error")
		assert.Equal(t, res, objAcc, "Login has unexpected result")
	})

	t.Run("account exists, wrong password", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.AccountMother{}.All()
		objAcc := &objArr[0]

		mockRep := repository.NewAccountsRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewAccount(mockRep, nil)
		res, err := model.LogIn(objAcc.Login, nWord)

		var nil_ptr *objects.Account = nil
		assert.Equal(t, errors.WrongPassword, err, "Login has unexpected error")
		assert.Equal(t, nil_ptr, res, "Login has unexpected result")
	})

	t.Run("account doesn't exists, wrong password", func(t *testing.T) {
		db, err := tests.StubConnecton()
		if err != nil {
			panic(err)
		}

		objArr := dbuilder.AccountMother{}.All()

		mockRep := repository.NewAccountsRep(db)
		err = mockRep.CreateList(objArr)
		if err != nil {
			panic(err)
		}

		model := models.NewAccount(mockRep, nil)
		res, err := model.LogIn(nWord, nWord)

		var nil_ptr *objects.Account = nil
		assert.Equal(t, errors.RecordNotFound, err, "Login has unexpected error")
		assert.Equal(t, nil_ptr, res, "Login has unexpected result")
	})

}

// Create account
func TestCreate(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		mockRep := new(mocks.AccountsRep)
		model := models.NewAccount(mockRep, nil)
		obj := dbuilder.AccountMother{}.Obj0()

		mockRep.On("Find", obj.Login).Return(nil, errors.RecordNotFound).Once()
		mockRep.On("Create", obj).Return(nil).Once()

		err := model.Create(obj)

		assert.Nil(t, err, "Create account have unexpected error")
		mockRep.AssertExpectations(t)
	})

	t.Run("already existing account", func(t *testing.T) {
		mockRep := new(mocks.AccountsRep)
		model := models.NewAccount(mockRep, nil)
		obj := dbuilder.AccountMother{}.Obj0()
	
		mockRep.On("Find", obj.Login).Return(obj, nil).Once()
		// mockRep.On("Create", obj).Return().Once()
	
		err := model.Create(obj)
	
		assert.Equal(t, errors.AccountExists, err, "Create account have unexpected error")
		mockRep.AssertExpectations(t)
	})
	
	t.Run("creation failed", func(t *testing.T) {
		mockRep := new(mocks.AccountsRep)
		model := models.NewAccount(mockRep, nil)
		obj := dbuilder.AccountMother{}.Obj0()
	
		mockRep.On("Find", obj.Login).Return(nil, errors.RecordNotFound).Once()
		mockRep.On("Create", obj).Return(err.New("")).Once()
	
		err := model.Create(obj)
	
		assert.Equal(t, errors.DBAdditionError, err, "Create account have unexpected error")
		mockRep.AssertExpectations(t)
	})
}

// Update role by admin
func TestUpdateRole(t *testing.T) {
	t.Run("successful operation", func(t *testing.T) {
		mockRep := new(mocks.AccountsRep)
		model := models.NewAccount(mockRep, nil)
		objCur := dbuilder.AccountMother{}.Obj0()
		objGoal := dbuilder.AccountMother{}.Obj1()
		objUdp := dbuilder.AccountMother{}.Obj1Udp()
	
		mockRep.On("Find", objCur.Login).Return(objCur, nil).Once()
		mockRep.On("Find", objGoal.Login).Return(objGoal, nil).Once()
		mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()
	
		err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)
	
		assert.Nil(t, err, "Update account have unexpected error")
		mockRep.AssertExpectations(t)
	})

	t.Run("admin account wasn't found", func(t *testing.T) {
		mockRep := new(mocks.AccountsRep)
		model := models.NewAccount(mockRep, nil)
		objCur := dbuilder.AccountMother{}.Obj0()
		objGoal := dbuilder.AccountMother{}.Obj1()
		objUdp := dbuilder.AccountMother{}.Obj1Udp()
	
		mockRep.On("Find", objCur.Login).Return(nil, errors.RecordNotFound).Once()
		// mockRep.On("Find", objGoal.Login).Return(objGoal, nil).Once()
		// mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()
	
		err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)
	
		assert.Equal(t, errors.UnknownAccount, err, "Update account have unexpected error")
		mockRep.AssertExpectations(t)
	})

	t.Run("actor is not admin", func(t *testing.T) {
		mockRep := new(mocks.AccountsRep)
		model := models.NewAccount(mockRep, nil)
		objCur := dbuilder.AccountMother{}.Obj1()
		objGoal := dbuilder.AccountMother{}.Obj0()
		objUdp := dbuilder.AccountMother{}.Obj1Udp()
	
		mockRep.On("Find", objCur.Login).Return(objCur, nil).Once()
		// mockRep.On("Find", objGoal.Login).Return(objGoal, nil).Once()
		// mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()
	
		err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)
	
		assert.Equal(t, errors.AccessDenied, err, "Update account have unexpected error")
		mockRep.AssertExpectations(t)
	})

	t.Run("updating account not found", func(t *testing.T) {
		mockRep := new(mocks.AccountsRep)
		model := models.NewAccount(mockRep, nil)
		objCur := dbuilder.AccountMother{}.Obj0()
		objGoal := dbuilder.AccountMother{}.Obj1()
		objUdp := dbuilder.AccountMother{}.Obj1Udp()
	
		mockRep.On("Find", objCur.Login).Return(objCur, nil).Once()
		mockRep.On("Find", objGoal.Login).Return(nil, errors.RecordNotFound).Once()
		// mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()
	
		err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)
	
		assert.Equal(t, errors.UnknownAccount, err, "Update account have unexpected error")
		mockRep.AssertExpectations(t)
	})
}
