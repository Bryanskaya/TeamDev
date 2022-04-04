package objects

import (
	"encoding/json"
)

type Account struct {
	Login          string `json:"login" gorm:"primary_key"`
	Role           string `json:"role"`
	Salt           string `json:"salt"`
	HashedPassword string `json:"hashed_password"`

	Grades []*Recipe			`gorm:"many2many:grades;"`
}

type AccountDTO struct {
	Login			string `json:"login" example:"admin"`
	Role			string `json:"role" example:"admin"`
	Password		string `json:"password" example:"admin"`
}

func (Account) TableName() string {
	return "accounts"
}

func (this *Account) ToDTO() *AccountDTO {
	dto := new(AccountDTO)
	jsonStr, _ := json.Marshal(this)
	json.Unmarshal(jsonStr, dto) 
	return dto
}

func (Account) ArrToDTO(src []Account) []AccountDTO {
	dst := make([]AccountDTO, len(src))
	for k, v := range src {
		dst[k] = *v.ToDTO()
	}
	return dst
}

func (this *AccountDTO) ToModel() *Account {
	mod := new(Account)

	jsonStr, _ := json.Marshal(this)
	json.Unmarshal(jsonStr, mod)
	mod.HashedPassword = this.Password
	return mod
}
