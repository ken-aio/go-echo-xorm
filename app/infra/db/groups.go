package db

import (
	"time"
)

type Groups struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	Name      string    `json:"name" xorm:"not null VARCHAR(255)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	CreatedBy string    `json:"created_by" xorm:"not null VARCHAR(64)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	UpdatedBy string    `json:"updated_by" xorm:"not null VARCHAR(64)"`
}

func (db *Groups) SetUpdatedBy(trace string) {
	db.UpdatedBy = trace
}

func (db *Groups) BeforeInsert() {
	db.CreatedAt = time.Now()
	db.UpdatedAt = time.Now()
	db.CreatedBy = db.UpdatedBy
}

func (db *Groups) BeforeUpdate() {
	db.UpdatedAt = time.Now()
}
