package entities

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int64          `json:"id" db:"id"`
	Email     string         `json:"email" db:"email"`
	Password  string         `json:"-" db:"password"`
	Token     sql.NullString `json:"-" db:"fb_token"`
	RoleID    int16          `json:"roles_id" db:"roles_id"`
	IsBlocked int16          `json:"blocked,omitempty" db:"blocked"`
	CreatedAt mysql.NullTime `json:"-" db:"created_at"`
	UpdatedAt mysql.NullTime `json:"-" db:"updated_at"`
}
