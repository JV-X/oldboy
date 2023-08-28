package router

import (
	"github.com/gin-gonic/gin"
	. "quickstart2/service"
)

func RouterInit(router *gin.Engine) {
	router.GET("index", Index)
	router.GET("login", Login)
	router.GET("books", GetBook)
	router.POST("auth", Auth)
	router.GET("timer", Timer)
	router.GET("reg", Reg)
	router.POST("reg", RegUser)
}
