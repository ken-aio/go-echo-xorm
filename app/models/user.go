package models

import (
	"github.com/go-xorm/xorm"
	"github.com/ken-aio/go-echo-xorm/app/infra/db"
)

// UserCreateRes new user
type UserCreateRes struct {
	ID int64 `json:"id"`
}

// UserListRes list userr struct
type UserListRes struct {
	ID   int64  `json:"id"`
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
func (u *User) Create() (*UserCreateRes, error) {
	_, err := u.Sess.InsertOne(u.Users)
	if err != nil {
		return nil, err
	}
	return &UserCreateRes{ID: u.Id}, nil
}

// List select user list
func (u *User) List() ([]*UserListRes, error) {
	var users []db.Users
	if err := u.Sess.Asc("id").Find(&users); err != nil {
		return nil, err
	}
	res := make([]*UserListRes, len(users))
	for i := 0; i < len(users); i++ {
		user := users[i]
		res[i] = &UserListRes{ID: user.Id, Name: user.Name}
	}
	return res, nil
}
