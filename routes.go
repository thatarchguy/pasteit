package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thatarchguy/pasteit/models"
)

// createPost add a new Post
func createPost(c *gin.Context) {
	// Allocate an empty Post
	post := &models.Post{}
	file, fileErr := c.FormFile("f")
	if fileErr != nil {
		log.Println(fileErr)
	}
	src, _ := file.Open()
	defer src.Close()
	var buf bytes.Buffer
	io.Copy(&buf, src)
	post.Post = buf.String()

	// Generate URI
	var postURI string
	var lastPost models.Post
	err := db.Last(&lastPost).Error
	//err := db.First(&lastPost, 4)
	if err != nil {
		// First post ever?!
		postURI = "aaa"
	} else {
		postURI = models.IncrementURI(lastPost.URI)
	}
	post.URI = postURI
	db.Save(&post)

	// Return
	c.String(http.StatusCreated, fmt.Sprintf("%v: http://%s/%s", post.ID, config.hostname, post.URI))
}

// fetchAllPost fetch all Posts
func fetchAllPost(c *gin.Context) {
	var post []models.Post

	db.Find(&post)

	if len(post) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No posts found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": post})
}

// fetchSinglePost fetch a single Post
func fetchSinglePost(c *gin.Context) {
	var post models.Post
	postURI := c.Param("id")

	db.Where("uri = ?", postURI).First(&post)

	if post.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No post found!"})
		return
	}

	c.String(http.StatusOK, post.Post)
}
