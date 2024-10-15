package controller

import (
    "gin.ksk318.me/model"
    "strconv"

    "github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    c.HTML(200, "user.html", gin.H{})
}

func CreateProcess(c *gin.Context) {
    name := c.PostForm("name")
    data := model.UserEntity{Name: name}
    data.Create()
    c.Redirect(301, "/")
}

func EditUser(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    data := model.GetOne(id)
    c.HTML(200, "edit.html", gin.H{"data": data})
}

func EditProcess(c *gin.Context) {
    id, _ := strconv.Atoi(c.PostForm("id"))
    data := model.GetOne(id)
    name := c.PostForm("name")
    data.Name = name
    data.Update()
    c.Redirect(301, "/")
}
