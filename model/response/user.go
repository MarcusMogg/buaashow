package response

// LoginRes response data when login success
type LoginRes struct {
	ID    uint   `json:"id"`
	Role  int    `json:"role"`
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
