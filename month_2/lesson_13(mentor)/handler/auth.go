package handler

import (
	"errors"
	"lesson_20/models"
)

func (h *handler) Exists(phone string) bool {
	return h.strg.Staff().Exists(models.ExistsReq{Phone: phone})
}

func (h *handler) Login(login, password string) (models.Staff, error) {

	staff, err := h.strg.Staff().GetByLogin(models.LoginRequest{
		Login:    login,
		Password: password,
	})

	if err != nil {
		return models.Staff{}, err
	}

	if staff.Login == login && staff.Password == password {
		return staff, nil
	}
	return models.Staff{}, errors.New("incorrect username or password")
}

func (h *handler) Register(req models.CreateStaff) (models.Staff, error) {

	_, err := h.strg.Staff().CreateStaff(req)
	if err != nil {
		return models.Staff{}, err
	}

	// staff, err := h.

	return models.Staff{}, nil
}
