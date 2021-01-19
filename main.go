package main

import (
    "github.com/kataras/iris/v12"
)

func main() {
    app := iris.New()
    app.Get("/", index)
    app.Get("/{username:string}", getCounter)
    app.Listen()
}

func index(c iris.Context) {
    c.Redirect("https://github.com/pikulet/kuute", iris.StatusPermanentRedirect)
}

func getCounter(c iris.Context) {

}
