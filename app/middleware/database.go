package middleware

import (
	"strings"

	"github.com/ken-aio/go-echo-xorm/app/infra/db"
	"github.com/labstack/echo"
)

// Database database middleware
func Database(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &CustomContext{c}
		con, err := cc.GetDB()
		if err != nil {
			return err
		}

		ua := strings.Join(c.Request().Header["User-Agent"], ",")
		if len(ua) > 20 {
			ua = ua[0:20]
		}
		trace := c.Request().Method + " " + c.Path() + " " + ua
		sess := con.NewSession()
		defer sess.Close()
		sess = sess.Before(func(bean interface{}) {
			bean.(db.ICommonColumn).SetUpdatedBy(trace)
		})
		cc.SetSession(sess)
		return next(cc)
	}
}
