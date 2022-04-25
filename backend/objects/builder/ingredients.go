package dbuilder

import "api/teamdev/objects"

type IngredientBuilder struct {
	Title string
}

func newIngredientBuilder() *IngredientBuilder {
	return new(IngredientBuilder)
}
func (this *IngredientBuilder) Build() *objects.Ingredient {
	return &objects.Ingredient{Title: this.Title}
}

type IngredientMother struct{}

func (IngredientMother) Obj0() *objects.Ingredient {
	b := newIngredientBuilder()
	b.Title = "молоко"
	return b.Build()
}
func (IngredientMother) Obj1() *objects.Ingredient {
	b := newIngredientBuilder()
	b.Title = "говядина"
	return b.Build()
}
func (IngredientMother) Obj2() *objects.Ingredient {
	b := newIngredientBuilder()
	b.Title = "топор"
	return b.Build()
}
func (IngredientMother) Obj3() *objects.Ingredient {
	b := newIngredientBuilder()
	b.Title = "тесак"
	return b.Build()
}
func (this IngredientMother) All() []objects.Ingredient {
	objArr := []objects.Ingredient{
		*this.Obj0(),
		*this.Obj1(),
		*this.Obj2(),
		*this.Obj3(),
	}
	return objArr
}
