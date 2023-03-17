package database

import (
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
)

type databaseModule interface {
	Query(query string) []map[string]interface{}
}

type databaseImpl struct {
	db            *sqlx.DB
	maxConnection int
	mutex         sync.Mutex
}

func (dm *databaseImpl) Query(query string) []map[string]interface{} {
	dm.mutex.Lock()
	defer dm.mutex.Unlock()

	fmt.Printf("[DB] Database URL: %s\n", dm.db.DriverName())
	// Perform database query here...
	return []map[string]interface{}{}
}

var databaseInstance *databaseImpl

func GetDatabaseModule(databaseURL string, maxConnections int) databaseModule {
	if databaseInstance == nil {
		db, err := sqlx.Open("postgres", databaseURL)
		if err != nil {
			panic("Error connecting to database: " + err.Error())
		}
		databaseInstance = &databaseImpl{
			db:            db,
			maxConnection: maxConnections,
		}
	}
	return databaseInstance
}
