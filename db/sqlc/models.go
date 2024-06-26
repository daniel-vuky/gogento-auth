// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type Gender string

const (
	Gender1 Gender = "1"
	Gender2 Gender = "2"
	Gender3 Gender = "3"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender `json:"gender"`
	Valid  bool   `json:"valid"` // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

type AdminUser struct {
	AdminID        int64              `json:"admin_id"`
	RoleID         int64              `json:"role_id"`
	Email          string             `json:"email"`
	Firstname      string             `json:"firstname"`
	Lastname       string             `json:"lastname"`
	HashedPassword string             `json:"hashed_password"`
	IsActive       bool               `json:"is_active"`
	LockExpires    pgtype.Timestamptz `json:"lock_expires"`
	CreatedAt      time.Time          `json:"created_at"`
}

type AuthorizationRole struct {
	RoleID      int64       `json:"role_id"`
	RoleName    string      `json:"role_name"`
	Description pgtype.Text `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
}

type AuthorizationRule struct {
	RuleID          int64     `json:"rule_id"`
	RoleID          int64     `json:"role_id"`
	IsAdministrator bool      `json:"is_administrator"`
	PermissionCode  string    `json:"permission_code"`
	IsAllowed       bool      `json:"is_allowed"`
	CreatedAt       time.Time `json:"created_at"`
}

type Customer struct {
	CustomerID        int64              `json:"customer_id"`
	Email             string             `json:"email"`
	Firstname         string             `json:"firstname"`
	Lastname          string             `json:"lastname"`
	Gender            NullGender         `json:"gender"`
	Dob               pgtype.Timestamptz `json:"dob"`
	HashedPassword    string             `json:"hashed_password"`
	PasswordChangedAt time.Time          `json:"password_changed_at"`
	CreatedAt         time.Time          `json:"created_at"`
}

type RefreshToken struct {
	RefreshTokenID int64     `json:"refresh_token_id"`
	CustomerID     int64     `json:"customer_id"`
	RefreshToken   string    `json:"refresh_token"`
	UserAgent      string    `json:"user_agent"`
	ClientIp       string    `json:"client_ip"`
	IsBlocked      bool      `json:"is_blocked"`
	ExpiredAt      time.Time `json:"expired_at"`
	CreatedAt      time.Time `json:"created_at"`
}
