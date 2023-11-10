package handler

import (
	"auth/api/response"
	"auth/config"
	"auth/models"
	"auth/pkg/helper"
	"auth/pkg/logger"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Router       /user [POST]
// @Summary      CREATES USER
// @Description  CREATES USER BASED ON GIVEN DATA
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUser  true  "user data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateUser(c *gin.Context) {
	var user models.CreateUser
	err := c.ShouldBind(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.User().CreateUser(c.Request.Context(), &user)
	if err != nil {
		fmt.Println("error User Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "login or phone number is already used, enter another one")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created", "id": resp})
}

// GetUser godoc
// @Router       /user/{id} [GET]
// @Summary      GET BY ID
// @Description  get user by ID
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID" format(uuid)
// @Success      200  {object}  models.User
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.User().GetUser(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error User Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ListUsers godoc
// @Router       /user [GET]
// @Summary      GET  ALL USERS
// @Description  gets all user based on limit, page and search by name
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page         query     int        false  "page"          minimum(1)     default(1)
// @Param   search         query     string        false  "search"
// @Success      200  {object}  models.GetAllUser
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllUser(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}

	resp, err := h.storage.User().GetAllUser(c.Request.Context(), &models.GetAllUserRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error User GetAllUser:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateUser godoc
// @Router       /user/{id} [PUT]
// @Summary      UPDATE USER BY ID
// @Description  UPDATES USER BASED ON GIVEN DATA AND ID
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of user" format(uuid)
// @Param        data  body      models.CreateUser  true  "user data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user.ID = c.Param("id")
	resp, err := h.storage.User().UpdateUser(c.Request.Context(), &user)
	if err != nil {
		h.log.Error("error User Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "updated user id": resp})
}

// DeleteUser godoc
// @Router       /user/{id} [DELETE]
// @Summary      DELETE USER BY ID
// @Description  DELETES USER BASED ON ID
// @Tags         USER
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of user" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.User().DeleteUser(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting user:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "deleted user id": resp})
}

// SignUp godoc
// @Router       /signup [POST]
// @Summary      CREATES USER
// @Description  CREATES USER BASED ON GIVEN DATA
// @Tags         AUTH
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateUser  true  "user data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) SignUp(c *gin.Context) {
	var user models.CreateUser
	err := c.ShouldBind(&user)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	if !helper.IsValidPhone(user.PhoneNumber) {
		c.JSON(http.StatusBadRequest, "invalid phone number")
		return
	}

	_, err = h.storage.User().GetByLogin(context.Background(), &models.LoginRequest{
		Login:    user.Login,
		Password: user.Password,
	})

	if err != nil {
		resp, err := h.storage.User().CreateUser(c.Request.Context(), &user)
		if err != nil {
			fmt.Println("error User Create:", err.Error())
			c.JSON(http.StatusInternalServerError, "login or phone number is already used, enter another one")
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "created", "id": resp})
		return
	}
	c.JSON(http.StatusBadGateway, gin.H{"error": "username is used, enter another one"})
}

// loginUser godoc
// @Router       /login [POST]
// @Summary      auth
// @Description  login
// @Tags         AUTH
// @Accept       json
// @Produce      json
// @Param        user    body   models.LoginRequest  true  "data of user"
// @Success      200  {object}  models.LoginRespond
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) Login(c *gin.Context) {
	var req models.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "BAD REQUEST", Message: "invalid fields in body"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	resp, err := h.storage.User().GetByLogin(context.Background(), &models.LoginRequest{
		Login:    req.Login,
		Password: req.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"not found user with login:": req.Login})
		return
	}

	if req.Password != resp.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "login or password didn't match"})
		return
	}

	m := make(map[string]interface{})
	m["login"] = resp.Login
	m["password"] = resp.Password
	m["phone_number"] = resp.PhoneNumber

	token, _ := helper.GenerateJWT(m, config.TokenExpireTime, config.JWTSecretKey)
	c.JSON(http.StatusCreated, models.LoginRespond{Token: token})
}
