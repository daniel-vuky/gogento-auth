package util

import (
	db "github.com/daniel-vuky/gogento-auth/db/sqlc"
	"math/rand"
	"time"
)

const (
	alphabet          string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	genderMale        string = "1"
	genderFemale      string = "2"
	genderNotProvided string = "3"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString
// Generates a random string of length n
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// RandomInt
// Generates a random integer between min and max
func RandomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

// RandomEmail
// Generates a random email
func RandomEmail() string {
	return RandomString(5) + "@unit_test.com"
}

// RandomGender
// Generates random gender base on enum value
func RandomGender() db.Gender {
	listGender := []db.Gender{
		db.Gender1,
		db.Gender2,
		db.Gender3,
	}

	return listGender[rand.Intn(len(listGender))]
}

// RandomDate
// Generates a random date
func RandomDate() time.Time {
	return time.Now().AddDate(-RandomInt(18, 50), 0, 0)
}
