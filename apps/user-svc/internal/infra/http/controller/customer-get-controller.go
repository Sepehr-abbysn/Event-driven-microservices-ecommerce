package controller

import (
	"net/http"
	"strings"
	"user-svc/internal/application"
	"user-svc/internal/domain/customer"
	"user-svc/internal/infra/database"
	"user-svc/internal/infra/http/helper"
)

func CustomerGet(w http.ResponseWriter, r *http.Request) {
	authorization := r.Header["Authorization"]
	if len(authorization) == 0 {
		helper.HandleHttpError(w, customer.ErrCustomerPermissionDenied)
		return
	}

	token := strings.Split(authorization[0], " ")
	if len(token) != 2 || token[0] != "Bearer" || token[1] == "" {
		helper.HandleHttpError(w, customer.ErrCustomerPermissionDenied)
		return
	}

	c := []customer.Customer{
		{
			ID:       "existing_id",
			Name:     "name",
			Email:    "email@email.com",
			Password: "hashed:my_pass",
		},
	}
	repo := database.NewInMemoryCustomerRepo(c)
	customerGetService := application.NewCustomerGetService(repo)

	res, err := customerGetService.Execute(token[1])
	if err != nil {
		helper.HandleHttpError(w, err)
		return
	}

	helper.HandleHttpSuccessJson(w, http.StatusOK, res)
}
