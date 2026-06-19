package users

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

//TODO: Finish lookup method
func (r *Repository) FindByUserOrEmail(value string) (*User, error) {
	query := `SELECT id, username, email, password_hash, display_name, role, created_at
			  FROM users
			  WHERE username = $1 OR email = $1`

	row := r.db.QueryRow(query, value)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.DisplayName, &user.Role, &user.CreatedAt)
}
