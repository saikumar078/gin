package main

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {

		data := gin.H{
			"message": "hello world",
		}
		ctx.JSON(
			200,
			data,
		)

	})

	router.Run(":8080")
}
