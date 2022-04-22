package repository

import (
	"api/teamdev/objects"

	"github.com/jinzhu/gorm"
)

type StepsRep interface {
	Create(obj *objects.Step) error
	CreateList(obj []objects.Step) error

	List(recipeId int) ([]objects.Step, error)
	FindSteps(id_rcp int) ([]objects.Step, error)
	FindStepByNum(id_rcp, step int) (objects.Step, error)
	FindStepLast(id_rcp int) objects.Step

	Delete(id_rcp, step int) error
	UpdateStep(id_rcp int, step int, obj *objects.Step) error
}

type PGStepsRep struct {
	db *gorm.DB
}

func NewStepsRep(db *gorm.DB) *PGStepsRep {
	return &PGStepsRep{db}
}

func (rep *PGStepsRep) List(recipeId int) (arr []objects.Step, err error) {
	err = rep.db.Where("recipe = ?", recipeId).Find(&arr).Error
	return
}

func (rep *PGStepsRep) FindSteps(id_rcp int) ([]objects.Step, error) {
	temp := []objects.Step{}
	err := rep.db.Where("recipe = ?", id_rcp).Order("num asc").Find(&temp).Error

	return temp, err
}

func (rep *PGStepsRep) FindStepByNum(id_rcp, step int) (objects.Step, error) {
	temp := objects.Step{}
	err := rep.db.Where("recipe = ? AND num = ?", id_rcp, step).Find(&temp).Error

	return temp, err
}

func (rep *PGStepsRep) FindStepLast(id_rcp int) objects.Step {
	temp := objects.Step{}
	rep.db.Where("recipe = ?", id_rcp).Order("num desc").First(&temp)

	return temp
}

func (rep *PGStepsRep) Create(obj *objects.Step) error {
	return rep.db.Create(obj).Error
}

func (rep *PGStepsRep) CreateList(objArr []objects.Step) error {
	for _, obj := range objArr {
		if err := rep.Create(&obj); err != nil {
			return err
		}
	}
	return nil
}

func (rep *PGStepsRep) Delete(id_rcp, step int) error {
	temp := objects.Step{}
	return rep.db.Where("recipe = ? AND num = ?", id_rcp, step).Delete(&temp).Error
}

func (rep *PGStepsRep) UpdateStep(id_rcp int, step int, obj *objects.Step) error {
	return rep.db.Model(&objects.Step{}).Where("recipe = ? AND num = ?", id_rcp, step).Update(objects.Step{Description: obj.Description, Title: obj.Title}).Error
}
