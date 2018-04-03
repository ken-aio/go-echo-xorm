package routes

import (
	"net/http"
	"time"

	"github.com/ken-aio/go-echo-xorm/app/handlers/v1"
	"github.com/ken-aio/go-echo-xorm/app/logics"
	"github.com/ken-aio/go-echo-xorm/app/middleware"
	"github.com/ken-aio/go-echo-xorm/config"
	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/swaggo/echo-swagger"
)

// Init initialize echo application
func Init() *echo.Echo {
	e := echo.New()

	config.LoadConfig(e)
	e.Debug = viper.GetBool("debug")

	middleware.Init(e)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/routes", func(c echo.Context) error {
		return c.JSON(http.StatusOK, e.Routes())
	})
	// ping
	e.GET("/ping", func(c echo.Context) error {
		resp := map[string]time.Time{"pong": time.Now()}
		return c.JSON(http.StatusOK, resp)
	})

	apiV1 := e.Group("api/v1")
	{
		users := apiV1.Group("/users")
		{
			u := v1.NewUserHandler(&logics.UserLogic{})
			users.POST("", u.UserCreate)
			users.GET("", u.UserList)
			users.PUT("/:id", u.UserUpdate)
			users.DELETE("/:id", u.UserDelete)
		}
		groups := apiV1.Group("/groups")
		{
			mems := groups.Group("/:group_id/members")
			{
				mems.GET("", v1.GroupMemberList)
			}
		}
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}
