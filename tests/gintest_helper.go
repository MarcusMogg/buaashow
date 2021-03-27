package tests

import (
	"buaashow/response"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/smartystreets/goconvey/convey"
)

var tr *gin.Engine

const baseDir = "./testcases"

type TestCase struct {
	OK   bool   `json:"ok"`
	Desc string `json:"desc"` //描述此testcase
	// Request
	Method  string                 `json:"method"` //请求类型
	URL     string                 `json:"url"`    //链接
	Headers map[string]string      `json:"header"` //headers
	Querys  map[string]string      `json:"query"`  // querys
	Body    map[string]interface{} `json:"body"`   //
	Resp    response.Response      `json:"resp"`   // 应该返回的类型
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
	b, _ := json.Marshal(t.Body)
	r = httptest.NewRequest(t.Method, t.URL+queryStr(t.Querys), bytes.NewBuffer(b))
	c.Request = r
	for k, v := range t.Headers {
		c.Request.Header.Set(k, v)
	}

	tr.ServeHTTP(w, r)
	return
}

type checkFunc func(a, b interface{}) bool

func (v *TestCase) Test(t *testing.T, check checkFunc) *response.Response {
	_, _, w := v.Request()
	var s response.Response
	err := json.Unmarshal(w.Body.Bytes(), &s)

	convey.Convey(v.Desc, t, func() {
		convey.So(err, convey.ShouldBeNil)
		convey.So(s.Code == v.Resp.Code, convey.ShouldBeTrue)
		convey.So(check(v.Resp.Data, s.Data), convey.ShouldBeTrue)
	})
	return &s
}

func NewTestCases(path string) []*TestCase {
	var res []*TestCase
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return res
	}
	err = json.Unmarshal(f, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func Get(key string, m interface{}) interface{} {
	return m.(map[string]interface{})[key]
}

func noCheck(a, b interface{}) bool {
	return true
}

func baseCheck(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func mapCheck(a, b interface{}) bool {
	am := a.(map[string]interface{})
	bm := b.(map[string]interface{})
	for k, v := range am {
		bv, ok := bm[k]
		if !ok {
			return false
		}
		if !baseCheck(v, bv) {
			return false
		}
	}
	return true
}

func arrayChech(a, b interface{}) bool {
	al := a.([]map[string]interface{})
	bl := a.([]map[string]interface{})
	if len(al) > len(bl) {
		return false
	}
	for i := range al {
		if !mapCheck(al[i], bl[i]) {
			return false
		}
	}
	return true
}
