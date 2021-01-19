package request

// LoginData 账号密码登录时传入参数
type LoginData struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// LoginTicketData 云平台登录时传入参数
type LoginTicketData struct {
	Authorization string `form:"authorization" json:"authorization" binding:"required"`
	ServiceURL    string `form:"url" json:"url" binding:"required"`
}
