package repository

import (
	"api/teamdev/errors"
	"api/teamdev/objects"

	"github.com/jinzhu/gorm"
)

type RecipesRep interface {
	Create(rcp *objects.Recipe) error
	CreateList(objArr []objects.Recipe) error

	List() []objects.Recipe
	FindByLogin(login string) ([]objects.Recipe, error)
	FindById(id int) (*objects.Recipe, error)
	GetLikedByLogin(login string) ([]objects.Recipe, error)
	GetAmountGrades(id int) int
	
	Delete(id int) error
	AddGrade(id int, login string) error
	DeleteGrade(id int, login string) error

	IsLiked(id_rcp int, login string) bool
}

type PGRecipesRep struct {
	db *gorm.DB
}

func NewRecipesRep(db *gorm.DB) *PGRecipesRep {
	return &PGRecipesRep{db}
}

func (this *PGRecipesRep) List() []objects.Recipe {
	temp := []objects.Recipe{}
	this.db.Find(&temp)
	return temp
}

func (this *PGRecipesRep) FindByLogin(login string)  ([]objects.Recipe, error) {
	temp := []objects.Recipe{}
	err := this.db.Where("author = ?", login).Find(&temp).Error
	switch err {
	case nil:
		return temp, nil
	case gorm.ErrRecordNotFound:
		return temp, nil
	default:
		return nil, errors.UnknownError
	}
}

func (this *PGRecipesRep) FindById(id int) (*objects.Recipe, error) {
	temp := new(objects.Recipe)
	err := this.db.Where("id = ?", id).Find(temp).Error
	switch err {
	case nil:
		return temp, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.UnknownRecipe
	default:
		return nil, errors.UnknownError
	}
}

func (this *PGRecipesRep) Create(obj *objects.Recipe) error {
	obj.Id = 0
	return this.db.Create(obj).Error
}
func (this *PGRecipesRep) CreateList(objArr []objects.Recipe) error {
	for _, obj := range objArr {
		if err := this.Create(&obj); err != nil {
			return err
		}
	}
	return nil
}

func (this *PGRecipesRep) Delete(id int) error {
	recipe := objects.Recipe{Id: id}

	err := this.db.Model(&recipe).Association("Categories").Clear().Error
	if err != nil { return err }
	err = this.db.Model(&recipe).Association("Ingredients").Clear().Error
	if err != nil { return err }
	err = this.db.Model(&recipe).Association("Grades").Clear().Error
	if err != nil { return err }
	
	return this.db.Where("id = ?", id).Delete(&objects.Recipe{}).Error
}

func (this *PGRecipesRep) AddGrade(id int, login string) error {
	recipe := objects.Recipe{Id: id}
	return this.db.Model(&recipe).Association("Grades").Append(&objects.Account{Login: login}).Error
}

func (this *PGRecipesRep) DeleteGrade(id int, login string) error {
	recipe := objects.Recipe{Id: id}
	return this.db.Model(&recipe).Association("Grades").Delete(&objects.Account{Login: login}).Error
}

func (this *PGRecipesRep) GetLikedByLogin(login string) ([]objects.Recipe, error) {
	temp := []objects.Recipe{}
	account := objects.Account{Login: login}

	err := this.db.Model(&account).Association("Grades").Find(&temp).Error
	return temp, err
}

func (this *PGRecipesRep) GetAmountGrades(id int) int {
	recipe := objects.Recipe{Id: id}
	return this.db.Model(&recipe).Association("Grades").Count()
}

func (this *PGRecipesRep) IsLiked(id_rcp int, login string) bool {
	recipe := objects.Recipe{Id: id_rcp}
	account := objects.Account{Login: login}

	err := this.db.Model(&recipe).Association("Grades").Find(&account).Error

	return err == nil
}
