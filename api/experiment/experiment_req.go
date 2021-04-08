package experiment

import (
	"buaashow/entity"
)

// SubmissionReq 提交作业时的请求
type SubmissionReq struct {
	Name      string            `json:"name"`
	Info      string            `json:"info"`
	Type      entity.SummitType `json:"type"`
	URL       string            `json:"url"`
	Readme    string            `json:"readme"`
	Thumbnail string            `json:"thumb"`
}
