package server

import (
  "github.com/gin-gonic/gin"
   "github.com/fallenstedt/gin-example/src/config"
  "log"
)

func Init() {
  c := config.Get()
  if c.Mode == "release" {
    gin.SetMode(gin.ReleaseMode)
  }
  printIntro()
  r := NewRouter()
  r.Run()
}

func printIntro() {
  log.Print("Welcome to a basic Gin example")
  log.Print("   \n  .\n  .\n . .\n  ...\n\\~~~~~/\n \\   /\n  \\ /\n   V\n   |\n   |\n  ---\n")
}