package logics

import (
	"time"

	"github.com/ken-aio/go-echo-xorm/app/middleware"
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
func (u *UserLogic) UserCreate(name string, birthdate time.Time, gender string) (uint64, error) {
	cc := u.Ctx.(*middleware.CustomContext)
	sess := cc.GetSession()
	if err := sess.Begin(); err != nil {
		return 0, err
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
		return 0, err
	}
	sess.Commit()
	return id, nil
}

// UserList list user
func UserList() ([]*models.UserListRes, error) {
	u := &models.User{}
	return u.List()
}
