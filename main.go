package main

import (
	context "context"
	"flag"
	"fmt"

	"log"
	"math/rand"

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

const (
	port       = ":8081"
	colorBytes = 3
)

type server struct{}

//Return a random color in the form of a color hex string
//e.g. "#00FF00" would be RRGGBB so Green
func (s *server) GetRandomColor(ctx context.Context, curr *pb.CurrentColor) (*pb.NewColor, error) {
	hex := "#" + randomHex()
	log.Printf("Client's current color: [#%v] sending [%v]", curr.Color, hex)
	return &pb.NewColor{Color: hex}, nil
}

//Create a random hex string of N digits
func randomHex() string {
	bytes := make([]byte, colorBytes)
	if _, err := rand.Read(bytes); err != nil {
		log.Panicln("Error generating random hex value", err)
	}
	return fmt.Sprintf("%X", bytes)
}

//Go application entrypoint
func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./web", true)))
	router.Use(static.Serve("/public", static.LocalFile("./public", true)))

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
