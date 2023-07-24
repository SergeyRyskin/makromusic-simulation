package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

// User struct to describe user object.
type User struct {
    ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
    CreatedAt  time.Time `db:"created_at" json:"created_at"`
    UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
    UserID     uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
    Name       string    `db:"name" json:"name" validate:"required,lte=255"`
    Surname    string    `db:"surname" json:"surname" validate:"required,lte=255"`
    Email      string    `db:"email" json:"email" validate:"required,lte=255"`
    Password      string    `db:"password" json:"password" validate:"required,lte=255"`
    UserStatus int       `db:"user_status" json:"user_status" validate:"required,len=1"`
    UserAttrs  UserAttrs `db:"user_attrs" json:"user_attrs" validate:"required,dive"`
}

// UserAttrs struct to describe user attributes.
type UserAttrs struct {
    Picture     string `json:"picture"`
    Description string `json:"description"`
    Rating      int    `json:"rating" validate:"min=1,max=10"`
}

// Value make the UserAttrs struct implement the driver.Valuer interface.
// This method simply returns the JSON-encoded representation of the struct.
func (b UserAttrs) Value() (driver.Value, error) {
    return json.Marshal(b)
}

// Scan make the UserAttrs struct implement the sql.Scanner interface.
// This method simply decodes a JSON-encoded value into the struct fields.
func (b *UserAttrs) Scan(value interface{}) error {
    j, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(j, &b)
}