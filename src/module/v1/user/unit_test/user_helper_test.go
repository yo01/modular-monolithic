package unit_test

import (
	"testing"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/user/helper"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Helper function to create a pointer to a string
func StringPointer(s string) *string {
	return &s
}

func TestPrepareToDetailUserResponse(t *testing.T) {
	// Create sample data for testing
	userData := []model.User{
		{
			ID:       uuid.New(),
			Email:    "user@example.com",
			FullName: "John Doe",
			RoleID:   uuid.New(),
			RoleName: StringPointer("Admin"),
		},
		// Add more sample data as needed
	}

	// Call the function with the sample data
	result := helper.PrepareToDetailUserResponse(userData)

	// Assert the expected values based on the sample data
	assert.NotNil(t, result)
	assert.Equal(t, userData[0].ID, result.ID)
	assert.Equal(t, userData[0].Email, result.Email)
	assert.Equal(t, userData[0].FullName, result.FullName)

	// Assert the expected RoleResponse values if RoleID is not nil
	if userData[0].RoleID != uuid.Nil {
		assert.NotNil(t, result.Role)
		assert.Equal(t, userData[0].RoleID, result.Role.ID)
		assert.Equal(t, *userData[0].RoleName, result.Role.Name)
	} else {
		assert.Nil(t, result.Role)
	}
}

func TestPrepareToLoginDetailUserResponse(t *testing.T) {
	// Prepare sample data for testing
	userData := []model.User{
		{
			ID:       uuid.New(),
			Email:    "test@example.com",
			FullName: "Test User",
			Password: StringPointer("testpassword"),
			RoleID:   uuid.New(),
			RoleName: StringPointer("Admin"),
		},
	}

	// Call the function with sample data
	resp := helper.PrepareToLoginDetailUserResponse(userData)

	// Perform assertions to ensure the response is as expected
	assert.NotNil(t, resp)
	assert.Equal(t, userData[0].ID, resp.ID)
	assert.Equal(t, userData[0].Email, resp.Email)
	assert.Equal(t, userData[0].FullName, resp.FullName)
	assert.Equal(t, *userData[0].Password, resp.Password)

	// Assert RoleResponse if RoleID is not nil
	if userData[0].RoleID != uuid.Nil {
		assert.NotNil(t, resp.Role)
		assert.Equal(t, userData[0].RoleID, resp.Role.ID)
		assert.Equal(t, *userData[0].RoleName, resp.Role.Name)
	} else {
		assert.Nil(t, resp.Role)
	}
}

func TestPrepareToUsersResponse(t *testing.T) {
	// Define test data
	testData := []model.User{
		{
			ID:       uuid.New(),
			FullName: "John Doe",
			Email:    "john.doe@example.com",
			RoleID:   uuid.New(),
			RoleName: &[]string{"Admin"}[0],
		},
		{
			ID:       uuid.New(),
			FullName: "Jane Doe",
			Email:    "jane.doe@example.com",
			RoleID:   uuid.New(),
			RoleName: &[]string{"User"}[0],
		},
	}

	// Call the function with test data
	resp := helper.PrepareToUsersResponse(testData)

	// Assert the expected result
	require.Len(t, resp, 2)
	assert.Equal(t, testData[0].ID, resp[0].ID)
	assert.Equal(t, testData[0].FullName, resp[0].FullName)
	assert.Equal(t, testData[0].Email, resp[0].Email)
	assert.Equal(t, testData[0].RoleID, resp[0].Role.ID)
	assert.Equal(t, *testData[0].RoleName, resp[0].Role.Name)

	assert.Equal(t, testData[1].ID, resp[1].ID)
	assert.Equal(t, testData[1].FullName, resp[1].FullName)
	assert.Equal(t, testData[1].Email, resp[1].Email)
	assert.Equal(t, testData[1].RoleID, resp[1].Role.ID)
	assert.Equal(t, *testData[1].RoleName, resp[1].Role.Name)
}
