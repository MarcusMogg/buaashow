package user

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

// EmailData for update email
type EmailData struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

// PasswordData for update password
type PasswordData struct {
	OldPassword string `form:"old" json:"old" binding:"required"`
	NewPassword string `form:"new" json:"new" binding:"required,gte=4,lse=16,asciiprint"`
}
