package models

import (
	"github.com/go-xorm/xorm"
	"github.com/ken-aio/go-echo-xorm/app/infra/db"
)

// UserCreateRes new user
type UserCreateRes struct {
	ID uint64 `json:"id"`
}

// UserListRes list userr struct
type UserListRes struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// User basic user struct
type User struct {
	*CommonModel
	*db.Users
}

// NewUserModel create user instance
func NewUserModel(s *xorm.Session) *User {
	user := &User{}
	user.CommonModel = &CommonModel{Sess: s}
	user.Users = &db.Users{}
	return user
}

// Create insert into user
func (u *User) Create() (uint64, error) {
	_, err := u.Sess.Insert(u.Users)
	if err != nil {
		return 0, err
	}
	return u.Id, nil
}

// List select user list
func (u *User) List() ([]*UserListRes, error) {
	// select
	//if err != nil {
	//	return nil, err
	//}

	//list := make([]*UserList, len(users))
	//for i := 0; i < len(users); i++ {
	//	user := users[i]
	//	list[i] = &UserList{user.ID, user.Name}
	//}
	return nil, nil
}
