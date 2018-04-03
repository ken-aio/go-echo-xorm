package db

import (
	"github.com/go-xorm/xorm"
	"github.com/spf13/viper"
)

// InitDB initialize DB
func InitDB() (*xorm.Engine, error) {
	dns := c("user") + ":" + c("password") + "@tcp([" + c("host") + "]:" + c("port") + ")/" + c("dbname")
	if len(c("option")) > 0 {
		dns = dns + "?" + c("option")
	}
	db, err := xorm.NewEngine(c("driver"), dns)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(viper.GetInt("db.settings.max_idle_conns"))
	db.SetMaxOpenConns(viper.GetInt("db.settings.max_idle_conns"))
	db.ShowSQL(viper.GetBool("db.settings.show_sql"))
	return db, nil
}

func c(key string) string {
	return viper.GetString("db.base." + key)
}
