//import (
//	"database/sql"
//	"time"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
// ...
//
// db, err := sql.Open("mysql", "user:password@/dbname")
// if err != nil {
//	panic(err)
//}
// See "Important settings" section.
// db.SetConnMaxLifetime(time.Minute * 3)
// db.SetMaxOpenConns(10)
// db.SetMaxIdleConns(10)

package main

import (
         "database/sql"
         "fmt"
         "log"
         "os"
	"github.com/go-sql-driver/mysql"
)
var db *sql.DB

func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}
