package objects

import (
	"encoding/json"
	"time"
)

type Recipe struct {
	Id          int       `json:"id" gorm:"primary_key"`
	Author      string    `json:"author" gorm:"foreignkey:Login"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	PortionNum  int       `json:"portion_num"`
	PicUrl      string    `json:"pic_url"`

	Categories       []*Category   `gorm:"many2many:recipe_category;"`
	Ingredients      []*Ingredient `gorm:"many2many:recipe_ingredient;"`
	RecipeIngredient []*RecipeIngredient
	Grades           []*Account `gorm:"many2many:grades;"`
}

func (Recipe) TableName() string {
	return "recipes"
}

func (obj *Recipe) ToDTO() *RecipeDTO {
	dto := new(RecipeDTO)
	jsonStr, _ := json.Marshal(obj)
	json.Unmarshal(jsonStr, dto)
	return dto
}

func (Recipe) ArrToDTO(src []Recipe) []RecipeDTO {
	dst := make([]RecipeDTO, len(src))
	for k, v := range src {
		dst[k] = *v.ToDTO()
	}
	return dst
}

type RecipeDTO struct {
	Id          int    `json:"id"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	PortionNum  int    `json:"portion_num"`
	PicUrl      string `json:"pic_url"`
}

func (obj *RecipeDTO) ToModel() *Recipe {
	mod := new(Recipe)
	if obj.CreatedAt == "" {
		obj.CreatedAt = "0001-01-01T00:00:00.000Z"
	}

	jsonStr, _ := json.Marshal(obj)
	json.Unmarshal(jsonStr, mod)
	return mod
}
