package service

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Auth(c *gin.Context) {
	user := c.PostForm("user")
	pwd := c.PostForm("pwd")
	if user == "admin" && pwd == "123456" {
		c.String(200, "登陆成功！")
	} else {
		c.String(200, "用户名或密码错误！")
	}
}

func Reg(context *gin.Context) {
	context.HTML(200, "reg.html", nil)
}

func RegUser(context *gin.Context) {
	user := context.PostForm("user")
	pwd := context.PostForm("pwd")
	gender := context.PostForm("gender")
	province := context.PostForm("province")

	context.JSON(200, gin.H{
		"user":     user,
		"pwd":      pwd,
		"province": province,
		"gender":   gender,
	})

}
