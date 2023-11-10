package handler

import (
	"app/models"
	"app/pkg/logger"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// CreateBranch godoc
// @Router       /branch [POST]
// @Summary      CREATES BRANCH
// @Description  CREATES BRANCH BASED ON GIVEN DATA
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        data  body      models.CreateBranch  true  "branch data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreateBranch(c *gin.Context) {
	var branch models.CreateBranch

	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Branch().CreateBranch(c.Request.Context(), &branch)
	if err != nil {
		fmt.Println("error Branch Create:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Branch successfully created", "id": resp})
}

// GetBranch godoc
// @Router       /branch/{id} [GET]
// @Summary      GET BY ID
// @Description  get branch by ID
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Branch ID" format(uuid)
// @Success      200  {object}  models.Branch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetBranch(c *gin.Context) {
	response := models.Branch{}
	id := c.Param("id")

	ok, err := h.redis.Cache().Get(c.Request.Context(), id, &response)
	if err != nil {
		fmt.Println("not found branch in redis: ", err)
	}
	if ok {
		c.JSON(http.StatusOK, response)
		return
	}

	resp, err := h.storage.Branch().GetBranch(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error Branch Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

	err = h.redis.Cache().Create(c.Request.Context(), id, &resp, time.Hour)
	if err != nil {
		fmt.Println("error creating branch in redis: ", err)
	}
}

// ListBranchs godoc
// @Router       /branch [GET]
// @Summary      GET  ALL BRANCHS
// @Description  gets all branch based on limit, page and search by name
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param   limit         query     int        false  "limit"          minimum(1)     default(10)
// @Param   page         query     int        false  "page"          minimum(1)     default(1)
// @Param   search         query     string        false  "search"
// @Success      200  {object}  models.GetAllBranch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllBranch(c *gin.Context) {
	h.log.Info("request GetAllBranch")
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

	resp, err := h.storage.Branch().GetAllBranch(c.Request.Context(), &models.GetAllBranchRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Branch GetAllBranch:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllBranch")
	c.JSON(http.StatusOK, resp)
}

// UpdateBranch godoc
// @Router       /branch/{id} [PUT]
// @Summary      UPDATE BRANCH BY ID
// @Description  UPDATES BRANCH BASED ON GIVEN DATA AND ID
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of branch" format(uuid)
// @Param        data  body      models.CreateBranch  true  "branch data"
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdateBranch(c *gin.Context) {
	var branch models.Branch
	err := c.ShouldBind(&branch)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	branch.ID = c.Param("id")
	resp, err := h.storage.Branch().UpdateBranch(c.Request.Context(), &branch)
	if err != nil {
		h.log.Error("error Branch Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update branch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch successfully updated", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), branch.ID)
	if err != nil {
		fmt.Println("error deleting branch in redis: ", err)
	}
}

// DeleteBranch godoc
// @Router       /branch/{id} [DELETE]
// @Summary      DELETE BRANCH BY ID
// @Description  DELETES BRANCH BASED ON ID
// @Tags         BRANCH
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of branch" format(uuid)
// @Success      200  {string}  string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Branch().DeleteBranch(c.Request.Context(), &models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting branch:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Branch successfully deleted", "id": resp})

	err = h.redis.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error deleting branch in redis: ", err)
	}
}
