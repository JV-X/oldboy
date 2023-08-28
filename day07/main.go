package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "quickstart2/router"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	RouterInit(r)
	err := r.Run("127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
	}
}
