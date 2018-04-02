package db

import (
	"time"
)

type Users struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Name      string    `json:"name" xorm:"not null VARCHAR(255)"`
	Birthdate time.Time `json:"birthdate" xorm:"DATE"`
	Gender    string    `json:"gender" xorm:"not null VARCHAR(10)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	CreatedBy string    `json:"created_by" xorm:"not null VARCHAR(64)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	UpdatedBy string    `json:"updated_by" xorm:"not null VARCHAR(64)"`
}

func (db *Users) SetUpdatedBy(trace string) {
	db.UpdatedBy = trace
}

func (db *Users) BeforeInsert() {
	db.CreatedAt = time.Now()
	db.UpdatedAt = time.Now()
	db.CreatedBy = db.UpdatedBy
}

func (db *Users) BeforeUpdate() {
	db.UpdatedAt = time.Now()
}
