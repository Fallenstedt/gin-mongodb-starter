package main

import (
  "log"
  "os"
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  
  mode := os.Getenv("MODE")
  if mode == "release" {
    gin.SetMode(gin.ReleaseMode)
  }

  if err != nil {
    log.Fatal("Error loading .env file")
  }

  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

  r.Run()
}