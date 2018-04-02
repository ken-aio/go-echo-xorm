package db

import (
	"time"
)

type Groups struct {
	Id        int64     `xorm:"pk autoincr BIGINT(20)"`
	Name      string    `xorm:"not null VARCHAR(255)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	CreatedBy string    `xorm:"not null VARCHAR(64)"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedBy string    `xorm:"not null VARCHAR(64)"`
}
