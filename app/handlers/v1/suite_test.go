package v1

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ken-aio/go-echo-xorm/app/middleware"
	"github.com/labstack/echo"
)

func TestMain(m *testing.M) {
	// before testing
	code := m.Run()
	// after testing
	os.Exit(code)
}

// buildContext テスト用のcontextを生成する
func buildContext(method string, url string, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, url, strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()

	c := e.NewContext(req, res)
	return c, res
}

// prepareHandler middlewareを準備したhandlerを用意する
func prepareHandler(next echo.HandlerFunc) echo.HandlerFunc {
	handler := middleware.Database(next)
	return handler
}

// toJSON structをjsonに変換する
func toJSON(strct interface{}) string {
	bytes, _ := json.Marshal(strct)
	return string(bytes)
}

// bindRes responseのjsonをstructにbindする
func bindRes(jsonStr string, strct interface{}) {
	if err := json.Unmarshal(([]byte)(jsonStr), &strct); err != nil {
		panic(err)
	}
}
