package repository

import (
	"database/sql"
)

type ExampleRepository struct {
	db *sql.DB
}

func NewExampleRepository(db *sql.DB) *ExampleRepository {
	return &ExampleRepository{
		db: db,
	}
}

func (r *ExampleRepository) GetAll() ([]string, error) {
	q := `SELECT name FROM example`

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}

	var names []string
	for rows.Next() {
		var name string 
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	rows.Close()

	return names, nil
}
