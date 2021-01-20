package response

// UserInfoRes common user info
type UserInfoRes struct {
	ID    uint   `json:"id"`
	Role  int    `json:"role"`
	Email string `json:"email"`
}

// LoginRes response data when login success
type LoginRes struct {
	UserInfoRes
	Token string `json:"token"`
}

// TicketRes sso response
type TicketRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		ID   string `json:"id"`
		Role int    `json:"role"`
	} `json:"data"`
}
