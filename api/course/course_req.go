package course

// for create course
type courseData struct {
	CID  uint   `json:"name_id" binding:"required"`
	Info string `json:"info"`
	TID  uint   `json:"tid" binding:"required"`
}

type studentsData struct {
	Accounts []string `json:"accounts" binding:"required,ascii"`
	// Names    []string `json:"names" binding:"required"`
}

type courseName struct {
	NID       uint   `json:"id"`
	Name      string `json:"name"`
	Info      string `json:"info"`
	Thumbnail string `json:"thumb"`
}
