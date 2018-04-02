package middleware

import (
	"github.com/go-xorm/xorm"
	"github.com/ken-aio/go-echo-xorm/app/infra/db"
	"github.com/labstack/echo"
)

// DBKey db key
const (
	DBKey      = "db.engine"
	SessionKey = "db.session"
)

// CustomContext 独自のコンテキスト
type CustomContext struct {
	echo.Context
}

// GetDB DB接続を取得する
func (d *CustomContext) GetDB() (*xorm.Engine, error) {
	iCon := d.Get(DBKey)
	if iCon != nil {
		con := iCon.(*xorm.Engine)
		// TODO for debug
		con.ShowSQL(true)
		return con, nil
	}
	con, err := db.InitDB()
	if err != nil {
		return nil, err
	}
	d.Set(DBKey, con)
	return con, nil
}

// GetSession DBセッションをContextにsetする
func (d *CustomContext) GetSession() *xorm.Session {
	return d.Get(SessionKey).(*xorm.Session)
}

// SetSession DBセッションをContextにsetする
func (d *CustomContext) SetSession(s *xorm.Session) {
	d.Set(SessionKey, s)
}
