package core

import "github.com/gin-gonic/gin"

func GetAllClass(context *gin.Context) {
	context.HTML(200, "class.html", nil)
}
