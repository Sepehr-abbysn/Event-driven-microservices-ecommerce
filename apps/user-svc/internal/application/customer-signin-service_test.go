package application

import (
	"testing"
	"user-svc/internal/domain/customer"
	"user-svc/internal/infra/database"
	"user-svc/internal/infra/encryption"

	"github.com/stretchr/testify/assert"
)

func TestCustomerSignin(t *testing.T) {
	c := []customer.Customer{
		{
			ID:       "existing_id",
			Name:     "existing_name",
			Email:    "existing_email",
			Password: "hashed:existing_password",
		},
	}

	repo := database.NewInMemoryCustomerRepo(c)
	hasher := encryption.NewStubPasswordHasher()
	tkGen := encryption.NewStubTokenGenerator()
	service := NewCustomerSigninService(repo, hasher, tkGen)

	t.Run("customer not found", func(t *testing.T) {
		input := &customer.SigninCustomerIn{
			Email:    "no_email",
			Password: "no_password",
		}

		_, err := service.Execute(input)
		assert.Equal(t, err, customer.ErrCustomerInvalidCredential)
		assert.NotNil(t, err)
	})

	t.Run("customer invalid password", func(t *testing.T) {
		input := &customer.SigninCustomerIn{
			Email:    "existing_email",
			Password: "wrong_password",
		}

		_, err := service.Execute(input)
		assert.Equal(t, err, customer.ErrCustomerInvalidCredential)
		assert.NotNil(t, err)
	})

	t.Run("customer signin success", func(t *testing.T) {
		input := &customer.SigninCustomerIn{
			Email:    "existing_email",
			Password: "existing_password",
		}

		res, err := service.Execute(input)
		assert.Equal(t, err, nil)
		assert.Equal(t, &customer.SigninCustomerOut{
			AccessToken: "existing_id",
			Customer: customer.CustomerOut{
				ID:    "existing_id",
				Name:  "existing_name",
				Email: "existing_email",
			},
		}, res)
	})
}
