package course

// RegisterData only admin regist teachers
type RegisterData struct {
	Account  string `form:"account" json:"account" binding:"required,gte=4"`
	Password string `form:"password" json:"password" binding:"required,gte=4,lte=16,ascii"`
}
