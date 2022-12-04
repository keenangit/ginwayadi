package main

import (
	"fmt"
	"ginwayadi/config"
	"ginwayadi/controllers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/twinj/uuid"
)

func main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(CORSMiddleware())
	r.Use(RequestIDMiddleware())
	// r.SetTrustedProxies([]string{"192.168.1.2"})

	go config.InitWA()

	v1 := r.Group("/v1")
	{
		qr := new(controllers.QRController)
		v1.GET("/wa/qr", BasicAuthMiddleware(), qr.GetQR)

		wa := new(controllers.MessageController)
		v1.POST("/wa/send", BasicAuthMiddleware(), wa.Send)

	}

	port := os.Getenv("PORT")

	log.Printf("\n\n PORT: %s \n ENV: %s \n SSL: %s \n Version: %s \n\n", port, os.Getenv("ENV"), os.Getenv("SSL"), os.Getenv("API_VERSION"))

	r.Run(":" + port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//CORSMiddleware ...
//CORS (Cross-Origin Resource Sharing)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("X-Frame-Options", "SAMEORIGIN")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

//RequestIDMiddleware ...
//Generate a unique ID and attach it to each request for future reference or use
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

var auth = new(controllers.AuthController)

//BasicAuth ...
func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.Valid(c)
		c.Next()
	}
}
