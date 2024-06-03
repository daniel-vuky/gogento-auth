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

// CreateRandomRefreshToken
// Creates a random refresh token for testing
func CreateRandomRefreshToken(t *testing.T) db.RefreshToken {
	randomCustomer := CreateCustomer(t)
	arg := &db.InsertRefreshTokenParams{
		CustomerID:   randomCustomer.CustomerID,
		RefreshToken: util.RandomString(32),
		UserAgent:    util.RandomString(12),
		ClientIp:     util.RandomString(12),
		IsBlocked:    false,
		ExpiredAt:    util.RandomDate(),
	}
	refreshToken, err := testStore.InsertRefreshToken(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, refreshToken)
	require.Equal(t, arg.CustomerID, refreshToken.CustomerID)
	require.Equal(t, arg.RefreshToken, refreshToken.RefreshToken)
	require.Equal(t, arg.UserAgent, refreshToken.UserAgent)
	require.Equal(t, arg.ClientIp, refreshToken.ClientIp)
	require.Equal(t, arg.IsBlocked, refreshToken.IsBlocked)
	require.Equal(t, arg.ExpiredAt, refreshToken.ExpiredAt)
	require.NotEmpty(t, refreshToken.CreatedAt)

	return refreshToken
}

// TestDeleteRefreshToken
// Test the DeleteRefreshToken function
func TestDeleteRefreshToken(t *testing.T) {
	refreshToken := CreateRandomRefreshToken(t)
	deletedRefreshToken, err := testStore.DeleteRefreshToken(context.Background(), refreshToken.CustomerID)
	compareRefreshToken(t, err, &refreshToken, &deletedRefreshToken)
}

// TestGetRefreshToken
// Test the GetRefreshToken function
func TestGetRefreshToken(t *testing.T) {
	refreshToken := CreateRandomRefreshToken(t)
	existedRefreshToken, err := testStore.GetRefreshToken(context.Background(), refreshToken.CustomerID)
	compareRefreshToken(t, err, &refreshToken, &existedRefreshToken)
}

// compareRefreshToken
// Compare refresh token data
func compareRefreshToken(
	t *testing.T,
	err error,
	fromRefreshToken *db.RefreshToken,
	targetRefreshToken *db.RefreshToken,
) {
	require.NoError(t, err)
	require.NotEmpty(t, fromRefreshToken)
	require.NotEmpty(t, targetRefreshToken)
	require.Equal(t, fromRefreshToken.RefreshTokenID, targetRefreshToken.RefreshTokenID)
	require.Equal(t, fromRefreshToken.CustomerID, targetRefreshToken.CustomerID)
	require.Equal(t, fromRefreshToken.RefreshToken, targetRefreshToken.RefreshToken)
	require.Equal(t, fromRefreshToken.UserAgent, targetRefreshToken.UserAgent)
	require.Equal(t, fromRefreshToken.ClientIp, targetRefreshToken.ClientIp)
	require.Equal(t, fromRefreshToken.IsBlocked, targetRefreshToken.IsBlocked)
	require.Equal(t, fromRefreshToken.ExpiredAt, targetRefreshToken.ExpiredAt)
	require.Equal(t, fromRefreshToken.CreatedAt, targetRefreshToken.CreatedAt)
}

// TestInsertRefreshToken
// Test the InsertRefreshToken function
func TestInsertRefreshToken(t *testing.T) {
	refreshToken := CreateRandomRefreshToken(t)
	createdRefreshToken, err := testStore.GetRefreshToken(context.Background(), refreshToken.CustomerID)
	require.NoError(t, err)
	require.NotEmpty(t, createdRefreshToken)
	require.Equal(t, refreshToken.CustomerID, createdRefreshToken.CustomerID)
	require.Equal(t, refreshToken.RefreshToken, createdRefreshToken.RefreshToken)
	require.Equal(t, refreshToken.UserAgent, createdRefreshToken.UserAgent)
	require.Equal(t, refreshToken.ClientIp, createdRefreshToken.ClientIp)
	require.Equal(t, refreshToken.IsBlocked, createdRefreshToken.IsBlocked)
	require.Equal(t, refreshToken.ExpiredAt, createdRefreshToken.ExpiredAt)
	require.WithinDuration(t, refreshToken.CreatedAt, createdRefreshToken.CreatedAt, time.Second)
}

// TestUpdateRefreshToken
// Test the UpdateRefreshToken function
func TestUpdateRefreshToken(t *testing.T) {
	refreshToken := CreateRandomRefreshToken(t)
	updateRefreshTokenParam := &db.UpdateRefreshTokenParams{
		CustomerID: refreshToken.CustomerID,
		RefreshToken: pgtype.Text{
			String: util.RandomString(32),
			Valid:  true,
		},
		UserAgent: pgtype.Text{
			String: util.RandomString(12),
			Valid:  true,
		},
		ClientIp: pgtype.Text{
			String: util.RandomString(12),
			Valid:  true,
		},
		IsBlocked: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
		ExpiredAt: pgtype.Timestamptz{
			Time:  util.RandomDate(),
			Valid: true,
		},
	}
	updatedRefreshToken, err := testStore.UpdateRefreshToken(context.Background(), updateRefreshTokenParam)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRefreshToken)
	require.Equal(t, updateRefreshTokenParam.CustomerID, updatedRefreshToken.CustomerID)
	require.NotEqual(t, updateRefreshTokenParam.RefreshToken, updatedRefreshToken.RefreshToken)
	require.NotEqual(t, updateRefreshTokenParam.UserAgent, updatedRefreshToken.UserAgent)
	require.NotEqual(t, updateRefreshTokenParam.ClientIp, updatedRefreshToken.ClientIp)
	require.NotEqual(t, updateRefreshTokenParam.IsBlocked, updatedRefreshToken.IsBlocked)
	require.NotEqual(t, updateRefreshTokenParam.ExpiredAt, updatedRefreshToken.ExpiredAt)
	require.WithinDuration(t, updatedRefreshToken.CreatedAt, updatedRefreshToken.CreatedAt, time.Second)
}
