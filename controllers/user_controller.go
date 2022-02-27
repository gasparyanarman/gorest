package controllers

import "github.com/gin-gonic/gin"

func GetHomePage(c *gin.Context) {
	c.JSON(200, "Hello World")
}
