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

  log.Print("   .\n  .\n . .\n  ...\n\\~~~~~/\n \\   /\n  \\ /\n   V\n   |\n   |\n  ---\n")
  r := NewRouter()
  r.Run()
}