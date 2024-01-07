package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/itsprason/upld/handlers"
)

func main() {
	router := gin.Default()

	// Serve static files (CSS, JS, images, etc.)
	router.Static("/static", "./static")
  router.Static("/uploads", "./uploads")


	// Load HTML templates from the templates directory
	router.LoadHTMLGlob("templates/*")

	// Define a route for the root path ("/") to render HTML
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.PUT("/:filename", handlers.UploadFile)

	// Run the server
	router.Run(":8080")
}
