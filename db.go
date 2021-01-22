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
    Count   int         `sql:",notnull"`
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

    count := 1

    user := User {
        Name: username,
        Count: count,
    }

    // insert
    created, err := kdb.db.Model(&user).
        Where("name = ?name").
        OnConflict("DO NOTHING").
        SelectOrInsert()

    if err != nil {
        panic(err)
    }

    if !created {
        // update
        _, err = kdb.db.Model(&user).
            Where("name = ?name").
            Set("count = count + 1").
            Returning("count").
            Update(&count)
    }

    fmt.Println(count)

    return 0
}

func (kdb *KuuteDB) shutdown () {
    kdb.db.Close()
}
