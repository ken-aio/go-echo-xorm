package logics

import (
	"github.com/go-xorm/xorm"
	"github.com/ken-aio/go-echo-xorm/app/middleware"
	"github.com/labstack/echo"
)

// CommonLogic 共通ロジック
type CommonLogic struct {
	Ctx echo.Context
}

// GetSession セッション情報を取得する
func (c *CommonLogic) GetSession() *xorm.Session {
	cc := c.Ctx.(*middleware.CustomContext)
	return cc.GetSession()
}
