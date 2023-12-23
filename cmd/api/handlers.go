package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"

	"github.com/danilomartinelli/users/internal/response"
)

func (app *application) status(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Status": "OK",
	}

	err := response.JSON(w, http.StatusOK, data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	DebtId string `json:"debtId"`
}

func (app *application) getUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")

	if userId == "" {
		app.badRequest(w, r, fmt.Errorf("userId is required"))
		return
	}

	debtId := uuid.New().String()

	user := User{
		Id:     userId,
		Name:   "Danilo Martinelli",
		DebtId: debtId,
	}

	err := response.JSON(w, http.StatusOK, user)
	if err != nil {
		app.serverError(w, r, err)
	}
}
