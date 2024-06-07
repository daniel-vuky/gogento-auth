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

// CreateRandomRole
// Create a new role with random data for testing
func CreateRandomRole(t *testing.T) db.AuthorizationRole {
	newRole := &db.InsertAuthorizationRoleParams{
		RoleName: util.RandomString(6),
		Description: pgtype.Text{
			String: util.RandomString(32),
			Valid:  true,
		},
	}

	createdRole, err := testStore.InsertAuthorizationRole(context.Background(), newRole)
	require.NoError(t, err)
	require.NotEmpty(t, createdRole)
	require.Equal(t, newRole.RoleName, createdRole.RoleName)
	require.Equal(t, newRole.Description, createdRole.Description)
	require.NotEmpty(t, createdRole.CreatedAt)

	return createdRole
}

// TestCreateAuthorizationRole
// Test the DeleteAuthorizationRole function
func TestDeleteAuthorizationRole(t *testing.T) {
	role := CreateRandomRole(t)
	deletedRole, err := testStore.DeleteAuthorizationRole(context.Background(), role.RoleID)
	compareRole(t, err, &role, &deletedRole)
}

// TestGetAuthorizationRole
// Test the GetAuthorizationRole function
func TestGetAuthorizationRole(t *testing.T) {
	role := CreateRandomRole(t)
	existedRole, err := testStore.GetAuthorizationRole(context.Background(), role.RoleID)
	compareRole(t, err, &role, &existedRole)
}

// compareRole
// Compare role data
func compareRole(
	t *testing.T,
	err error,
	fromRole *db.AuthorizationRole,
	targetRole *db.AuthorizationRole,
) {
	require.NoError(t, err)
	require.NotEmpty(t, fromRole)
	require.NotEmpty(t, targetRole)
	require.Equal(t, fromRole.RoleID, targetRole.RoleID)
	require.Equal(t, fromRole.RoleName, targetRole.RoleName)
	require.Equal(t, fromRole.Description, targetRole.Description)
	require.WithinDuration(t, fromRole.CreatedAt, targetRole.CreatedAt, time.Second)
}

// TestInsertAuthorizationRole
// Test the InsertAuthorizationRole function
func TestInsertAuthorizationRole(t *testing.T) {
	CreateRandomRole(t)
}

// TestUpdateAuthorizationRole
// Test the UpdateAuthorizationRole function
func TestUpdateAuthorizationRole(t *testing.T) {
	role := CreateRandomRole(t)
	updatedParams := &db.UpdateAuthorizationRoleParams{
		RoleID: role.RoleID,
		RoleName: pgtype.Text{
			String: util.RandomString(12),
			Valid:  true,
		},
		Description: pgtype.Text{
			String: util.RandomString(64),
			Valid:  true,
		},
	}
	updatedRole, err := testStore.UpdateAuthorizationRole(context.Background(), updatedParams)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRole)
	require.Equal(t, role.RoleID, updatedRole.RoleID)
	require.NotEqual(t, role.RoleName, updatedRole.RoleName)
	require.NotEqual(t, role.Description, updatedRole.Description)
	require.WithinDuration(t, role.CreatedAt, updatedRole.CreatedAt, time.Second)
}
