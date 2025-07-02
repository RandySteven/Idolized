package repositories

import (
	"database/sql"
)

type Repositories struct {
}

func NewRepositories(db *sql.DB) *Repositories {
	_, _ = newTransaction(db)
	return &Repositories{}
}
