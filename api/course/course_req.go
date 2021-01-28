package course

import "buaashow/entity"

// for create course
type courseData struct {
	Name string `json:"name" binding:"require,gte=4,lte=32"`
	Info string `json:"info"`
	entity.Term
}

type studentsData struct {
	Accounts []string `json:"accounts" binding:"require,ascii"`
	Names    []string `json:"names" binding:"require"`
}
