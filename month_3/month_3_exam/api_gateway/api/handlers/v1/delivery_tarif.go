package v1

import (
	"context"
	"fmt"
	"net/http"

	order_service "api_gateway/genproto/order_service"

	"github.com/gin-gonic/gin"
)

// DeliveryTariff godoc
// @ID create-delivery
// @Router /v1/delivery/create [POST]
// @Summary create delivery
// @Description create delivery
// @Tags delivery
// @Accept json
// @Produce json
// @Param delivery body order_service.CreateDeliveryTariffRequest true "delivery"
// @Success 200 {object} models.ResponseModel{data=order_service.Response}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateDeliveryTariff(c *gin.Context) {
	var delivery order_service.CreateDeliveryTariffRequest

	if err := c.BindJSON(&delivery); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.DeliveryTariff().Create(c.Request.Context(), &delivery)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating delivery", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllDeliveryTariff godoc
// @ID get-delivery
// @Router /v1/delivery/list [GET]
// @Summary get list delivery
// @Description get delivery
// @Tags delivery
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param type query string false "type"
// @Success 200 {object} models.ResponseModel{data=order_service.ListDeliveryTariffResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllDeliveryTariff(c *gin.Context) {
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

	resp, err := h.services.DeliveryTariff().List(
		c.Request.Context(),
		&order_service.ListDeliveryTariffRequest{
			Limit:     int32(limit),
			Page:      int32(page),
			Search:    c.Query("search"),
			TarifType: c.Query("type"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error get list delivery", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-DeliveryTariffByID godoc
// @ID get-delivery-byID
// @Router /v1/delivery/get/{delivery_id} [GET]
// @Summary get delivery by ID
// @Description get delivery
// @Tags delivery
// @Accept json
// @Produce json
// @Param delivery_id path string true "delivery_id"
// @Success 200 {object} models.ResponseModel{data=order_service.DeliveryTariff}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetDeliveryTariff(c *gin.Context) {

	delivery_id := c.Param("delivery_id")

	resp, err := h.services.DeliveryTariff().Get(
		context.Background(),
		&order_service.IdRequest{
			Id: delivery_id,
		},
	)

	if !handleError(h.log, c, err, "error get delivery by id") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// Update DeliveryTariff godoc
// @ID update_delivery
// @Router /v1/delivery/update/{delivery_id} [PUT]
// @Summary Update DeliveryTariff
// @Description Update DeliveryTariff by ID
// @Tags delivery
// @Accept json
// @Produce json
// @Param delivery body order_service.UpdateDeliveryTariffRequest true "delivery"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateDeliveryTariff(c *gin.Context) {

	var delivery = order_service.UpdateDeliveryTariffRequest{}

	err := c.ShouldBindJSON(&delivery)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding", err.Error())
		return
	}

	resp, err := h.services.DeliveryTariff().Update(c.Request.Context(), &delivery)
	if !handleError(h.log, c, err, "error while getting delivery") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Delete DeliveryTariff godoc
// @ID delete-delivery
// @Router /v1/delivery/delete/{delivery_id} [DELETE]
// @Summary delete delivery
// @Description Delete DeliveryTariff
// @Tags delivery
// @Accept json
// @Produce json
// @Param delivery_id path string true "delivery_id"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteDeliveryTariff(c *gin.Context) {

	var id order_service.IdRequest

	id.Id = c.Param("delivery_id")

	resp, err := h.services.DeliveryTariff().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while deleting delivery", err.Error())
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}
