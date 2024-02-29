package application

import (
	"errors"
	"testing"
	"user-svc/internal/domain/customer"
	"user-svc/internal/infra/database"
	"user-svc/internal/infra/encryption"

	"github.com/stretchr/testify/assert"
)

func TestCustomerSignup(t *testing.T) {
	c := []customer.Customer{
		{
			ID:       "existing_id",
			Name:     "existing_name",
			Email:    "existing_email",
			Password: "existing_password",
		},
	}

	repo := database.NewInMemoryCustomerRepo(c)
	hasher := encryption.NewCustomPasswordHasher()
	service := NewCustomerSignupService(repo, hasher)

	t.Run("customer already exists", func(t *testing.T) {

		input := &customer.CreateCustomerIn{
			Name:     "existing_name",
			Email:    "existing_email",
			Password: "existing_password",
		}

		err := service.Execute(input)
		assert.Equal(t, err, customer.ErrCustomerAlreadyExists)
		assert.NotNil(t, err)
	})

	t.Run("customer validation error", func(t *testing.T) {

		input := &customer.CreateCustomerIn{
			Name:     "",
			Email:    "any_email",
			Password: "any_password",
		}

		err := service.Execute(input)
		assert.Equal(t, err, errors.New("invalid customer name"))
		assert.NotNil(t, err)
	})
}
