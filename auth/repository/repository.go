package repository

import (
	"fmt"

	"github.com/jacobsngoodwin/wordmem/auth/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // used for sqlx access to postgresql
)

// Create initializes and combines the various data resourses
// with concrete implementations having db instance injected
func Create(dataSourceName string) (*Repository, error) {
	db, err := sqlx.Open("postgres", dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("error opening db: %w", err)
	}

	// Verify database connection is working
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}

	return &Repository{
		UserRepository: UserRepositoryInit(db),
	}, nil

}

// Repository defines required interfaces to construct valid repository
type Repository struct {
	model.UserRepository
}
