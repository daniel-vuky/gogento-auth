package sqlc

import (
	"context"
	db "github.com/daniel-vuky/gogento-auth/db/sqlc"
	"github.com/daniel-vuky/gogento-auth/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// CreateCustomer
// Create a new customer with random data for testing
func CreateCustomer(t *testing.T) db.Customer {
	insertCustomerParam := &db.InsertCustomerParams{
		Email:     util.RandomEmail(),
		Firstname: util.RandomString(5),
		Lastname:  util.RandomString(5),
		Gender: db.NullGender{
			Gender: util.RandomGender(),
			Valid:  true,
		},
		Dob: pgtype.Timestamptz{
			Time:  util.RandomDate(),
			Valid: true,
		},
		HashedPassword:    util.RandomString(32),
		PasswordChangedAt: util.RandomDate(),
	}
	createdCustomer, err := testStore.InsertCustomer(context.Background(), insertCustomerParam)
	require.NoError(t, err)
	require.NotEmpty(t, createdCustomer)
	require.NotEmpty(t, createdCustomer.Email)
	require.NotEmpty(t, createdCustomer.Firstname)
	require.NotEmpty(t, createdCustomer.Lastname)
	require.NotEmpty(t, createdCustomer.Gender)
	require.NotEmpty(t, createdCustomer.Dob)
	require.NotEmpty(t, createdCustomer.HashedPassword)
	require.NotEmpty(t, createdCustomer.PasswordChangedAt)
	require.NotEmpty(t, createdCustomer.CreatedAt)

	return createdCustomer
}

// TestGetCustomer
// Test the GetCustomer function
func TestGetCustomer(t *testing.T) {
	customer := CreateCustomer(t)
	existedCustomer, err := testStore.GetCustomer(context.Background(), customer.Email)
	require.NoError(t, err)
	require.NotEmpty(t, existedCustomer)
	require.Equal(t, customer.Email, existedCustomer.Email)
	require.Equal(t, customer.Firstname, existedCustomer.Firstname)
	require.Equal(t, customer.Lastname, existedCustomer.Lastname)
	require.Equal(t, customer.Gender, existedCustomer.Gender)
	require.Equal(t, customer.Dob, existedCustomer.Dob)
	require.Equal(t, customer.HashedPassword, existedCustomer.HashedPassword)
	require.Equal(t, customer.PasswordChangedAt, existedCustomer.PasswordChangedAt)
	require.WithinDuration(t, customer.PasswordChangedAt, existedCustomer.PasswordChangedAt, time.Second)
}

// TestUpdateUser
// Test the UpdateUser function
func TestUpdateUser(t *testing.T) {
	customer := CreateCustomer(t)
	updateUserParam := &db.UpdateCustomerParams{
		Email: customer.Email,
		Firstname: pgtype.Text{
			String: util.RandomString(10),
			Valid:  true,
		},
		Lastname: pgtype.Text{
			String: util.RandomString(10),
			Valid:  true,
		},
		Gender: db.NullGender{
			Gender: util.RandomGender(),
			Valid:  true,
		},
		Dob: pgtype.Timestamptz{
			Time:  util.RandomDate(),
			Valid: true,
		},
		HashedPassword: pgtype.Text{
			String: util.RandomString(32),
			Valid:  true,
		},
		PasswordChangedAt: pgtype.Timestamptz{
			Time:  util.RandomDate(),
			Valid: true,
		},
	}
	updatedCustomer, err := testStore.UpdateCustomer(context.Background(), updateUserParam)
	require.NoError(t, err)
	require.NotEmpty(t, updatedCustomer)
	require.Equal(t, customer.Email, updatedCustomer.Email)
	require.NotEqual(t, customer.Firstname, updatedCustomer.Firstname)
	require.NotEqual(t, customer.Lastname, updatedCustomer.Lastname)
	require.NotEmpty(t, updatedCustomer.Gender)
	require.NotEqual(t, customer.Dob, updatedCustomer.Dob)
	require.NotEqual(t, customer.HashedPassword, updatedCustomer.HashedPassword)
	require.NotEqual(t, customer.PasswordChangedAt, updatedCustomer.PasswordChangedAt)
	require.WithinDuration(t, customer.CreatedAt, updatedCustomer.CreatedAt, time.Second)
}

func TestInsertCustomer(t *testing.T) {
	// Case success
	randomCustomer := CreateCustomer(t)

	// Case email is empty
	insertCustomerParam := &db.InsertCustomerParams{
		Email:     "",
		Firstname: util.RandomString(5),
		Lastname:  util.RandomString(5),
		Gender: db.NullGender{
			Gender: util.RandomGender(),
			Valid:  true,
		},
		Dob: pgtype.Timestamptz{
			Time:  util.RandomDate(),
			Valid: true,
		},
		HashedPassword:    util.RandomString(32),
		PasswordChangedAt: util.RandomDate(),
	}
	createdCustomer, err := testStore.InsertCustomer(context.Background(), insertCustomerParam)
	require.Error(t, err)
	require.Empty(t, createdCustomer)

	// Case email unique
	insertCustomerParam = &db.InsertCustomerParams{
		Email:     randomCustomer.Email,
		Firstname: util.RandomString(5),
		Lastname:  util.RandomString(5),
		Gender: db.NullGender{
			Gender: util.RandomGender(),
			Valid:  true,
		},
		Dob: pgtype.Timestamptz{
			Time:  util.RandomDate(),
			Valid: true,
		},
		HashedPassword:    util.RandomString(32),
		PasswordChangedAt: util.RandomDate(),
	}
	createdCustomer, err = testStore.InsertCustomer(context.Background(), insertCustomerParam)
	require.Error(t, err)
	require.Empty(t, createdCustomer)
}
