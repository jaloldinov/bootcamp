package handler

import (
	"app/models"
	"app/pkg/helper"
	"app/pkg/logger"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateStaff godoc
// @Router       /staff [POST]
// @Summary      CREATES STAFF
// @Description  CREATES STAFF BASED ON GIVEN DATA
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateStaff  true  "staff data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateStaff(c *gin.Context) {

	var staff models.CreateStaff
	err := c.ShouldBindJSON(&staff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	hashedPass, err := helper.GeneratePasswordHash(staff.Password)
	if err != nil {
		h.log.Error("error while generating hash password:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	staff.Password = string(hashedPass)

	resp, err := h.storage.Staff().CreateStaff(c.Request.Context(), &staff)
	if err != nil {
		fmt.Println("error Staff Create:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "staff successfully created", "id": resp})
}

// GetStaff godoc
// @Router       /staff/{id} [GET]
// @Summary      GET BY ID
// @Description  get staff by ID
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Staff ID" format(uuid)
// @Success      200  {object}  models.Staff
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetStaff(c *gin.Context) {
	response := models.Staff{}
	id := c.Param("id")

	ok, err := h.redis.Cache().Get(c.Request.Context(), id, &response)
	if err != nil {
		fmt.Println("not found staff in redis cache")
	}

	if ok {
		c.JSON(http.StatusOK, response)
		return
	}

	resp, err := h.storage.Staff().GetStaff(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error staff Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})

	err = h.redis.Cache().Create(c.Request.Context(), id, resp, time.Hour)
	if err != nil {
		fmt.Println("error staff Create in redis cache:", err.Error())
	}
}

// ListStaffes godoc
// @Router       /staff [GET]
// @Summary      GET  ALL STAFFS
// @Description  gets all staffs based on limit, page and search by name
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page         query     int        false  "page"          minimum(1)     default(1)
// @Param   search         query     string        false  "search"
// @Success      200  {object}  models.GetAllStaff
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllStaff(c *gin.Context) {
	h.log.Info("request GetALLstaff")
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

	resp, err := h.storage.Staff().GetAllStaff(c.Request.Context(), &models.GetAllStaffRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error  Getallstaff:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to getAllStaff")
	c.JSON(http.StatusOK, resp)
}

// UpdateStaffs godoc
// @Router       /staff/{id} [PUT]
// @Summary      UPDATE STAFF BY ID
// @Description  UPDATES STAFF BASED ON GIVEN DATA AND ID
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staff" format(uuid)
// @Param        data  body      models.CreateStaff  true  "staff data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaff(c *gin.Context) {
	var staff models.Staff
	err := c.ShouldBind(&staff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	staff.ID = c.Param("id")
	resp, err := h.storage.Staff().UpdateStaff(c.Request.Context(), &staff)
	if err != nil {
		h.log.Error("error Staff Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "staff successfully updated", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), staff.ID)
	if err != nil {
		fmt.Println("error staff Delete in redis cache:", err.Error())
	}
}

// DeleteStaff godoc
// @Router       /staff/{id} [DELETE]
// @Summary      DELETE STAFF BY ID
// @Description  DELETES STAFF BASED ON ID
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staff" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteStaff(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Staff().DeleteStaff(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting staff:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "staff successfully deleted", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error staff Delete in redis cache:", err.Error())
	}

}

// func (h *Handler) ChangeBalance(c *gin.Context) {

// }

// CHANGE PASSWORD godoc
// @Router       /staff/change-password/{id} [POST]
// @Summary      UPDATE STAFF PASSWORD BY ID
// @Description  UPDATES STAFF PASSWORD BASED ON GIVEN OLD AND NEW PASSWORD
// @Tags         STAFF
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of staff" format(uuid)
// @Param        data  body      models.ChangePassword  true  "staff data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateStaffPassword(c *gin.Context) {
	var req models.ChangePasswordRequest
	id := c.Param("id")

	err := c.ShouldBind(&req)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	staff, err := h.storage.Staff().GetStaff(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error  get staff:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error: not found staff with id:": id})
		return
	}

	err = helper.ComparePasswords([]byte(staff.Password), []byte(req.OldPassword))
	if err != nil {
		h.log.Error("error: password didn't match:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error:": "password didn't match:"})
		return
	}

	hashedPass, err := helper.GeneratePasswordHash(req.NewPassword)
	if err != nil {
		h.log.Error("error while generating hash password:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	req.Id = id
	req.NewPassword = string(hashedPass)
	resp, err := h.storage.Staff().ChangePassword(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("error change password:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "password updated", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error staff Delete in redis cache:", err.Error())
	}
}
