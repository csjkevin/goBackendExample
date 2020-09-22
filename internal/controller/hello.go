package controller

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, `<h1>It works!</h1><i>Go Backend Example</i>`)
}
