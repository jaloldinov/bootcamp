package v1

import (
	"context"
	"fmt"
	"net/http"
	"time"

	user_service "api_gateway/genproto/user_service"
	"api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Branch godoc
// @ID create-branch
// @Router /v1/branch/create [POST]
// @Summary create branch
// @Description create branch
// @Tags branch
// @Accept json
// @Produce json
// @Param branch body user_service.CreateBranchRequest true "branch"
// @Success 200 {object} models.ResponseModel{data=user_service.Response}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateBranch(c *gin.Context) {
	var branch user_service.CreateBranchRequest

	if err := c.BindJSON(&branch); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Branch().Create(c.Request.Context(), &branch)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating branch", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllBranch godoc
// @ID get-branch
// @Router /v1/branch/list [GET]
// @Summary get list branch
// @Description get branch
// @Tags branch
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param name query integer false "name"
// @Param created_from query string false "created_from"
// @Param created_to query string false "created_to"
// @Success 200 {object} models.ResponseModel{data=user_service.ListBranchResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllBranch(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		fmt.Println(err)
		return
	}

	page, err := h.ParseQueryParam(c, "page", "1")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := h.services.Branch().List(
		c.Request.Context(),
		&user_service.ListBranchRequest{
			Limit:         int32(limit),
			Page:          int32(page),
			Name:          c.Query("name"),
			CreatedAtFrom: c.Query("created_from"),
			CreatedAtTo:   c.Query("created_to"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error get list branch", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-BranchByID godoc
// @ID get-branch-byID
// @Router /v1/branch/get/{branch_id} [GET]
// @Summary get branch by ID
// @Description get branch
// @Tags branch
// @Accept json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Success 200 {object} models.ResponseModel{data=user_service.Branch{}}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetBranch(c *gin.Context) {

	id := user_service.IdRequest{
		Id: c.Param("branch_id"),
	}

	resp, err := h.services.Branch().Get(
		context.Background(),
		&id,
	)

	if !handleError(h.log, c, err, "error get branch by id") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// Update Branch godoc
// @ID update_branch
// @Router /v1/branch/update/{branch_id} [PUT]
// @Summary Update Branch
// @Description Update Branch by ID
// @Tags branch
// @Accept json
// @Produce json
// @Param branch body user_service.UpdateBranchRequest true "branch"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateBranch(c *gin.Context) {

	var branch = user_service.UpdateBranchRequest{}

	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding", err.Error())
		return
	}

	resp, err := h.services.Branch().Update(c.Request.Context(), &branch)
	if !handleError(h.log, c, err, "error while getting branch") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// / Delete Branch godoc
// @ID delete-branch
// @Router /v1/branch/delete/{branch_id} [DELETE]
// @Summary delete branch
// @Description Delete Branch
// @Tags branch
// @Accept json
// @Produce json
// @Param branch_id path string true "branch_id"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteBranch(c *gin.Context) {

	var id user_service.IdRequest

	id.Id = c.Param("branch_id")

	resp, err := h.services.Branch().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while deleting branch", err.Error())
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// @Router       /v1/branch/list/active [get]
// @Summary      List Branch
// @Description  get Branch
// @Tags         branch
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success 200 {object} models.ResponseModel{data=user_service.ListBranchResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetListActiveBranch(c *gin.Context) {

	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		fmt.Println(err)
		return
	}

	page, err := h.ParseQueryParam(c, "page", "1")
	if err != nil {
		fmt.Println(err)
		return
	}

	timeNow := time.Now().Format("15:04:05")

	resp, err := h.services.Branch().ListActive(c.Request.Context(), &user_service.ListBranchActiveRequest{
		Page:  int32(page),
		Limit: int32(limit),
		Time:  timeNow,
	})

	if err != nil {
		h.log.Error("error getting active branchs:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}
