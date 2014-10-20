package models

import (
	"regexp"

	"github.com/revel/revel"
)

type User struct {
	Id        int64  `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Email     string `db:"email" json:"email"`
}

var emailRegex = regexp.MustCompile("\\w*@\\w+[.]\\w+")

func (user *User) Validate(v *revel.Validation) {
	v.Check(user.FirstName,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))

	v.Check(user.LastName,
		revel.ValidRequired(),
		revel.ValidMaxSize(25))

	v.Check(user.Email,
		revel.ValidMatch(emailRegex))
}
