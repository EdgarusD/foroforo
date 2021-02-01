package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var db []Thread

type Thread struct {
	Author  string `json:"author" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Body    string `json:"body" binding:"required"`
	key     string
	deleted bool
}

type Key struct {
	Key string `json:"key" binding:"required"`
}

func postThread(c *gin.Context) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	var newThread Thread

	if c.Bind(&newThread) == nil {
		fmt.Println("se agrego un hilo")

		num := r1.Intn(2147483647)
		num2 := r1.Intn(2147483647)
		s := strconv.Itoa(num)
		s2 := strconv.Itoa(num2)
		key := s + s2
		newThread.key = key

		db = append(db, newThread)

		c.JSON(200, gin.H{
			"id":  len(db) - 1,
			"key": key,
		})
	}
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{"threads": db})
}

func deleteThread(c *gin.Context) {
	var key Key
	if c.Bind(&key) == nil {
		for i := 0; i < len(db); i++ {
			if key.Key == db[i].key {
				db[i].deleted = true
			}
		}
	}
}

func main() {
	r := gin.Default()
	r.POST("/thread", postThread)
	r.DELETE("/thread/:id", deleteThread)
	r.GET("/", getHome)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
