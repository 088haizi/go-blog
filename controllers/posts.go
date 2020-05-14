package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/088haizi/go-blog/models"
	"github.com/088haizi/go-blog/database"
)

func Index(c *gin.Context)  {
	var posts []database.Post
	err := models.All(&posts)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

func Create(c *gin.Context)  {
	var post database.Post
	c.BindJSON(&post)
	err := models.Create(&post)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, post)
	}
}
