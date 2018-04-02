package db

import (
	"time"
)

type GroupMembers struct {
	Id        int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	IsAdmin   int       `json:"is_admin" xorm:"not null TINYINT(1)"`
	UserId    int64     `json:"user_id" xorm:"not null index BIGINT(20)"`
	GroupId   int64     `json:"group_id" xorm:"not null index BIGINT(20)"`
	CreatedAt time.Time `json:"created_at" xorm:"not null DATETIME"`
	CreatedBy string    `json:"created_by" xorm:"not null VARCHAR(64)"`
	UpdatedAt time.Time `json:"updated_at" xorm:"not null DATETIME"`
	UpdatedBy string    `json:"updated_by" xorm:"not null VARCHAR(64)"`
}

func (db *GroupMembers) SetUpdatedBy(trace string) {
	db.UpdatedBy = trace
}

func (db *GroupMembers) BeforeInsert() {
	db.CreatedAt = time.Now()
	db.UpdatedAt = time.Now()
	db.CreatedBy = db.UpdatedBy
}

func (db *GroupMembers) BeforeUpdate() {
	db.UpdatedAt = time.Now()
}
