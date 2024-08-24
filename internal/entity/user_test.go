package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "j@j.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestValidatePassword(t *testing.T) {
	user, _ := NewUser("John Doe", "j@j.com", "123456")
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("12345"))
	assert.NotEqual(t, user.Password, "123456")
}
