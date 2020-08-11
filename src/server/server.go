package server

import (
  "log"
  "os"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func Init() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  mode := os.Getenv("MODE")
  if mode == "release" {
    gin.SetMode(gin.ReleaseMode)
  }
  

  r := NewRouter()
  r.Run()
}