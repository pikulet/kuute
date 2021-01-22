package kuute

import (
    "fmt"

    "os"
    "github.com/joho/godotenv"

    "github.com/go-pg/pg/v10"
)

var db *pg.DB

type User struct {
    tableName struct {} `pg:"kuute"`
    Id      int         `pg:",pk"`
    Name    string
    Count   int         `sql:",notnull"`
}

func (u User) String() string {
    return fmt.Sprintf("User<%s %d>", u.Name, u.Count)
}

func Init() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    options, _ := pg.ParseURL(os.Getenv("DATABASE_URL"))
    db := pg.Connect(options)

    // Check connection
    _, err = db.Exec("SELECT 1")
    if err != nil {
        panic(err)
    }

}

func Shutdown() {
    db.Close()
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

    return count
}

