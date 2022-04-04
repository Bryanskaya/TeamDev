package objects

import (
	"encoding/json"
	"strings"
)

type Ingredient struct {
	Title string 		`json:"title" gorm:"primary_key"`

	Recipes []*Recipe	`gorm:"many2many:recipe_category;"`
}

func (Ingredient) TableName() string {
	return "ingredients"
}
func (this *Ingredient) ToDTO() *IngredientDTO {
	dto := new(IngredientDTO)
	jsonStr, _ := json.Marshal(this)
	json.Unmarshal(jsonStr, dto)
	return dto
}
func (Ingredient) ArrToDTO(src []Ingredient) []IngredientDTO {
	dst := make([]IngredientDTO, len(src))
	for k, v := range src {
		dst[k] = *v.ToDTO()
	}
	return dst
}

type IngredientDTO struct {
	Title  string `json:"title"`
	Amount string `json:"amount"`
}

func (this *IngredientDTO) ToModel(idRecipe int) *RecipeIngredient {
	mod := new(RecipeIngredient)
	mod.Recipe_id = idRecipe
	mod.Amount = this.Amount
	mod.Ingredient_id = strings.ToLower(this.Title)
	return mod
}

func (IngredientDTO) ToArrModel(idRecipe int, src []IngredientDTO) *[]RecipeIngredient {
	dst := make([]RecipeIngredient, len(src))
	for i, value := range src {
		dst[i] = *value.ToModel(idRecipe)
	}
	return &dst
}