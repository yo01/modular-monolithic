package unit_test

import (
	"testing"

	"modular-monolithic/model"
	"modular-monolithic/module/v1/auth/helper"

	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/bcrypt"
)

func StringPointer(s string) *string {
	return &s
}

func TestHash(t *testing.T) {
	// Define a sample password
	password := "password123!"

	// Call the Hash function to hash the password
	hashedPassword, err := helper.Hash(password)
	if err != nil {
		t.Errorf("Hash function returned an error: %v", err)
	}

	// Compare the length of the hashed password
	if len(hashedPassword) == 0 {
		t.Error("Hashed password is empty")
	}

	// Compare the hashed password with the expected length
	if len(hashedPassword) != 60 {
		t.Errorf("Hashed password length is incorrect, got: %d, want: %d", len(hashedPassword), bcrypt.DefaultCost)
	}
}

func TestVerifyPassword(t *testing.T) {
	// Generate a hashed password from a known password
	knownPassword := "password123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(knownPassword), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Error generating hashed password: %v", err)
	}

	// Test verifying the correct password
	if err = helper.VerifyPassword(string(hashedPassword), knownPassword); err != nil {
		t.Errorf("VerifyPassword returned an error for correct password: %v", err)
	}

	// Test verifying an incorrect password
	incorrectPassword := "wrongpassword"
	if err = helper.VerifyPassword(string(hashedPassword), incorrectPassword); err == nil {
		t.Error("VerifyPassword did not return an error for incorrect password")
	}
}

func TestBycryptPassword(t *testing.T) {
	// Create a sample User data with a plain text password
	user := &model.User{
		Email:    "test@gmail.com",
		Password: StringPointer("plain_text_password"),
	}

	// Call the BycryptPassword function to hash the password
	if err := helper.BycryptPassword(user); err != nil {
		t.Fatalf("BycryptPassword returned an error: %v", err)
	}

	// Verify that the password in the User model is hashed
	assert.NotEqual(t, "plain_text_password", user.Password, "Password should be hashed")
}
