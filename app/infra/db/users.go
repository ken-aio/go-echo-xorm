package db

import (
	"time"
)

type ICommonColumn interface {
	SetUpdatedBy(trace string)
}

type Users struct {
	Id        uint64    `xorm:"pk autoincr BIGINT(20)"`
	Name      string    `xorm:"not null VARCHAR(255)"`
	Birthdate time.Time `xorm:"DATE"`
	Gender    string    `xorm:"not null VARCHAR(10)"`
	CreatedAt time.Time `xorm:"not null DATETIME"`
	CreatedBy string    `xorm:"not null VARCHAR(64)"`
	UpdatedAt time.Time `xorm:"not null DATETIME"`
	UpdatedBy string    `xorm:"not null VARCHAR(64)"`
}

func (db *Users) SetUpdatedBy(trace string) {
	db.UpdatedBy = trace
}

func (db *Users) BeforeInsert() {
	db.CreatedAt = time.Now()
	db.UpdatedAt = time.Now()
	db.CreatedBy = db.UpdatedBy
}
