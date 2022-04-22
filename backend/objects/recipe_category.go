package objects

type RecipeCategory struct {
	Recipe_id      int    `gorm:"foreignkey:Id" json:"recipe"`
	Category_title string `gorm:"foreignkey:title" json:"category"`
}

func (RecipeCategory) TableName() string {
	return "recipe_category"
}
