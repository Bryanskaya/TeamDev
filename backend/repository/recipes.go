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
	FindByTitle(title string) ([]objects.Recipe, error)
	FindByLogin(login string) ([]objects.Recipe, error)
	FindByCategory(title string) ([]objects.Recipe, error)
	FindById(id int) (*objects.Recipe, error)
	GetLikedByLogin(login string) ([]objects.Recipe, error)
	GetAmountGrades(id int) int
	GetAllCategories() ([]objects.Category, error)

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

func (rep *PGRecipesRep) List() []objects.Recipe {
	temp := []objects.Recipe{}
	rep.db.Find(&temp)
	return temp
}

func (rep *PGRecipesRep) FindByTitle(title string) ([]objects.Recipe, error) {
	temp := []objects.Recipe{}
	err := rep.db.Where("LOWER(title) LIKE LOWER(?)", "%"+title+"%").Find(&temp).Error
	switch err {
	case nil:
		return temp, nil
	case gorm.ErrRecordNotFound:
		return temp, nil
	default:
		return nil, errors.UnknownError
	}
}
func (rep *PGRecipesRep) FindByLogin(login string) ([]objects.Recipe, error) {
	temp := []objects.Recipe{}
	err := rep.db.Where("author = ?", login).Find(&temp).Error
	switch err {
	case nil:
		return temp, nil
	case gorm.ErrRecordNotFound:
		return temp, nil
	default:
		return nil, errors.UnknownError
	}
}

func (rep *PGRecipesRep) FindByCategory(title string) ([]objects.Recipe, error) {
	temp := []objects.Recipe{}
	category := objects.Category{Title: title}

	err := rep.db.Model(&category).Association("Recipes").Find(&temp).Error

	switch err {
	case nil:
		return temp, nil
	case gorm.ErrRecordNotFound:
		return temp, nil
	default:
		return nil, errors.UnknownError
	}
}

func (rep *PGRecipesRep) FindById(id int) (*objects.Recipe, error) {
	temp := new(objects.Recipe)
	err := rep.db.Where("id = ?", id).Find(temp).Error
	switch err {
	case nil:
		return temp, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.UnknownRecipe
	default:
		return nil, errors.UnknownError
	}
}

func (rep *PGRecipesRep) Create(obj *objects.Recipe) error {
	obj.Id = 0
	return rep.db.Create(obj).Error
}
func (rep *PGRecipesRep) CreateList(objArr []objects.Recipe) error {
	for _, obj := range objArr {
		if err := rep.Create(&obj); err != nil {
			return err
		}
	}
	return nil
}

func (rep *PGRecipesRep) Delete(id int) error {
	recipe := objects.Recipe{Id: id}

	err := rep.db.Model(&recipe).Association("Categories").Clear().Error
	if err != nil {
		return err
	}
	err = rep.db.Model(&recipe).Association("Ingredients").Clear().Error
	if err != nil {
		return err
	}
	err = rep.db.Model(&recipe).Association("Grades").Clear().Error
	if err != nil {
		return err
	}

	return rep.db.Where("id = ?", id).Delete(&objects.Recipe{}).Error
}

func (rep *PGRecipesRep) AddGrade(id int, login string) error {
	recipe := objects.Recipe{Id: id}
	return rep.db.Model(&recipe).Association("Grades").Append(&objects.Account{Login: login}).Error
}

func (rep *PGRecipesRep) DeleteGrade(id int, login string) error {
	recipe := objects.Recipe{Id: id}
	return rep.db.Model(&recipe).Association("Grades").Delete(&objects.Account{Login: login}).Error
}

func (rep *PGRecipesRep) GetLikedByLogin(login string) ([]objects.Recipe, error) {
	temp := []objects.Recipe{}
	account := objects.Account{Login: login}

	err := rep.db.Model(&account).Association("Grades").Find(&temp).Error
	return temp, err
}

func (rep *PGRecipesRep) GetAmountGrades(id int) int {
	recipe := objects.Recipe{Id: id}
	return rep.db.Model(&recipe).Association("Grades").Count()
}

func (rep *PGRecipesRep) GetAllCategories() ([]objects.Category, error) {
	temp := []objects.Category{}
	err := rep.db.Find(&temp).Error
	switch err {
	case nil:
		return temp, nil
	case gorm.ErrRecordNotFound:
		return temp, nil
	default:
		return nil, errors.UnknownError
	}
}

func (rep *PGRecipesRep) IsLiked(id_rcp int, login string) bool {
	recipe := objects.Recipe{Id: id_rcp}
	account := objects.Account{Login: login}

	err := rep.db.Model(&recipe).Association("Grades").Find(&account).Error

	return err == nil
}
