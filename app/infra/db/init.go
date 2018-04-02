package db

import (
	"github.com/go-xorm/xorm"
)

// InitDB initialize DB
func InitDB() (*xorm.Engine, error) {
	db, err := xorm.NewEngine("mysql", "sampledb:sampledb@tcp([localhost]:3306)/sampledb?parseTime=true&loc=Asia%2FTokyo") // TODO 設定ファイルにする
	if err != nil {
		return nil, err
	}
	// TODO: Connection pool setting
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	return db, nil
}
