package controllers

import (
	"example/rest-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Array to store the posts
var posts = []models.Post{
	models.Post{
		Id:          1,
		Title:       "Post 1",
		Description: "My First Post",
		Author:      "Kishan",
	},
	models.Post{
		Id:          2,
		Title:       "Post 2",
		Description: "My Second Post",
		Author:      "Amit",
	},
}

// Retrive all posts
func GetAllPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": posts})
}

// Get post by id
func GetPost(c *gin.Context) {
	// convert string id to int
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Issue while parsing Id"})
		return
	}
	for _, data := range posts {
		if i == data.Id {
			c.JSON(http.StatusOK, gin.H{"data": data})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "No data found"})
}

// create a new poost
func CreatePost(c *gin.Context) {
	var input models.CreatePostInput
	// bind the user input with CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	post := models.Post{Id: len(posts) + 1, Title: input.Title, Description: input.Description, Author: input.Author}
	posts = append(posts, post)
	c.JSON(http.StatusCreated, gin.H{"data": post})
}

// update a new poost
func UpdatePost(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Issue while parsing Id"})
		return
	}
	var input models.CreatePostInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index, data := range posts {
		if i == data.Id {
			posts = append(posts[:index], posts[index+1:]...)
			post := models.Post{Id: i, Title: input.Title, Description: input.Description, Author: input.Author}
			posts = append(posts, post)
			c.JSON(http.StatusOK, gin.H{"data": post})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "No post found for given id"})
}

// delete a poost
func DeletePost(c *gin.Context) {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Issue while parsing Id"})
		return
	}
	for index, data := range posts {
		if i == data.Id {
			posts = append(posts[:index], posts[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"data": data})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": "No post found for given id"})
}
