package tests

import "testing"

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

}

func teacherCreate(t *testing.T) {

}

func studentCreate(t *testing.T) {

}

func userLogin(t *testing.T) {

}

func TestUser(t *testing.T) {
	// for order
	t.Run("admin_login", adminLogin)
	t.Run("teacher_create", teacherCreate)
	t.Run("student_create", studentCreate)
	t.Run("user_login", userLogin)
}
