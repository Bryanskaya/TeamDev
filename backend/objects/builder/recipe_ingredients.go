package dbuilder

import "api/teamdev/objects"

type RecipeIngredientBuilder struct {
	Recipe_id 		int
	Ingredient_id   string
	Amount 			string
}
func newRecipeIngredientBuilder() *RecipeIngredientBuilder { 
	return new(RecipeIngredientBuilder) 
}
func (this *RecipeIngredientBuilder) Build() *objects.RecipeIngredient { 
	return &objects.RecipeIngredient{Recipe_id: this.Recipe_id, 
		Ingredient_id: this.Ingredient_id, Amount: this.Amount} 
}

type RecipeIngredientMother struct{}
func (RecipeIngredientMother) Obj0() *objects.RecipeIngredient {
	b := newRecipeIngredientBuilder()
	b.Recipe_id = 1
	b.Ingredient_id = "молоко"
	b.Amount = "1 стакан"
	return b.Build()
}
func (RecipeIngredientMother) Obj1() *objects.RecipeIngredient {
	b := newRecipeIngredientBuilder()
	b.Recipe_id = 2
	b.Ingredient_id = "соль"
	b.Amount = "1 щепотка"
	return b.Build()
}
func (RecipeIngredientMother) Obj2() *objects.RecipeIngredient {
	b := newRecipeIngredientBuilder()
	b.Recipe_id = 1
	b.Ingredient_id = "соль"
	b.Amount = "по вкусу"
	return b.Build()
}
func (this RecipeIngredientMother) All() []objects.RecipeIngredient {
	objArr := []objects.RecipeIngredient {
		*this.Obj0(), 
		*this.Obj1(),
		*this.Obj2(),
	}
	return objArr
}
