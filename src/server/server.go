package server

import (
  "log"
  "os"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func Init() {
  mode := os.Getenv("MODE")
  if mode == "release" {
    gin.SetMode(gin.ReleaseMode)
  }
  
  err := godotenv.Load()
  
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  r := NewRouter()
  r.Run()
}