package v1

import (
	"context"
	"fmt"
	"net/http"

	catalog_service "api_gateway/genproto/catalog_service"

	"github.com/gin-gonic/gin"
)

// Product godoc
// @ID create-product
// @Router /v1/product/create [POST]
// @Summary create product
// @Description create product
// @Tags product
// @Accept json
// @Produce json
// @Param product body catalog_service.CreateProductRequest true "product"
// @Success 200 {object} models.ResponseModel{data=catalog_service.Response}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateProduct(c *gin.Context) {
	var product catalog_service.CreateProductRequest

	if err := c.BindJSON(&product); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Product().Create(c.Request.Context(), &product)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating product", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllProduct godoc
// @ID get-product
// @Router /v1/product/list [GET]
// @Summary get list product
// @Description get product
// @Tags product
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param type query integer false "type"
// @Param category_id query integer false "category_id"
// @Success 200 {object} models.ResponseModel{data=catalog_service.ListProductResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllProduct(c *gin.Context) {
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

	categoryId, err := h.ParseQueryParam(c, "category_id", "0")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := h.services.Product().List(
		c.Request.Context(),
		&catalog_service.ListProductRequest{
			Limit:    int32(limit),
			Page:     int32(page),
			Search:   c.Query("search"),
			Type:     c.Query("type"),
			Category: int32(categoryId),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error get list product", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-ProductByID godoc
// @ID get-product-byID
// @Router /v1/product/get/{product_id} [GET]
// @Summary get product by ID
// @Description get product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 200 {object} models.ResponseModel{data=catalog_service.Product}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetProduct(c *gin.Context) {

	product_id := c.Param("product_id")

	resp, err := h.services.Product().Get(
		context.Background(),
		&catalog_service.IdRequest{
			Id: product_id,
		},
	)

	if !handleError(h.log, c, err, "error get product by id") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// Update Product godoc
// @ID update_product
// @Router /v1/product/update/{product_id} [PUT]
// @Summary Update Product
// @Description Update Product by ID
// @Tags product
// @Accept json
// @Produce json
// @Param product body catalog_service.UpdateProductRequest true "product"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateProduct(c *gin.Context) {

	var product = catalog_service.UpdateProductRequest{}

	err := c.ShouldBindJSON(&product)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding", err.Error())
		return
	}

	resp, err := h.services.Product().Update(c.Request.Context(), &product)
	if !handleError(h.log, c, err, "error while getting product") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// / Delete Product godoc
// @ID delete-product
// @Router /v1/product/delete/{product_id} [DELETE]
// @Summary delete product
// @Description Delete Product
// @Tags product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteProduct(c *gin.Context) {

	var id catalog_service.IdRequest

	id.Id = c.Param("product_id")

	resp, err := h.services.Product().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while deleting product", err.Error())
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}
