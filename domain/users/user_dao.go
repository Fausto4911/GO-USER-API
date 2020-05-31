package users

import (
	"fmt"
	"github.com/fausto4911/GO-USER-API/utils/date_utils"
	"github.com/fausto4911/GO-USER-API/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	result := usersDB[u.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.Id))
	}
	u.Id = result.Id
	u.Email = result.Email
	u.Password = result.Password
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.UpdateAt = result.UpdateAt
	u.CreateAt = result.CreateAt
	return nil
}

func (u *User) Save() *errors.RestErr {
	current := usersDB[u.Id]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exist", u.Id))
	}
	u.CreateAt = date_utils.GetNowString()
	usersDB[u.Id] = u
	return nil
}
