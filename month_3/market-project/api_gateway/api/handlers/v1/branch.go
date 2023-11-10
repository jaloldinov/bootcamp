package v1

import (
	"context"
	"errors"
	"net/http"

	branch_service "api_gateway/genproto"
	"api_gateway/pkg/util"

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
// @Param branch body branch_service.CreateBranchRequest true "branch"
// @Success 200 {object} models.ResponseModel{data=branch_service.Branch} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateBranch(c *gin.Context) {
	var branch branch_service.CreateBranchRequest

	if err := c.BindJSON(&branch); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.BranchService().Create(c.Request.Context(), &branch)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating branch", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllBranch godoc
// @ID get-branch
// @Router /v1/branch/list [GET]
// @Summary get branch all
// @Description get branch
// @Tags branch
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param offset query integer false "offset"
// @Success 200 {object} models.ResponseModel{data=branch_service.ListBranchResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllBranch(c *gin.Context) {
	limit, err := h.ParseQueryParam(c, "limit", "10")
	if err != nil {
		return
	}

	offset, err := h.ParseQueryParam(c, "offset", "0")
	if err != nil {
		return
	}

	resp, err := h.services.BranchService().List(
		c.Request.Context(),
		&branch_service.ListBranchRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error getting all branchs", err)
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
// @Success 200 {object} models.ResponseModel{data=branch_service.GetBranchResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetBranch(c *gin.Context) {

	branch_id := c.Param("branch_id")

	if !util.IsValidUUID(branch_id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "branch id is not valid", errors.New("branch id is not valid"))
		return
	}

	resp, err := h.services.BranchService().Get(
		context.Background(),
		&branch_service.IdRequest{
			Id: branch_id,
		},
	)
	if !handleError(h.log, c, err, "error while getting branch") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)

}

// Update Branch godoc
// @ID update_branch
// @Router /v1/branch/update/{branch_id} [PUT]
// @Summary Update Branch
// @Description Update Branch by ID
// @Tags branch
// @Accept json
// @Produce json
// @Param        branch_id       path    string     true    "Branch ID to update"
// @Param branch body branch_service.CreateBranchRequest true "branch"
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateBranch(c *gin.Context) {

	var branch = branch_service.UpdateBranchRequest{}

	branch.Id = c.Param("branch_id")
	err := c.ShouldBindJSON(&branch)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding", err.Error())
		return
	}

	resp, err := h.services.BranchService().Update(c.Request.Context(), &branch)
	if !handleError(h.log, c, err, "error while getting branch") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
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
// @Success 200 {object} models.ResponseModel{data=models.Status} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteBranch(c *gin.Context) {

	var id branch_service.IdRequest
	id.Id = c.Param("branch_id")

	if !util.IsValidUUID(id.Id) {
		h.handleErrorResponse(c, http.StatusBadRequest, "branch id is not valid", errors.New("branch id is not valid"))
		return
	}

	resp, err := h.services.BranchService().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while deleting branch", err.Error())
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "ok", resp)
}
