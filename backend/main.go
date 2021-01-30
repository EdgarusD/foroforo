package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Thread struct {
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
}

func main() {
	db := []Thread{}
	r := gin.Default()
	r.POST("/thread", func(c *gin.Context) {
		var thread Thread
		if c.Bind(&thread) == nil {
			db = append(db, thread)
			fmt.Println("se agrego un hilo")
			c.JSON(200, gin.H{"id": len(db) - 1})
		}
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"threads": db})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
