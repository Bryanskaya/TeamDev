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
func (builder *AccountBuilder) Build() *objects.Account {
	builder.Salt, builder.HashedPassword = utils.HashPassword(builder.Login, builder.Salt)
	return &objects.Account{Login: builder.Login, Role: builder.Role,
		Salt: builder.Salt, HashedPassword: builder.HashedPassword}
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

func (mother AccountMother) All() []objects.Account {
	objArr := []objects.Account{
		*mother.Obj0(),
		*mother.Obj1(),
		*mother.Obj2(),
	}
	return objArr
}
