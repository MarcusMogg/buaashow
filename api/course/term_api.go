package course

import (
	"buaashow/entity"
	"buaashow/response"
	"time"

	"github.com/gin-gonic/gin"
)

var seasons = [...]time.Month{time.February, time.July, time.September}

// GetTerms gdoc
// @Tags term
// @Summary 获取学期信息，从用户创建时间到当前时间段 [2,6]春[7,8]夏,[9-1]秋,需用户登录
// @Produce application/json
// @Success 200 {array} entity.Term
// @Router /terms [get]
func GetTerms(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	response.OkWithData(terms(u), c)
}

func terms(user *entity.MUser) []entity.Term {
	beginY, beginM := tsToYM(user.CreatedAt)
	curY, curM := tsToYM(time.Now())
	res := make([]entity.Term, 0)
	if beginM < time.February {
		res = append(res, entity.Term{Year: beginY - 1, Season: 2})
	}
	for i, j := range seasons {
		if beginM >= j {
			res = append(res, entity.Term{Year: beginY, Season: i})
		}
	}
	for i := beginY + 1; i < curY; i++ {
		res = append(res, entity.Term{Year: i, Season: 0},
			entity.Term{Year: i, Season: 1},
			entity.Term{Year: i, Season: 2})
	}
	for i, j := range seasons {
		if curM >= j {
			res = append(res, entity.Term{Year: curY, Season: i})
		}
	}
	return res
}

func tsToYM(now time.Time) (int, time.Month) {
	ny := now.Year()
	nm := now.Month()
	return ny, nm
}
