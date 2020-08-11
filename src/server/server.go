package server

import (
  "github.com/gin-gonic/gin"
   "github.com/fallenstedt/gin-example/src/config"
)

func Init() {

  con, err := config.Init()

  if err != nil {
    panic("Failed to load config")
  }

  if con.Mode == "release" {
    gin.SetMode(gin.ReleaseMode)
  }
  

  r := NewRouter()
  r.Run()
}