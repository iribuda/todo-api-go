package configs

import (
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
)

// Erstellung von der Verbindung mit der Datenbank
func NewMySQLStorage(cfg mysql.Config)(*sql.DB, error){
	// öffnet Datenbank, die durch driver name und data source name (DSN) angegeben wird
	// DSN - username, password, host, port, database name
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}