package course

// for create course
type courseData struct {
	Name string `json:"name" binding:"required,gte=4,lte=32"`
	Info string `json:"info"`
	TID  uint   `json:"tid"`
}

// TODO：add name
type studentsData struct {
	Accounts []string `json:"accounts" binding:"required,ascii"`
	// Names    []string `json:"names" binding:"required"`
}
