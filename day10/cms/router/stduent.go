package router

import (
	"github.com/gin-gonic/gin"
	. "temp/core"
)

func InitRouter(r *gin.Engine) {
	r.GET("/student", GetAllStudent)
	r.GET("/student/add", AddStudent)
	r.GET("/course", GetAllCourse)
	r.GET("/class", GetAllClass)
}
