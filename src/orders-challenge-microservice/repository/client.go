package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"sync"
)

type (
	DBClient interface {
		GetClient() (*sql.DB, error)
	}
	DBClientImpl struct {
	}
)

var (
	once          sync.Once
	dbInstance    *sql.DB
	errorInstance error
)

func NewDBClient() DBClient {
	return DBClientImpl{}
}

func (dbci DBClientImpl) GetClient() (*sql.DB, error) {
	return getClientInstance()
}

func getClientInstance() (*sql.DB, error) {
	once.Do(func() {
		psqlInfo := "host=localhost port=5432 user=dbuser password=admin2021 dbname=todoapp sslmode=disable" //TODO: env var
		dbInstance, errorInstance = sql.Open("postgres", psqlInfo)
	})
	return dbInstance, errorInstance
}
