package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	"main/db"
	"main/helper"
	"main/router"

	"github.com/gin-contrib/cors"
)

func main() {
	ctx := context.Background()

	// Set up a Gin release mode gin.DebugMode, gin.ReleaseMode
	gin.SetMode(gin.ReleaseMode)

	// Set up a Gin router
	r := gin.Default()
	// Maximum allowed file size is 32 MB
	r.MaxMultipartMemory = 1 << 32

	// Set up allow cors
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Set up logging a file.
	setupLogging(r)
	// Set up .env
	setupEnv()
	// Set up mongo
	db.SetupMongoDBClient(ctx)
	// Set up redis
	setupRedisClient(ctx)
	// Set up upload dir
	setupUploadDir()
	// Set up routers
	router.SetupRouter(r)

	// Start the server
	r.Run(":8080")
	fmt.Println("Server is running")
}

func setupEnv() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
}

func setupLogging(r *gin.Engine) {
	// Create a new logrus logger instance
	logger := logrus.New()

	// Set the formatter to the TextFormatter with a timestamp
	logger.SetFormatter(&logrus.TextFormatter{})

	// // Set the output to a file
	file, err := os.OpenFile("gin.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		logger.Info("Failed to log to file, using default stderr")
	}

	logger.SetOutput(file)

	r.Use(gin.LoggerWithWriter(logger.Writer()))
}

func setupRedisClient(ctx context.Context) {
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Check for errors
	_, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}

	// Set data
	// err := client.Set(ctx, "1234", "this is ok", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// // Get data
	// val, err := client.Get(ctx, "1234").Result()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("redis key =", val)
}

func setupUploadDir() {
	dirPath := os.Getenv(helper.EnvDir) + "/upload"
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
}
