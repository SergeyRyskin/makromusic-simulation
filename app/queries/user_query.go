package queries

import (
    "github.com/google/uuid"
    "github.com/jmoiron/sqlx"
    "github.com/SergeyRyskin/makromusic-simulation/app/models"
)

// UserQueries struct for queries from User model.
type UserQueries struct {
    *sqlx.DB
}

// GetUsers method for getting all users.
func (q *UserQueries) GetUsers() ([]models.User, error) {
    // Define users variable.
    users := []models.User{}

    // Define query string.
    query := `SELECT * FROM users`

    // Send query to database.
    err := q.Get(&users, query)
    if err != nil {
        // Return empty object and error.
        return users, err
    }

    // Return query result.
    return users, nil
}

// GetUser method for getting one user by given ID.
func (q *UserQueries) GetUser(id uuid.UUID) (models.User, error) {
    // Define user variable.
    user := models.User{}

    // Define query string.
    query := `SELECT * FROM users WHERE id = $1`

    // Send query to database.
    err := q.Get(&user, query, id)
    if err != nil {
        // Return empty object and error.
        return user, err
    }

    // Return query result.
    return user, nil
}

// CreateUser method for creating user by given User object.
func (q *UserQueries) CreateUser(b *models.User) error {
    // Define query string.
    query := `INSERT INTO users VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

    // Send query to database.
    _, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Name, b.Surname, b.Email, b.Password, b.UserStatus, b.UserAttrs)
    if err != nil {
        // Return only error.
        return err
    }

    // This query returns nothing.
    return nil
}

// UpdateUser method for updating user by given User object.
func (q *UserQueries) UpdateUser(id uuid.UUID, b *models.User) error {
    // Define query string.
    query := `UPDATE users SET updated_at = $2, name = $3, surname = $4, email = $5, password = $6, user_status = $7, user_attrs = $8 WHERE id = $1`

    // Send query to database.
    _, err := q.Exec(query, id, b.UpdatedAt, b.Name, b.Surname, b.Email, b.Password, b.UserStatus, b.UserAttrs)
    if err != nil {
        // Return only error.
        return err
    }

    // This query returns nothing.
    return nil
}

// DeleteUser method for delete user by given ID.
func (q *UserQueries) DeleteUser(id uuid.UUID) error {
    // Define query string.
    query := `DELETE FROM users WHERE id = $1`

    // Send query to database.
    _, err := q.Exec(query, id)
    if err != nil {
        // Return only error.
        return err
    }

    // This query returns nothing.
    return nil
}