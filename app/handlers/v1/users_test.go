package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ken-aio/go-echo-xorm/app/logics"
	"github.com/ken-aio/go-echo-xorm/app/models"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type UserLogicMock struct{}

func (u *UserLogicMock) NewUserLogic(c echo.Context) logics.IUserLogic {
	return &UserLogicMock{}
}

func (u *UserLogicMock) UserCreate(name string, birthdate time.Time, gender string) (*models.UserCreateRes, error) {
	return &models.UserCreateRes{ID: 1}, nil
}

func (u *UserLogicMock) UserList() ([]*models.UserListRes, error) {
	resp := []*models.UserListRes{
		&models.UserListRes{ID: 1, Name: "test1"},
		&models.UserListRes{ID: 2, Name: "test2"},
	}
	return resp, nil
}

func TestUserCreate(t *testing.T) {
	body := &userCreateReq{Name: "test", Birthdate: time.Now(), Gender: "male"}
	c, res := buildContext(echo.POST, "/api/v1/users", toJSON(body))

	h := NewUserHandler(&UserLogicMock{})
	if assert.NoError(t, h.UserCreate(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		actual := &models.UserCreateRes{}
		json.Unmarshal(([]byte)(res.Body.String()), &actual)
		assert.Equal(t, actual.ID, int64(1))
	}
}

func TestUserList(t *testing.T) {
	c, res := buildContext(echo.GET, "/api/v1/users", "")

	h := NewUserHandler(&UserLogicMock{})
	if assert.NoError(t, h.UserList(c)) {
		assert.Equal(t, http.StatusOK, res.Code)
		actual := make([]models.UserListRes, 0)
		json.Unmarshal(([]byte)(res.Body.String()), &actual)
		assert.Equal(t, 2, len(actual))
	}
}

func TestUserUpdate(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewUserHandler(&UserLogicMock{})
	if assert.NoError(t, h.UserUpdate(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"update ok\"", rec.Body.String())
	}
}

func TestUserDelete(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/api/v1/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := NewUserHandler(&UserLogicMock{})
	if assert.NoError(t, h.UserDelete(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "\"delete ok\"", rec.Body.String())
	}
}
