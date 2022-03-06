package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
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
		host := os.Getenv("dbHost")
		port := os.Getenv("dbPort")
		password := os.Getenv("dbPassword")
		name := os.Getenv("dbName")
		user := os.Getenv("dbUser")
		if host == "" || port == "" || password == "" || name == "" || user == "" {
			errorInstance = errors.New("database en vars are not populated")
		} else {
			psqlInfo := fmt.Sprintf("host=%+v port=%+v user=%+v password=%+v dbname=%+v sslmode=disable",
				host, port, user, password, name)
			dbInstance, errorInstance = sql.Open("postgres", psqlInfo)
		}
	})
	return dbInstance, errorInstance
}
