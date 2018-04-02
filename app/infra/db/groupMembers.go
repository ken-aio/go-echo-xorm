package db

import (
	"time"
)

type GroupMembers struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	IsAdmin   int       `xorm:"not null TINYINT(1)"`
	UserId    int64     `xorm:"not null index BIGINT(20)"`
	GroupId   int64     `xorm:"not null index BIGINT(20)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	CreatedBy string    `xorm:"not null VARCHAR(64)"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedBy string    `xorm:"not null VARCHAR(64)"`
}
