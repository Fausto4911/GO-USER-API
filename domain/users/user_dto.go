package users

import (
	"github.com/fausto4911/GO-USER-API/utils/errors"
	"strings"
	"time"
)

type User struct {
	Id        int64      `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
}


func (u *User) Validate() *errors.RestErr {
	u.Email = strings.TrimSpace(strings.ToLower(u.Email))
	if u.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}
