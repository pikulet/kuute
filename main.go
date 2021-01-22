package main

import (
    "fmt"

    "io/ioutil"
    "net/http"

    "time"
    "crypto/md5"

    "github.com/gin-gonic/gin"
)

var kuuteDB *KuuteDB

func main() {
    gin.SetMode(gin.ReleaseMode)

    kuuteDB = InitKuuteDB()
    defer kuuteDB.shutdown()

    router  := gin.Default()
    router.GET("/", index)
    router.GET("/:name", getCounter)

    server := & http.Server{
        Handler:    router,
    }

    server.SetKeepAlivesEnabled(false)
    server.ListenAndServe()
}

func index (c *gin.Context) {
    c.Redirect(http.StatusMovedPermanently, "https://github.com/pikulet/kuute")
    c.Abort()
}

func getCounter (c *gin.Context) {
    name := c.Param("name")
    count := kuuteDB.getCounter(name)
    img := getShieldsIOImage(count)

    // to block caching
    etag := fmt.Sprintf("%x", md5.Sum(img))
    c.Header("Cache-Control", "no-cache")
    c.Header("Content-Type", "image/svg+xml;charset=utf-8")
    c.Header("ETag", etag)
    c.Header("Expires", time.Now().UTC().Format(http.TimeFormat))
    c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
    c.String(http.StatusOK, string(img))
}

const (
    ShieldsIO string = "https://img.shields.io/badge/⭐ Views-%d-00bcc9"
)

func getShieldsIOImage(count int) []byte {
    site := fmt.Sprintf(ShieldsIO, count)
    resp, err := http.Get(site)
    if err != nil {
        panic (err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    return body
}
