package models

import "github.com/go-xorm/xorm"

// CommonModel モデルに共通のstruct
type CommonModel struct {
	Sess *xorm.Session
}
