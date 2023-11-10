package handler

import (
	"fmt"
	"market/api/response"
	"market/config"
	"market/models"
	"market/pkg/helper"
	"market/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// create person handler
// @Router       /login [post]
// @Summary      create person
// @Description  api for create persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        person    body     models.LoginReq  true  "data of person"
// @Success      200  {object}  models.LoginRes
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) Login(c *gin.Context) {
	var req models.LoginReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "BAD REQUEST", Message: "invalid fields in body"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	fmt.Println(req)
	// resp, err := h.redisStrg.Person().GetByUsername(models.RequestByUsername{
	// 	Username: req.Username,
	// })
	if err != nil {
		fmt.Println("error Person GetByUsername:", err.Error())
		res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	// err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(req.Password))
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "INVALID Password", Message: "invalid password"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	m := make(map[string]interface{})
	// m["user_id"] = resp.Id
	// m["branch_id"] = resp.BranchId
	token, err := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)
	c.JSON(http.StatusCreated, models.LoginRes{Token: token})
}
