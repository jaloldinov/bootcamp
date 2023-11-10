package handler

import (
	"app/api/response"
	"app/config"
	"app/models"
	"app/pkg/helper"
	"app/pkg/logger"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create person handler
// @Router       /login [post]
// @Summary      create staff
// @Description  api for create staffes
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        staff    body     models.LoginRequest  true  "data of staff"
// @Success      200  {object}  models.LoginRespond
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetByUsername(c *gin.Context) {
	var req models.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "BAD REQUEST", Message: "invalid fields in body"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	resp, err := h.storage.Staff().GetByUsername(context.Background(), &models.RequestByUsername{
		Username: req.Username,
	})

	if err != nil {
		fmt.Println("error Person GetByUsername:", err.Error())
		res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	if req.Password != resp.Password {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "INVALID Password", Message: "invalid password"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	m := make(map[string]interface{})
	m["user_id"] = resp.ID
	m["branch_id"] = resp.BranchID
	token, _ := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)
	c.JSON(http.StatusCreated, models.LoginRespond{Token: token})
}

/*
func (h *Handler) Exists(phone string) bool {
	return h.storage.Staff().Exists(models.ExistsReq{Phone: phone})
}

func (h *Handler) Login(login, password string) (models.Staff, error) {

	staff, err := h.storage.Staff().GetByLogin(models.LoginRequest{
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

func (h *Handler) Register(req models.CreateStaff) (models.Staff, error) {

	_, err := h.storage.Staff().CreateStaff(req)
	if err != nil {
		return models.Staff{}, err
	}

	// staff, err := h.

	return models.Staff{}, nil
}


*/
