package util

import (
	db "github.com/daniel-vuky/gogento-auth/db/sqlc"
	"github.com/daniel-vuky/gogento-auth/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestRandomString
// Test RandomString function
func TestRandomString(t *testing.T) {
	randomString := util.RandomString(5)
	require.NotEmpty(t, randomString)
	require.Equal(t, 5, len(randomString))
}

// TestRandomInt
// Test RandomInt function
func TestRandomInt(t *testing.T) {
	randomInt := util.RandomInt(1, 10)
	require.GreaterOrEqual(t, randomInt, 1)
	require.LessOrEqual(t, randomInt, 10)
}

// TestRandomEmail
// Test RandomEmail function
func TestRandomEmail(t *testing.T) {
	randomEmail := util.RandomEmail()
	require.NotEmpty(t, randomEmail)
	require.Contains(t, randomEmail, "@unit_test.com")
}

// TestRandomGender
// Test RandomGender function
func TestRandomGender(t *testing.T) {
	randomGender := util.RandomGender()
	require.NotEmpty(t, randomGender)
	require.IsType(t, db.Gender(""), randomGender)
}

// TestRandomDate
// Test RandomDate function
func TestRandomDate(t *testing.T) {
	randomDate := util.RandomDate()
	require.NotEmpty(t, randomDate)
	require.IsType(t, time.Time{}, randomDate)
}
