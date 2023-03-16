package database

import "fmt"

type databaseModule interface {
	Query(query string) []map[string]interface{}
}

type databaseImpl struct{}

func (dm *databaseImpl) Query(query string) []map[string]interface{} {
	fmt.Printf("[DB] Querying: %s\n", query)
	// Perform database query here...
	return []map[string]interface{}{}
}

var databaseInstance *databaseImpl

func GetDatabaseModule() databaseModule {
	if databaseInstance == nil {
		databaseInstance = &databaseImpl{}
	}
	return databaseInstance
}
