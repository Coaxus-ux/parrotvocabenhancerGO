package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	r := gin.Default()

	r.GET("/audio/:name", func(c *gin.Context) {
		name := c.Param("name")
		accent := c.Query("accent")
		fmt.Println("name: ", name)
		fmt.Println("accent: ", accent)
		if accent != "uk" && accent != "us" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Accent must be 'uk' or 'us'"})
			return
		}

		audioPath := filepath.Join("audios", accent, name+".mp3")

		c.File(audioPath)
	})

	r.Run(":8080")
}
