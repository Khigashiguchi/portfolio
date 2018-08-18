package infrastructure

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql" // mysql driverを使うため
    "fmt"
)

// NewDB return *database.sql.DB struct.
func NewDB(conf Config) (*sql.DB, error) {
    // Open Database
    db, err := sql.Open(
        "mysql",
        fmt.Sprintf(
            "%s:%s@tcp(%s:%d)/%s?sql_mode=''%s'",
            conf.DB.User,
            conf.DB.Password,
            conf.DB.Host,
            conf.DB.Port,
            conf.DB.Name,
            conf.DB.SQLMode))
    if err != nil {
        return nil, err
    }
    if err := db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}
