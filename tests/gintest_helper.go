package tests

import (
	"buaashow/response"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

var tr *gin.Engine

type TestCase struct {
	Desc string //描述此testcase
	// Request
	Method  string            //请求类型
	URL     string            //链接
	Headers map[string]string //headers
	Querys  map[string]string // querys
	Body    string            //
	// Response
	ShowBody bool              //是否展示返回
	Resp     response.Response // 应该返回的类型
}

func queryStr(m map[string]string) string {
	// len(nil) == 0
	if len(m) == 0 {
		return ""
	}
	res := "?"
	sp := ""
	for k, v := range m {
		res += fmt.Sprintf("%s%s=%s", sp, k, v)
		if len(sp) == 0 {
			sp = "&"
		}
	}
	return res
}

func (t *TestCase) Request() (c *gin.Context, r *http.Request, w *httptest.ResponseRecorder) {
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	r = httptest.NewRequest(t.Method, t.URL+queryStr(t.Querys), bytes.NewBufferString(t.Body))
	c.Request = r
	for k, v := range t.Headers {
		c.Request.Header.Set(k, v)
	}

	tr.ServeHTTP(w, r)
	return
}

func (v *TestCase) Test(t *testing.T) *response.Response {
	_, _, w := v.Request()
	var s response.Response
	err := json.Unmarshal(w.Body.Bytes(), &s)

	convey.Convey(v.Desc, t, func() {
		if v.ShowBody {
			fmt.Printf("接口返回%s\n", w.Body.String())
		}
		convey.So(err, convey.ShouldBeNil)
		convey.So(s.Code == v.Resp.Code, convey.ShouldBeTrue)
		convey.So(reflect.DeepEqual(s.Data, v.Resp.Data), convey.ShouldBeTrue)
	})
	return &s
}
