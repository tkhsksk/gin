package controller

import (
    "gin.ksk318.me/model"
    "strconv"

    "github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
    datas := model.GetAll()
    c.HTML(200, "user.html", gin.H{"datas": datas})
}

func CreateProcess(c *gin.Context) {
    // postを受け取ってdbに登録する
    name := c.PostForm("name")
    data := model.User{Name: name}
    data.Create()
    c.Redirect(301, "/")
}

func EditUser(c *gin.Context) {
    // パラメータのidからデータを取得
    id, _ := strconv.Atoi(c.Param("id"))
    data := model.GetOne(id)
    // dataを受け渡し
    c.HTML(200, "edit.html", gin.H{"data": data})
}

func EditProcess(c *gin.Context) {
    // idからデータを取得
    id, _ := strconv.Atoi(c.PostForm("id"))
    data := model.GetOne(id)
    // nameをpostから取得して設定
    name := c.PostForm("name")
    data.Name = name
    // dataをupdate
    data.Update()
    c.Redirect(301, "/")
}

func SearchUser(c *gin.Context) {
    name := c.Query("name")
    datas := model.Search(name)
    c.HTML(200, "search.html", gin.H{"datas":datas, "query":name})
}

