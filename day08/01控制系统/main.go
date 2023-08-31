package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
)

func getRequest(c *gin.Context) {
	fmt.Println("c.Request:::", c.Request)
	fmt.Println(c.Request.Method)     // GET
	fmt.Println(c.Request.URL)        // /test
	fmt.Println(c.Request.RemoteAddr) // 127.0.0.1:58581
	fmt.Println(c.Request.Header)
	fmt.Println(c.Request.Header["User-Agent"])
	fmt.Println(c.GetHeader("User-Agent"))
	fmt.Println(c.ClientIP())

	c.JSON(200, gin.H{
		"c.Request": "123",
	})
}
func main() {
	r := gin.Default()
	r.Static("/apple", "./banana")
	r.LoadHTMLGlob("./temps/*")
	r.GET("request/info", getRequest)
	r.GET("article/:username/p/:id", func(context *gin.Context) {
		username := context.Param("username")
		id := context.Param("id")
		fmt.Println(username, id)
		fmt.Println(reflect.TypeOf(id))

		context.JSON(200, gin.H{})
	})
	r.GET("wn/jobs", func(context *gin.Context) {
		kd := context.DefaultQuery("kd", "python3")
		gj := context.Query("gj")
		context.JSON(200, gin.H{
			"kd": kd,
			"gj": gj,
		})
	})
	r.POST("/reg", func(context *gin.Context) {
		fmt.Println("name:::", context.PostForm("name"))
		fmt.Println("hobby:::", context.PostForm("hobby"))
		fmt.Println("hobby:::", context.PostFormArray("hobby"))
		context.JSON(200, gin.H{})
	})
	r.POST("/user", func(context *gin.Context) {
		user := User{}
		err := context.ShouldBind(&user)
		if err != nil {
			fmt.Println("user", user)
			fmt.Println("err:::", err)
			return
		}
		fmt.Println("user:::", user)
		fmt.Println(context.Request.Header["Content-Type"])

		context.Redirect(301, "/resp/xml")
	})
	r.GET("resp/xml", func(context *gin.Context) {
		context.XML(200, gin.H{"message": "this is a message!"})
	})
	r.GET("/resp/index.html", func(context *gin.Context) {
		cart := []string{"iphone15", "mac", "ipad"}

		context.HTML(200, "index.html", gin.H{"cart": cart})

	})

	r.Run("127.0.0.1:9090")
}

type User struct {
	Name  string `json:"Name" form:"name"`
	Hobby string `json:"Hobby" form:"hobby"`
}
