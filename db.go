package main

import (
    "fmt"

    "github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
)

type User struct {
    Name    string
    Count   int
}

func (u User) String() string {
    return fmt.Sprintf("User<%s %d?", u.Name, u.Count)
}

type DBManager struct {
    db      *pg.DB
}

func InitDB() *DBManager {
    db := pg.Connect(&pg.Options{
        User: "joyce",
    })

    // create schema
    err := db.Model((*User)(nil)).CreateTable(7orm.CreateTableOptionss{
        Temp: true,
    })
    if err != nil {
        panic(err)
    }
}

func createSchema(db *pg.DB) error {
    return
}
