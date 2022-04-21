package dbuilder

import (
	"api/teamdev/objects"
	"api/teamdev/utils"
)

type AccountBuilder struct {
	Login          string
	Role           string
	Salt           string
	HashedPassword string
}

func newAccountBuilder() *AccountBuilder {
	return new(AccountBuilder)
}
func (this *AccountBuilder) Build() *objects.Account {
	this.Salt, this.HashedPassword = utils.HashPassword(this.Login, this.Salt)
	return &objects.Account{Login: this.Login, Role: this.Role,
		Salt: this.Salt, HashedPassword: this.HashedPassword}
}

/*
WARNING: don't create new accounts in one test, use old one (hash generated differently every time)
	// BAD:
	objArr := AccountMother{}.All()
	objAcc := AccountMother{}.Obj0()

	// GOOD:
	objArr := AccountMother{}.All()
	objAcc := &objArr[0]
*/
type AccountMother struct{}

func (AccountMother) Obj0() *objects.Account {
	b := newAccountBuilder()
	b.Login = "test1"
	b.Role = "admin"
	return b.Build()
}

func (AccountMother) Obj1() *objects.Account {
	b := newAccountBuilder()
	b.Login = "test2"
	b.Role = "user"
	return b.Build()
}

func (AccountMother) Obj1Udp() *objects.Account {
	b := newAccountBuilder()
	b.Login = "test2"
	b.Role = "admin"
	return b.Build()
}

func (AccountMother) Obj2() *objects.Account {
	b := newAccountBuilder()
	b.Login = "test3"
	b.Role = "user"
	return b.Build()
}

func (this AccountMother) All() []objects.Account {
	objArr := []objects.Account{
		*this.Obj0(),
		*this.Obj1(),
		*this.Obj2(),
	}
	return objArr
}
