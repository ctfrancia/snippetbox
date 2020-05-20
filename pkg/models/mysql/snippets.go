package mysql

import (
	"database/sql"

	"github.com/ctfrancia/snippetbox/pkg/models"
)

// SnippetModel struct
type SnippetModel struct {
	DB *sql.DB
}

// Insert inserts into DB
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?,?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// Get specific snippet
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest returns the last 10 snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
