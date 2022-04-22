package dbuilder

import "api/teamdev/objects"

type StepBuilder struct {
	Recipe      int
	Num         int
	Description string
	Title       string
}

func newStepBuilder() *StepBuilder {
	return new(StepBuilder)
}
func (builder *StepBuilder) Build() *objects.Step {
	return &objects.Step{Recipe: builder.Recipe, Num: builder.Num,
		Description: builder.Description, Title: builder.Title}
}

type StepMother struct{}

func (StepMother) Obj0() *objects.Step {
	b := newStepBuilder()
	b.Recipe = 1
	b.Num = 1
	b.Description = "Подогреваем молоко"
	b.Title = "Шаг 1"
	return b.Build()
}
func (StepMother) Obj1() *objects.Step {
	b := newStepBuilder()
	b.Recipe = 1
	b.Num = 2
	b.Description = "Добавлем яйца"
	b.Title = "Шаг 2"
	return b.Build()
}
func (StepMother) Obj2() *objects.Step {
	b := newStepBuilder()
	b.Recipe = 3
	b.Num = 1
	b.Description = "Подогреваем кефир"
	b.Title = "Шаг 1"
	return b.Build()
}

func (mother StepMother) All() []objects.Step {
	objArr := []objects.Step{
		*mother.Obj0(), 
		*mother.Obj1(),
		*mother.Obj2(),
	}
	return objArr
}
