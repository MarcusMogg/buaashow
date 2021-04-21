package experiment

import (
	"buaashow/entity"
)

// SubmissionReq 提交作业时的请求
type SubmissionReq struct {
	// require when teacher submit
	Account   string            `json:"account"`
	Name      string            `json:"name"`
	Info      string            `json:"info"`
	Type      entity.SummitType `json:"type"`
	SrcURL    string            `json:"src_url"`
	DistURL   string            `json:"dist_url"`
	Readme    string            `json:"readme"`
	Thumbnail string            `json:"thumb"`
}
