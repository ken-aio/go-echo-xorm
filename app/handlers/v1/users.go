package v1

import (
	"net/http"
	"time"

	"github.com/ken-aio/go-echo-xorm/app/logics"
	"github.com/labstack/echo"
)

type (
	userCreateReq struct {
		Name      string    `json:"name"`
		Birthdate time.Time `json:"birthdate"`
		Gender    string    `json:"gender"`
	}
	userHandler struct {
		logic logics.IUserLogic
	}
)

// NewUserHandler new user handler
func NewUserHandler(l logics.IUserLogic) *userHandler {
	return &userHandler{logic: l}
}

// UserCreate User create API
// @Summary User create API
// @Description create new user
// @Accept  json
// @Produce  json
// @Param   body     body    v1.userCreateReq     true        "user create parameter"
// @Success 200 {object} models.UserCreateRes	""
// @Router /api/v1/users [post]
func (h *userHandler) UserCreate(c echo.Context) error {
	u := &userCreateReq{}
	if err := c.Bind(u); err != nil {
		return err
	}
	logic := h.logic.NewUserLogic(c)
	resp, err := logic.UserCreate(u.Name, u.Birthdate, u.Gender)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

// UserList User list API
// @Summary User list API
// @Description list users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserListRes	""
// @Router /api/v1/users [get]
func (h *userHandler) UserList(c echo.Context) error {
	logic := h.logic.NewUserLogic(c)
	resp, err := logic.UserList()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, resp)
}

// UserUpdate User update API
// @Summary User update API
// @Description update users
// @Accept  json
// @Produce  json
// @Param   user_id     path    int     true        "user id parameter"
// @Success 200 string string	""
// @Router /api/v1/users/{user_id} [put]
func (h *userHandler) UserUpdate(c echo.Context) error {
	return c.JSON(http.StatusOK, "update ok")
}

// UserDelete User delete API
// @Summary User delete API
// @Description delete users
// @Accept  json
// @Produce  json
// @Param   user_id     path    int     true        "user id parameter"
// @Success 200 string string	""
// @Router /api/v1/users/{user_id} [delete]
func (h *userHandler) UserDelete(c echo.Context) error {
	return c.JSON(http.StatusOK, "delete ok")
}
