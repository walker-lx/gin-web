package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// get
	r.GET("/gin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	// AsciiJSON
	r.GET("someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(200, data)
	})
	r.Run(":8000")
}
