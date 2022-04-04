package repository

import (
	"api/teamdev/errors"
	"api/teamdev/objects"

	"github.com/jinzhu/gorm"
)

type AccountsRep interface {
	Create(obj *objects.Account) error
	CreateList(obj []objects.Account) error

	Find(login string) (*objects.Account, error)
	FindLikedRecipe(id_rcp int) ([]objects.Account, error)
	
	UpdateRole(login, role string) error
}

type PGAccountsRep struct {
	db *gorm.DB
}

func NewAccountsRep(db *gorm.DB) *PGAccountsRep {
	return &PGAccountsRep{db}
}

func (this *PGAccountsRep) Create(obj *objects.Account) error {
	return this.db.Create(obj).Error
}

func (this *PGAccountsRep) CreateList(objArr []objects.Account) error {
	for _, obj := range objArr {
		if err := this.Create(&obj); err != nil {
			return err
		}
	}
	return nil
}

func (this *PGAccountsRep) UpdateRole(login, role string) error {
	return this.db.Model(&objects.Account{}).Where("login = ?", login).Update("role", role).Error
}

func (this *PGAccountsRep) Find(login string) (*objects.Account, error) {
	temp := new(objects.Account)
	err := this.db.Where("login = ?", login).First(temp).Error
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		temp, err = nil, errors.RecordNotFound
		break
	default:
		temp, err = nil, errors.UnknownError
	}

	return temp, err
}

func (this *PGAccountsRep) FindLikedRecipe(id_rcp int) ([]objects.Account, error) {
	temp := []objects.Account{}
	recipe := objects.Recipe{Id: id_rcp}

	err := this.db.Model(&recipe).Association("Grades").Find(&temp).Error

	return temp, err
}