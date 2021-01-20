package main

import (
    "fmt"

    "os"
    "github.com/joho/godotenv"

    "github.com/go-pg/pg/v10"
)

type User struct {
    tableName struct {} `pg:"kuute"`
    Id      int         `pg:",pk"`
    Name    string
    Count   int
}

func (u User) String() string {
    return fmt.Sprintf("User<%s %d>", u.Name, u.Count)
}

type KuuteDB struct {
    db      *pg.DB
}

func InitKuuteDB() *KuuteDB {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    
    db := pg.Connect(&pg.Options{
        User:       os.Getenv("DB_USER"), 
        Password:   os.Getenv("DB_PASSWORD"),
        Database:   os.Getenv("DB_DATABASE"),
        Addr:       os.Getenv("DB_ADDR"),
    })

    // Check connection
    _, err = db.Exec("SELECT 1")
    if err != nil {
        panic(err)
    }

    return &KuuteDB{ db }
}

func (kdb *KuuteDB) getCounter (username string) int {

    user := new(User)
    err := kdb.db.Model(user).
        Column("count").
        Where("name = ?", username).
        Select()

    if err != nil {
        panic(err)
    }
    
    fmt.Println(user)

    return 0
}

func (kdb *KuuteDB) shutdown () {
    kdb.db.Close()
}
