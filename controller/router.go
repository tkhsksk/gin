package controller

import (
    "github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
    router := gin.Default()
    router.LoadHTMLGlob("view/*")

    router.GET("/", Home)
    router.GET("/user", CreateUser)

    return router
}