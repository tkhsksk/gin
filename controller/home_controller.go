package controller

import (
    "github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
    // トップページ
    c.HTML(200, "index.html", gin.H{"title": "ホーム",})
}
