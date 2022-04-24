package dbuilder

import "api/teamdev/objects"

type CategoryBuilder struct {
	Title string
}
func newCategoryBuilder() *CategoryBuilder { 
	return new(CategoryBuilder) 
}
func (this *CategoryBuilder) Build() *objects.Category { 
	return &objects.Category{Title: this.Title} 
}

type CategoryMother struct{}
func (CategoryMother) Obj0() *objects.Category {
	b := newCategoryBuilder()
	b.Title = "завтраки"
	return b.Build()
}
func (CategoryMother) Obj1() *objects.Category {
	b := newCategoryBuilder()
	b.Title = "салаты"
	return b.Build()
}
func (this CategoryMother) All() []objects.Category {
	objArr := []objects.Category{
		*this.Obj0(), 
		*this.Obj1(),
	}
	return objArr
}