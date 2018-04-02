package logics

import (
	"time"

	"github.com/ken-aio/go-echo-xorm/app/models"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// UserLogic user logic
type UserLogic struct {
	*CommonLogic
}

// NewUserLogic create new user instance
func NewUserLogic(c echo.Context) *UserLogic {
	return &UserLogic{&CommonLogic{Ctx: c}}
}

// UserCreate new user create
func (u *UserLogic) UserCreate(name string, birthdate time.Time, gender string) (*models.UserCreateRes, error) {
	sess := u.GetSession()
	if err := sess.Begin(); err != nil {
		return nil, err
	}
	user := models.NewUserModel(sess)
	user.Name = name
	user.Birthdate = birthdate
	user.Gender = gender
	id, err := user.Create()
	if err != nil {
		if err := sess.Rollback(); err != nil {
			logrus.Error("Error in user_logic.Create rollback.", err)
		}
		return nil, err
	}
	sess.Commit()
	return id, nil
}

// UserList list user
func (u *UserLogic) UserList() ([]*models.UserListRes, error) {
	sess := u.GetSession()
	user := models.NewUserModel(sess)
	return user.List()
}
