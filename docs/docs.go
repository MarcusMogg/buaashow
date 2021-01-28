// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mogg",
            "url": "https://github.com/MarcusMogg"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/course": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "获取与当前用户相关的课程(教师创建、学生加入) 需用户登录",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.CourseResp"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "创建课程 需教师登录",
                "parameters": [
                    {
                        "description": "课程信息",
                        "name": "coursedata",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/course.courseData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.CourseResp"
                        }
                    }
                }
            }
        },
        "/course/{cid}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "删除学生,需用户登录，当前用户需要是课程创建者",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
                        "name": "cid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/course/{cid}/student/{uid}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "删除学生,需用户登录，当前用户有课程管理权限",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
                        "name": "cid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/course/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "获取课程信息, 需用户登录，当前用户需要与课程相关",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.CourseResp"
                        }
                    }
                }
            }
        },
        "/course/{id}/students": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "查看课程学生列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.UserInfoRes"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "course"
                ],
                "summary": "导入学生, 需用户登录，当前用户有课程管理权限",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Course ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "学生账号",
                        "name": "accounts",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/course.studentsData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/course.studentsData"
                        }
                    }
                }
            }
        },
        "/terms": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "term"
                ],
                "summary": "获取学期信息，从用户创建时间到当前时间段 [2,6]春[7,8]夏,[9-1]秋,需用户登录",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Term"
                            }
                        }
                    }
                }
            }
        },
        "/user/email": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "修改邮箱, 需用户登录",
                "parameters": [
                    {
                        "description": "新邮箱",
                        "name": "ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.emailData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.loginRes"
                        }
                    }
                }
            }
        },
        "/user/info": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取当前用户信息，需用户登录",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.UserInfoRes"
                        }
                    }
                }
            }
        },
        "/user/info/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "获取指定id的用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.UserInfoRes"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "使用账号密码登录",
                "parameters": [
                    {
                        "description": "账号密码",
                        "name": "logindata",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.loginData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.loginRes"
                        }
                    }
                }
            }
        },
        "/user/password": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "修改密码, 需用户登录",
                "parameters": [
                    {
                        "description": "新旧密码",
                        "name": "ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.passwordData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.loginRes"
                        }
                    }
                }
            }
        },
        "/user/teacher": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "创建教师账号 需管理员登录",
                "parameters": [
                    {
                        "description": "账号密码必选，邮箱可选",
                        "name": "logindata",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.registerData"
                        }
                    }
                ]
            }
        },
        "/user/verify": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "使用云平台登录",
                "parameters": [
                    {
                        "description": "云平台返回的ticket",
                        "name": "ticket",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.loginTicketData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.loginRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "course.courseData": {
            "type": "object",
            "properties": {
                "info": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "season": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "course.studentsData": {
            "type": "object",
            "properties": {
                "accounts": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "entity.CourseResp": {
            "type": "object",
            "properties": {
                "cid": {
                    "type": "integer"
                },
                "info": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "season": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "entity.Term": {
            "type": "object",
            "properties": {
                "season": {
                    "type": "integer"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "entity.UserInfoRes": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "integer"
                }
            }
        },
        "user.emailData": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "user.loginData": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user.loginRes": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "user.loginTicketData": {
            "type": "object",
            "required": [
                "authorization",
                "url"
            ],
            "properties": {
                "authorization": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "user.passwordData": {
            "type": "object",
            "required": [
                "new",
                "old"
            ],
            "properties": {
                "new": {
                    "type": "string"
                },
                "old": {
                    "type": "string"
                }
            }
        },
        "user.registerData": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "buaashow",
	Description: "buaashow is a sample RESTful api server.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
