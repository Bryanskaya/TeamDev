package objects

type Grades struct {
	Recipe  int    `gorm:"primary_key;foreignkey:Id" json:"recipe"`
	Account string `gorm:"primary_key;foreignkey:Login" json:"account"`
}
