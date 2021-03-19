package user

// LoginData 账号密码登录时传入参数
type loginData struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// LoginTicketData 云平台登录时传入参数
type loginTicketData struct {
	Authorization string `form:"authorization" json:"authorization" binding:"required"`
	ServiceURL    string `form:"url" json:"url" binding:"required"`
}

// EmailData for update email
type emailData struct {
	Email string `form:"email" json:"email" binding:"required,email"`
}

// nameData for update name
type nameData struct {
	Name string `json:"name" binding:"required"`
}

// PasswordData for update password
type passwordData struct {
	OldPassword string `form:"old" json:"old" binding:"required"`
	NewPassword string `form:"new" json:"new" binding:"required,gte=4,lte=16,ascii"`
}

// RegisterData only admin regist teachers
type registerData struct {
	Account  string `form:"account" json:"account" binding:"required,gte=4"`
	Password string `form:"password" json:"password" binding:"required,gte=4,lte=16,ascii"`
}
