package user

import (
	"buaashow/entity"
)

// LoginRes response data when login success
type loginRes struct {
	entity.UserInfoRes
	Token string `json:"token"`
}

// TicketRes sso response
type ticketRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID   string `json:"id"`
		Role string `json:"role"`
	} `json:"data"`
}
