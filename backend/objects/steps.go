package objects

import "encoding/json"

type Step struct {
	Recipe      int    `json:"recipe" gorm:"UNIQUE_INDEX:stepindex;foreignkey:Id"`
	Num         int    `json:"num" gorm:"UNIQUE_INDEX:stepindex;autoincrement:true"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

func (Step) TableName() string {
	return "steps"
}

type StepDTO struct {
	Recipe      int    `json:"recipe"`
	Num         int    `json:"num"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

func (obj *Step) ToDTO() *StepDTO {
	dto := new(StepDTO)
	jsonStr, _ := json.Marshal(obj)
	json.Unmarshal(jsonStr, dto)
	return dto
}

func (Step) ArrToDTO(src []Step) []StepDTO {
	dst := make([]StepDTO, len(src))
	for k, v := range src {
		dst[k] = *v.ToDTO()
	}
	return dst
}

func (obj *StepDTO) ToModel() *Step {
	mod := new(Step)
	jsonStr, _ := json.Marshal(obj)
	json.Unmarshal(jsonStr, mod)
	return mod
}
