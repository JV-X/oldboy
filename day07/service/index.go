package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GetBook(c *gin.Context) {
	c.JSON(200, gin.H{
		"books": []string{"西游记", "三国演义", "金瓶梅", "国色天香"},
	})
}

func Timer(c *gin.Context) {
	now := time.Now().Unix()
	fmt.Println("now", now)
	c.HTML(200, "timer.html", gin.H{
		"now": now,
	})
}
