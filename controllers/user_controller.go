package controllers

import (
	Authors "Rest_Api/AuthorDTO"
	"Rest_Api/services"
	"github.com/gin-gonic/gin"
)

func SaveAuthor(c *gin.Context) {
	var author Authors.Author
	if err := c.ShouldBindJSON(&author); err != nil {
		panic(err)
	}

	err := services.CreateAuthor(&author)

	if err != nil {
		panic(err)
	}

	c.JSON(200, "The author has been created.")
}
