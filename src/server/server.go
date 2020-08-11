package server

import (
  "github.com/gin-gonic/gin"
   "github.com/fallenstedt/gin-example/src/config"
)

func Init() {
  c := config.Get()
  if c.Mode == "release" {
    gin.SetMode(gin.ReleaseMode)
  }

  r := NewRouter()
  r.Run()
}