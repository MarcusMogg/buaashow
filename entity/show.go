package entity

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 简略信息
type SummaryResp struct {
	CourseName string `json:"cname"`
	ShowID     `json:"-"`
	UserName   string `json:"uname"`
	Name       string `json:"title"`
	Info       string `json:"info"`
	Type       int    `json:"type"`
	URL        string `json:"url"`
}

// 检索参数 query
type SearchParam struct {
	CourseName string `form:"cname"`
	Recommend  string `form:"rec"`
	TermID     int    `form:"tid"`
	Title      string `form:"title"`
	PageNum    int    `form:"page"`
	PageSize   int    `form:"size"`
}

type ShowID struct {
	EID uint
	GID string
}

func (id *ShowID) Encode() string {
	s := fmt.Sprintf("%d@%s", id.EID, id.GID)
	return base64.URLEncoding.EncodeToString([]byte(s))
}

func DecodeShowID(s string) (*ShowID, error) {
	b, err := base64.URLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	sb := strings.Split(string(b), "@")
	if len(sb) != 2 {
		return nil, errors.New("存在错误字符")
	}

	eid, err := strconv.ParseUint(sb[0], 10, 0)
	if err != nil {
		return nil, errors.New("eid错误")
	}
	res := &ShowID{}
	res.EID = uint(eid)
	res.GID = sb[1]
	return res, nil
}
