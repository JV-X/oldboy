package core

import "github.com/gin-gonic/gin"

func GetAllCourse(context *gin.Context) {

	context.HTML(200, "course.html", nil)

}
