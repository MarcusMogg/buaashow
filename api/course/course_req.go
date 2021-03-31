package course

// for create course
type courseData struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
	TID  uint   `json:"tid" binding:"required"`
}

type studentsData struct {
	Accounts []string `json:"accounts" binding:"required,ascii"`
	// Names    []string `json:"names" binding:"required"`
}
