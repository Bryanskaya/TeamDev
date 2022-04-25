package repository

import (
	"api/teamdev/objects"
	"strings"

	"github.com/jinzhu/gorm"
)

type CategoriesRep interface {
	Create(obj *objects.Category) error
	CreateList(obj []objects.Category) error

	List() []objects.Category
	Find(ctg string) ([]objects.Category, error)
	Get(ctg string) (*objects.Category, error)
	FindRecipes(ctg string) ([]objects.Recipe, error)
	FindByRecipe(id_rcp int) ([]objects.Category, error)

	AddToRecipe(id_rcp int, ctg string) error
	ReplaceInRecipe(id_rcp int, arr []objects.Category) error
	DelFromRecipe(id_rcp int, ctg string) error
}

type PGCategoriesRep struct {
	db *gorm.DB
}

func NewCategotiesRep(db *gorm.DB) *PGCategoriesRep {
	return &PGCategoriesRep{db}
}

func (this *PGCategoriesRep) Create(obj *objects.Category) error {
	return this.db.Create(obj).Error
}

func (this *PGCategoriesRep) CreateList(objArr []objects.Category) error {
	for _, obj := range objArr {
		if err := this.Create(&obj); err != nil {
			return err
		}
	}
	return nil
}

func (this *PGCategoriesRep) List() []objects.Category {
	temp := []objects.Category{}
	this.db.Find(&temp)
	return temp
}

func (this *PGCategoriesRep) Find(ctg string) ([]objects.Category, error) {
	temp := []objects.Category{}
	err := this.db.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(ctg)+"%").Find(&temp).Error
	return temp, err
}

func (this *PGCategoriesRep) FindRecipes(ctg string) ([]objects.Recipe, error) {
	temp := []objects.Recipe{}
	category := objects.Category{Title: ctg}

	err := this.db.Model(&category).Association("Recipes").Find(&temp).Error

	return temp, err
}

func (this *PGCategoriesRep) FindByRecipe(id_rcp int) ([]objects.Category, error) {
	temp := []objects.Category{}
	recipe := objects.Recipe{Id: id_rcp}

	err := this.db.Model(&recipe).Association("Categories").Find(&temp).Error

	return temp, err
}

func (this *PGCategoriesRep) Get(ctg string) (*objects.Category, error) {
	temp := new(objects.Category)
	err := this.db.Where("LOWER(title) = ?", strings.ToLower(ctg)).Find(temp).Error
	return temp, err
}

func (this *PGCategoriesRep) AddToRecipe(id_rcp int, ctg string) error {
	recipe := objects.Recipe{Id: id_rcp}
	return this.db.Model(&recipe).Association("Categories").Append(&objects.Category{Title: ctg}).Error
}

func (this *PGCategoriesRep) ReplaceInRecipe(id_rcp int, arr []objects.Category) error {
	recipe := objects.Recipe{Id: id_rcp}
	return this.db.Model(&recipe).Association("Categories").Replace(arr).Error
}

func (this *PGCategoriesRep) DelFromRecipe(id_rcp int, ctg string) error {
	recipe := objects.Recipe{Id: id_rcp}
	return this.db.Model(&recipe).Association("Categories").Delete(&objects.Category{Title: ctg}).Error
}
