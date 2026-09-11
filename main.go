package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.New()

	r.Use(gin.Recovery())

	configurarRotas(r)

	r.Run(":8080")
}
