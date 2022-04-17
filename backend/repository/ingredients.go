package repository

import (
	"api/teamdev/objects"
	"strings"

	"github.com/jinzhu/gorm"
)

type IngredientsRep interface {
	Create(obj *objects.Ingredient) error
	CreateList(obj []objects.Ingredient) error

	Get(title string) (*objects.Ingredient, error)
	FindByRecipe(idRecipe int) ([]objects.RecipeIngredient, error)

	AddToRecipe(obj *objects.RecipeIngredient) error
	ReplaceInRecipe(id_rcp int, arr []objects.RecipeIngredient) error
	DelFromRecipe(idRecipe int, title string) error
}

type PGIngredientsRep struct {
	db *gorm.DB
}

func NewIngredientsRep(db *gorm.DB) *PGIngredientsRep {
	return &PGIngredientsRep{db}
}


func (rep *PGIngredientsRep) Create(obj *objects.Ingredient) error {
	return rep.db.Create(obj).Error
}

func (rep *PGIngredientsRep) CreateList(objArr []objects.Ingredient) error {
	for _, obj := range objArr {
		if err := rep.Create(&obj); err != nil {
			return err
		}
	}
	return nil
}

func (rep *PGIngredientsRep) Get(title string) (*objects.Ingredient, error) {
	temp := new(objects.Ingredient)
	err := rep.db.Where("LOWER(title) = ?", strings.ToLower(title)).Find(temp).Error
	return temp, err
}

func (rep *PGIngredientsRep) FindByRecipe(idRecipe int) ([]objects.RecipeIngredient, error) {
	temp := []objects.RecipeIngredient{}
	err := rep.db.Where(objects.RecipeIngredient{Recipe_id: idRecipe}).Find(&temp).Error
	return temp, err
}

func (rep *PGIngredientsRep) AddToRecipe(obj *objects.RecipeIngredient) error {
	return rep.db.Create(&obj).Error
}

func (rep *PGIngredientsRep) ReplaceInRecipe(id_rcp int, arr []objects.RecipeIngredient) error {
	recipe := objects.Recipe{Id: id_rcp}
	return rep.db.Model(&recipe).Association("RecipeIngredient").Replace(arr).Error
}

func (rep *PGIngredientsRep) DelFromRecipe(idRecipe int, title string) error {
	return rep.db.Where("recipe_id = ? AND ingredient_id = ?", idRecipe, title).Delete(&objects.RecipeIngredient{}).Error
}
