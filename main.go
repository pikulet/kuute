package kuute

import (
    "fmt"

    "io/ioutil"
    "net/http"
)

const (
    ShieldsIO string = "https://img.shields.io/badge/⭐ Views-%d-00bcc9"
)

func GetCounterBadge(name string) []byte {
    count := db.getCounter(name)
    return getShieldsIOBadge(count)
}

func getShieldsIOBadge(count int) []byte {
    site := fmt.Sprintf(ShieldsIO, count)
    resp, err := http.Get(site)
    if err != nil {
        panic (err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    return body
}
