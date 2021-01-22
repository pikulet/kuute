package main

import (
    "fmt"

    "io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
)

var kuuteDB *KuuteDB

func main() {
//    gin.SetMode(gin.ReleaseMode)

    kuuteDB = InitKuuteDB()
    defer kuuteDB.shutdown()

    r := gin.New()
    r.GET("/:name", getCounter)
    r.Run()
}

func index (c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "https://github.com/pikulet/kuute")
    c.Abort()
}

func getCounter (c *gin.Context) {
    name := c.Param("name")
    count := kuuteDB.getCounter(name)
    img := getShieldsIOImage(count)

    c.Header("Content-Type", "image/svg+xml;charset=utf-8")
    c.String(http.StatusOK, img)
}

const (
    ShieldsIO string = "https://img.shields.io/badge/Views-%d-00bcc9"
)

func getShieldsIOImage(count int) string {
    site := fmt.Sprintf(ShieldsIO, count)
    resp, err := http.Get(site)
    if err != nil {
        panic (err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    return string(body)
}
