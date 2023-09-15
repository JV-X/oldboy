package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	. "temp/router"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("student.html", "temps/base.html", "temps/student.html")
	r.AddFromFiles("course.html", "temps/base.html", "temps/course.html")
	r.AddFromFiles("class.html", "temps/base.html", "temps/class.html")
	return r
}

func main() {
	r := gin.Default()
	// 静态文件配置
	r.Static("/static", "./static")
	r.LoadHTMLGlob("./temps/*")
	r.HTMLRender = createMyRender()

	InitRouter(r)

	r.Run(":8082") // 监听并在 0.0.0.0:8080 上启动服务
}
