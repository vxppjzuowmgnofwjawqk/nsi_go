package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"nsi_go/internal/scheme"
)

type storage struct {
	db *sql.DB
}

type Storage interface {
	GetAllCountries() []scheme.NSI
	GetCountryById(id string) scheme.NSI
	CreateCountry(name string) scheme.NSI
	UpdateCountry(id, name string)
	DeleteCountry(id string)
	GetAllLanguages() []scheme.NSI
	GetLanguageById(id string) scheme.NSI
	CreateLanguage(name string) scheme.NSI
	UpdateLanguage(id, name string)
	DeleteLanguage(id string)
	GetAllStatuses() []scheme.NSI
	GetStatusById(id string) scheme.NSI
	CreateStatus(name string) scheme.NSI
	UpdateStatus(id, name string)
	DeleteStatus(id string)
}

func New() *storage {
	config := "user=postgres " +
		"password=130263 " +
		"host=localhost " +
		"port=5432 " +
		"dbname=nsi_go " +
		"sslmode=disable"
	db, err := sql.Open("postgres", config)
	if err != nil {
		fmt.Printf("an error occurred while opening the database:\n%v\n", err)
	}
	return &storage{
		db: db,
	}
}
