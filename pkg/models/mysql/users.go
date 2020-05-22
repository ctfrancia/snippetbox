package mysql

import (
	"database/sql"

	"github.com/ctfrancia/snippetbox/pkg/models"
)

// UserModel defines the structure
type UserModel struct {
	DB *sql.DB
}

// Insert into the Users table
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Authenticate is used to verify whether a user exists with the provided email address and password
// returning the user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// Get fetches the details of the specific user based on their user ID
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
