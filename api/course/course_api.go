package course

import (
	"buaashow/entity"
	"buaashow/response"
	"buaashow/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateCourse gdoc
// @Tags course
// @Summary 创建课程 需教师登录
// @accept application/json
// @Produce application/json
// @Param coursedata body courseData true "课程信息"
// @Success 200 {object} entity.CourseResp
// @Router /course [post]
func CreateCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	var req courseData
	if err := c.ShouldBindJSON(&req); err == nil {
		course := entity.MCourse{
			CID:     req.CID,
			Info:    req.Info,
			Teacher: u.Account,
			TID:     req.TID,
		}
		if resp, err := service.CreateCourse(&course, u); err == nil {
			response.OkWithData(resp, c)
		} else {
			response.FailWithMessage(err.Error(), c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// GetMyCourses gdoc
// @Tags course
// @Summary 获取与当前用户相关的课程(教师创建、学生加入) 需用户登录
// @accept application/json
// @Produce application/json
// @Success 200 {array} entity.CourseResp
// @Router /course [get]
func GetMyCourses(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	courses := service.GetMyCourses(u)
	response.OkWithData(courses, c)
}

// GetCourseInfo gdoc
// @Tags course
// @Summary 获取课程信息
// @Produce application/json
// @Param id path int true "Course ID"
// @Success 200 {object} entity.CourseResp
// @Router /course/{id} [get]
func GetCourseInfo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	res, err := service.GetCourseInfoByID(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(res, c)
	}
}

// CreateStudents gdoc
// @Tags course
// @Summary 导入学生, 需用户登录，当前用户有课程管理权限
// @Produce application/json
// @Param id path int true "Course ID"
// @Param accounts body studentsData true "学生账号"
// @Success 200 {object} studentsData
// @Router /course/{id}/students [post]
func CreateStudents(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	cid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	var req studentsData
	if err := c.ShouldBindJSON(&req); err == nil {
		fails, err := service.CreateStudentsToCourse(req.Accounts, uint(cid), u)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		} else {
			response.OkWithData(fails, c)
		}
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// GetStudents gdoc
// @Tags course
// @Summary 查看课程学生列表
// @Produce application/json
// @Param id path int true "Course ID"
// @Success 200 {array} entity.UserInfoRes
// @Router /course/{id}/students [get]
func GetStudents(c *gin.Context) {
	cid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	response.OkWithData(service.GetStudentsInCourse(uint(cid)), c)
}

// DeleteStudent gdoc
// @Tags course
// @Summary 删除学生,需用户登录，当前用户有课程管理权限
// @Produce application/json
// @Param cid path int true "Course ID"
// @Param uid path string true "Student Account"
// @Success 200
// @Router /course/{cid}/student/{uid} [delete]
func DeleteStudent(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	cid, err := strconv.ParseUint(c.Param("cid"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	uid := c.Param("uid")
	err = service.DeleteStudent(uint(cid), uid, u)
	if err == nil {
		response.Ok(c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// DeleteStudent gdoc
// @Tags course
// @Summary 删除所有学生,需用户登录，当前用户有课程管理权限
// @Produce application/json
// @Param cid path int true "Course ID"
// @Param uid path string true "Student Account"
// @Success 200
// @Router /course/{cid}/students [delete]
func DeleteAllStudent(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	cid, err := strconv.ParseUint(c.Param("cid"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	err = service.DeleteAllStudents(uint(cid), u)
	if err == nil {
		response.Ok(c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// DeleteCourse gdoc
// @Tags course
// @Summary 删除课程,需用户登录，当前用户需要是课程创建者
// @Produce application/json
// @Param cid path int true "Course ID"
// @Success 200
// @Router /course/{cid} [delete]
func DeleteCourse(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	cid, err := strconv.ParseUint(c.Param("cid"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}

	err = service.DeleteCourse(uint(cid), u)
	if err != nil {
		response.Ok(c)
	} else {
		response.FailWithMessage(err.Error(), c)
	}
}

// CreateExp godc
// @Tags exp
// @Summary 创建实验 需教师登录
// @Produce application/json
// @Param cid path int true "Course ID"
// @Param exp body entity.ExperimentReq true "实验信息"
// @Success 200
// @Router /course/{cid}/exp [post]
func CreateExp(c *gin.Context) {
	claim, ok := c.Get("user")
	if !ok {
		response.FailWithMessage("未通过jwt认证", c)
		return
	}
	u := claim.(*entity.MUser)
	cid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		zap.S().Debug(err)
		return
	}
	var req entity.ExperimentReq
	if err := c.ShouldBindJSON(&req); err == nil {
		//var begin, end time.Time
		//if begin, err = time.ParseInLocation(global.TimeTemplateSec, req.BeginTime, time.Local); err != nil {
		//	response.FailValidate(c)
		//	zap.S().Debug(err)
		//	return
		//}
		//if end, err = time.ParseInLocation(global.TimeTemplateSec, req.EndTime, time.Local); err != nil {
		//	response.FailValidate(c)
		//	zap.S().Debug(err)
		//	return
		//}
		exp := entity.MExperiment{
			CID:  uint(cid),
			Name: req.Name,
			Info: req.Info,
			Team: req.Team,
		}
		if err = service.CreateExp(&exp, u); err != nil {
			response.FailWithMessage(err.Error(), c)
			zap.S().Debug(err)
			return
		}
		response.OkWithData(gin.H{"eid": exp.ID}, c)
	} else {
		response.FailValidate(c)
		zap.S().Debug(err)
	}
}

// GetExps godc
// @Tags exp
// @Summary 获取课程相关的实验信息
// @Produce application/json
// @Success 200 {array} entity.ExperimentResponse
// @Router /course/{cid}/exps [get]
func GetExps(c *gin.Context) {
	cid, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		response.FailValidate(c)
		return
	}
	res, err := service.GetExpsByCID(uint(cid))
	if err != nil {
		zap.S().Debug(err.Error())
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(res, c)
	}
}
