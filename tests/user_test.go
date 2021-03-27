package tests

import (
	"path"
	"testing"
)

type userdata struct {
	account  string
	password string
	token    string
}

var adminToken string
var teachers = []userdata{
	{"teacher_1", "sfdsdf,a4s4", ""},
	{"teacher_2", "asd.a.d,12#", ""},
}
var students = []userdata{
	{"student_1", "as45sa4c55#", ""},
	{"student_2", "as45sa4c55#", ""},
	{"student_3", "as45sa4c55#", ""},
	{"student_4", "as45sa4c55#", ""},
	{"student_5", "as45sa4c55#", ""},
}

func adminLogin(t *testing.T) {
	cases := NewTestCases(path.Join(baseDir, "user", "adminLogin.json"))
	for _, i := range cases {
		resp := i.Test(t, mapCheck)
		if i.OK {
			adminToken = Get("token", resp.Data).(string)
		}
	}
}

func teacherCreate(t *testing.T) {
	cases := NewTestCases(path.Join(baseDir, "user", "teacherCreate.json"))
	for _, i := range cases {
		i.Headers["Authorization"] = adminToken
		i.Test(t, noCheck)
	}
}

func teacherLogin(t *testing.T) {
	cases := NewTestCases(path.Join(baseDir, "user", "teacherLogin.json"))
	cur := 0
	for _, i := range cases {
		resp := i.Test(t, mapCheck)
		if i.OK {
			if cur < 2 {
				teachers[cur].token = Get("token", resp.Data).(string)
			}
			cur++
		}
	}
}
func studentCreate(t *testing.T) {
	cases := NewTestCases(path.Join(baseDir, "user", "studentCreate.json"))
	cur := 0
	for _, i := range cases {
		i.Headers["Authorization"] = adminToken
		resp := i.Test(t, noCheck)
		if i.OK {
			teachers[cur].token = Get("token", resp.Data).(string)
			cur++
		}
	}
}

func userLogin(t *testing.T) {
	cases := NewTestCases(path.Join(baseDir, "user", "userLogin.json"))
	cur := 0
	for _, i := range cases {
		resp := i.Test(t, mapCheck)
		if i.OK {
			if cur < 2 {
				teachers[cur].token = Get("token", resp.Data).(string)
			}
			cur++
		}
	}
}

func TestUser(t *testing.T) {
	// for order
	t.Run("admin_login", adminLogin)
	//t.Run("teacher_create", teacherCreate)
	//t.Run("teacher_login", teacherLogin)
	//t.Run("student_create", studentCreate)
	//t.Run("user_login", userLogin)
}
