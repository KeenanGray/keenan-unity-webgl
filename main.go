package main

import (
	"flag"
	//"fmt"
	"net/http"
	//"time"
	//	"log"
	//"github.com/gorilla/mux"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	//"os"
	//"fmt"
	//"path/filepath"
)

//Go application entrypoint
func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	router.Use(static.Serve("/public", static.LocalFile("./web/public", true)))

	// Setup route group for the API
	api := router.Group("api/")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Server is running",
			})
		})
	}
	router.Run()
}
