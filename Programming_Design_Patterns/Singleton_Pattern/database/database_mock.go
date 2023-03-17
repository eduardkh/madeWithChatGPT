package database

type databaseMock struct{}

func (dm *databaseMock) Query(query string) []map[string]interface{} {
	// Return mock data
	return []map[string]interface{}{
		{"id": 1, "name": "Alice"},
		{"id": 2, "name": "Bob"},
		{"id": 3, "name": "Charlie"},
	}
}

// GetDatabaseMock returns a databaseModule that returns mock data
func GetDatabaseMock() databaseModule {
	return &databaseMock{}
}
