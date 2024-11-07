package main

import (
    "gin.ksk318.me/controller"
    "html/template"
    "log"
)

func main() {
    router := controller.GetRouter()
    router.Static("/assets", "./assets")

    // ポート8090で表示
    router.Run(":8090")
}

func loadTemplate(name string) *template.Template {
    t, err := template.ParseFiles(
        "view/_head.html",
        "view/_header.html",
    )
    if err != nil {
        log.Fatalf("template error: %v", err)
    }
    return t
}