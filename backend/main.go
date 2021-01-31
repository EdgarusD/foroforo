package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Thread struct {
	Author string `json:"author" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
	key    string
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	db := []Thread{}
	r := gin.Default()
	r.POST("/thread", func(c *gin.Context) {
		var newThread Thread
		err := c.Bind(&newThread)
		fmt.Println(err)
		if err == nil {
			fmt.Println("se agrego un hilo")
			num := r1.Intn(2147483647)
			num2 := r1.Intn(2147483647)
			s := strconv.Itoa(num)
			s2 := strconv.Itoa(num2)
			key := s + s2
			newThread.key = key
			db = append(db, newThread)
			c.JSON(200, gin.H{"id": len(db) - 1, "key": key})
		}
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"threads": db})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
