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

// CreateRandomRule
// Create a random rule for testing
func CreateRandomRule(t *testing.T, roleID int64) db.AuthorizationRule {
	newRule := &db.InsertAuthorizationRuleParams{
		RoleID:          roleID,
		IsAdministrator: false,
		PermissionCode:  util.RandomString(10),
		IsAllowed:       true,
	}
	rule, err := testStore.InsertAuthorizationRule(context.Background(), newRule)
	compareRule(t, err, &db.AuthorizationRule{
		RoleID:          newRule.RoleID,
		IsAdministrator: newRule.IsAdministrator,
		PermissionCode:  newRule.PermissionCode,
		IsAllowed:       newRule.IsAllowed,
	}, &rule)

	return rule
}

// TestDeleteAuthorizationRule
// Test the DeleteAuthorizationRule function
func TestDeleteAuthorizationRule(t *testing.T) {
	role := CreateRandomRole(t)
	rule := CreateRandomRule(t, role.RoleID)
	deletedRule, err := testStore.DeleteAuthorizationRule(context.Background(), rule.RuleID)
	compareRule(t, err, &rule, &deletedRule)
}

// TestGetAuthorizationRule
// Test the GetAuthorizationRule function
func TestGetAuthorizationRule(t *testing.T) {
	role := CreateRandomRole(t)
	rule := CreateRandomRule(t, role.RoleID)
	existedRule, err := testStore.GetAuthorizationRule(context.Background(), rule.RuleID)
	compareRule(t, err, &rule, &existedRule)
}

// compareRule
// Compare rules data
func compareRule(
	t *testing.T,
	err error,
	fromRule *db.AuthorizationRule,
	targetRule *db.AuthorizationRule,
) {
	require.NoError(t, err)
	require.NotEmpty(t, fromRule)
	require.NotEmpty(t, targetRule)
	require.Equal(t, fromRule.RoleID, targetRule.RoleID)
	require.Equal(t, fromRule.IsAdministrator, targetRule.IsAdministrator)
	require.IsType(t, true, targetRule.IsAdministrator)
	require.Equal(t, fromRule.PermissionCode, targetRule.PermissionCode)
	require.Equal(t, fromRule.IsAllowed, targetRule.IsAllowed)
	require.IsType(t, true, targetRule.IsAllowed)
}

// TestGetAuthorizationRuleByRole
// Test the GetAuthorizationRuleByRole function
func TestGetAuthorizationRuleByRole(t *testing.T) {
	listRule := make([]db.AuthorizationRule, 10)
	randomRole := CreateRandomRole(t)
	for i := 0; i < 10; i++ {
		listRule[i] = CreateRandomRule(t, randomRole.RoleID)
	}
	listCreatedRule, err := testStore.GetAuthorizationRuleByRole(context.Background(), randomRole.RoleID)
	require.NoError(t, err)
	require.NotEmpty(t, listCreatedRule)
	require.Equal(t, len(listCreatedRule), 10)

	for index, rule := range listCreatedRule {
		compareRule(t, nil, &listRule[index], &rule)
	}
}

// TestInsertAuthorizationRule
// Test the InsertAuthorizationRule function
func TestInsertAuthorizationRule(t *testing.T) {
	// Normal case
	role := CreateRandomRole(t)
	newRule := &db.InsertAuthorizationRuleParams{
		RoleID:          role.RoleID,
		IsAdministrator: false,
		PermissionCode:  util.RandomString(10),
		IsAllowed:       true,
	}
	rule, err := testStore.InsertAuthorizationRule(context.Background(), newRule)
	compareRule(t, err, &db.AuthorizationRule{
		RoleID:          newRule.RoleID,
		IsAdministrator: newRule.IsAdministrator,
		PermissionCode:  newRule.PermissionCode,
		IsAllowed:       newRule.IsAllowed,
	}, &rule)

	// Case RoleID and PermissionCode unique
	newRule = &db.InsertAuthorizationRuleParams{
		RoleID:          role.RoleID,
		IsAdministrator: false,
		PermissionCode:  rule.PermissionCode,
		IsAllowed:       true,
	}
	rule, err = testStore.InsertAuthorizationRule(context.Background(), newRule)
	require.Error(t, err)
	require.Empty(t, rule)
}

// TestInsertMultipleAuthorizationRule
// Test the InsertMultipleAuthorizationRule function
func TestInsertMultipleAuthorizationRule(t *testing.T) {
	var listRoleId []int64
	var listPermissionCode []string
	var listIsAdministrator, listIsAllowed []bool
	role := CreateRandomRole(t)
	for i := 0; i < 10; i++ {
		listRoleId = append(listRoleId, role.RoleID)
		listPermissionCode = append(listPermissionCode, util.RandomString(20))
		listIsAdministrator = append(listIsAdministrator, false)
		listIsAllowed = append(listIsAllowed, i%2 == 0)
	}
	insertMultipleParams := &db.InsertMultipleAuthorizationRulesParams{
		listRoleId,
		listIsAdministrator,
		listPermissionCode,
		listIsAllowed,
	}
	listCreatedRules, err := testStore.InsertMultipleAuthorizationRules(context.Background(), insertMultipleParams)
	require.NoError(t, err)
	require.Equal(t, len(listCreatedRules), 10)
	for index, rule := range listCreatedRules {
		compareRule(t, nil, &db.AuthorizationRule{
			RoleID:          listRoleId[index],
			IsAdministrator: listIsAdministrator[index],
			PermissionCode:  listPermissionCode[index],
			IsAllowed:       listIsAllowed[index],
		}, &rule)
	}
}

// TestIsAllowed
func TestIsAllowed(t *testing.T) {
	// Case allowed
	role := CreateRandomRole(t)
	rule := CreateRandomRule(t, role.RoleID)
	isAllowedParams := &db.IsAllowedParams{
		RoleID:         role.RoleID,
		PermissionCode: rule.PermissionCode,
	}
	isAllowed, err := testStore.IsAllowed(context.Background(), isAllowedParams)
	require.NoError(t, err)
	require.Equal(t, rule.IsAllowed, isAllowed)
}

// TestUpdateAuthorizationRule
// Test the UpdateAuthorizationRule function
func TestUpdateAuthorizationRule(t *testing.T) {
	role := CreateRandomRole(t)
	rule := CreateRandomRule(t, role.RoleID)
	newRule := &db.UpdateAuthorizationRuleParams{
		RuleID: rule.RuleID,
		RoleID: pgtype.Int8{
			Int64: role.RoleID,
			Valid: true,
		},
		IsAdministrator: pgtype.Bool{
			Bool:  true,
			Valid: true,
		},
		PermissionCode: pgtype.Text{
			String: util.RandomString(20),
			Valid:  true,
		},
		IsAllowed: pgtype.Bool{
			Bool:  false,
			Valid: true,
		},
	}
	updatedRule, err := testStore.UpdateAuthorizationRule(context.Background(), newRule)
	require.NoError(t, err)
	require.NotEmpty(t, updatedRule)
	require.Equal(t, newRule.RuleID, updatedRule.RuleID)
	require.NotEqual(t, newRule.IsAdministrator, updatedRule.IsAdministrator)
	require.IsType(t, true, updatedRule.IsAdministrator)
	require.NotEqual(t, newRule.PermissionCode, updatedRule.PermissionCode)
	require.NotEqual(t, newRule.IsAllowed, updatedRule.IsAllowed)
	require.IsType(t, true, updatedRule.IsAllowed)
	require.NotEmpty(t, updatedRule.CreatedAt)
	require.WithinDuration(t, rule.CreatedAt, updatedRule.CreatedAt, time.Second)
}
