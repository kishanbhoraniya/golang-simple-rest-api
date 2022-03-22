package models

// Post model
type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

// create, Update Post model
type CreatePostInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Author      string `json:"author" binding:"required"`
}
