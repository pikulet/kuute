package main

import (
    "github.com/kataras/iris/v12"
)

var kuuteDB *KuuteDB

func main() {
    kuuteDB = InitKuuteDB()
    defer kuuteDB.shutdown()

    app := iris.New()
    app.Logger().SetLevel("info")

    app.Get("/", index) 
    app.Get("/{username:string regexp(^[a-zA-z0-9-]*$) max(39)}", getCounter) 
    app.Listen(":8080")
}

func index (c iris.Context) {
    c.Redirect("https://github.com/pikulet/kuute", iris.StatusPermanentRedirect)
}

func getCounter (c iris.Context) {
    user := c.Params().Get("username")
    c.Application().Logger().Infof("User: %s", user)

    kuuteDB.getCounter(user)
    //c.Text("%d", count)
}
