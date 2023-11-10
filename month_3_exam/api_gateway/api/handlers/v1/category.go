package v1

import (
	"context"
	"fmt"
	"net/http"

	catalog_service "api_gateway/genproto/catalog_service"

	"github.com/gin-gonic/gin"
)

// Category godoc
// @ID create-category
// @Router /v1/category/create [POST]
// @Summary create category
// @Description create category
// @Tags category
// @Accept json
// @Produce json
// @Param category body catalog_service.CreateCategoryRequest true "category"
// @Success 200 {object} models.ResponseModel{data=catalog_service.Response}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateCategory(c *gin.Context) {
	var category catalog_service.CreateCategoryRequest

	if err := c.BindJSON(&category); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Category().Create(c.Request.Context(), &category)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating category", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllCategory godoc
// @ID get-category
// @Router /v1/category/list [GET]
// @Summary get list category
// @Description get category
// @Tags category
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Success 200 {object} models.ResponseModel{data=catalog_service.ListCategoryResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllCategory(c *gin.Context) {
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

	resp, err := h.services.Category().List(
		c.Request.Context(),
		&catalog_service.ListCategoryRequest{
			Limit:  int32(limit),
			Page:   int32(page),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error get list category", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-CategoryByID godoc
// @ID get-category-byID
// @Router /v1/category/get/{category_id} [GET]
// @Summary get category by ID
// @Description get category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 200 {object} models.ResponseModel{data=catalog_service.Category}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetCategory(c *gin.Context) {

	category_id := c.Param("category_id")

	resp, err := h.services.Category().Get(
		context.Background(),
		&catalog_service.IdRequest{
			Id: category_id,
		},
	)

	if !handleError(h.log, c, err, "error get category by id") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// Update Category godoc
// @ID update_category
// @Router /v1/category/update/{category_id} [PUT]
// @Summary Update Category
// @Description Update Category by ID
// @Tags category
// @Accept json
// @Produce json
// @Param category body catalog_service.UpdateCategoryRequest true "category"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateCategory(c *gin.Context) {

	var category = catalog_service.UpdateCategoryRequest{}

	err := c.ShouldBindJSON(&category)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding", err.Error())
		return
	}

	resp, err := h.services.Category().Update(c.Request.Context(), &category)
	if !handleError(h.log, c, err, "error while getting category") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// / Delete Category godoc
// @ID delete-category
// @Router /v1/category/delete/{category_id} [DELETE]
// @Summary delete category
// @Description Delete Category
// @Tags category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteCategory(c *gin.Context) {

	var id catalog_service.IdRequest

	id.Id = c.Param("category_id")

	resp, err := h.services.Category().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while deleting category", err.Error())
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}
