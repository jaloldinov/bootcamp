package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	order_service "api_gateway/genproto/order_service"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/logger"

	"github.com/gin-gonic/gin"
)

// UpdateStatus Order godoc
// @ID update_order_status
// @Router /v1/order/update/status/{order_id} [PUT]
// @Summary Update Order STATUS
// @Description Update status by ID and status type
// @Tags order
// @Accept json
// @Produce json
// @Param order body order_service.UpdateOrderStatusRequest true "order"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateOrderStatus(c *gin.Context) {
	var order order_service.UpdateOrderStatusRequest
	err := c.ShouldBindJSON(&order)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.services.Order().UpdateStatus(c.Request.Context(), &order)

	if err != nil {
		fmt.Println("error Order Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	orderid := strconv.Itoa(int(order.Id))
	if order.Status == "finished" {
		res, err := h.services.Order().Get(c.Request.Context(), &order_service.IdRequest{Id: orderid})

		if err != nil {
			fmt.Println("error Order Get:", err.Error())
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		_, err = h.services.Client().UpdateOrder(c.Request.Context(), &user_service.UpdateClientsOrderRequest{
			Id:               res.ClientId,
			TotalOrdersCount: 1,
			TotalOrdersSum:   res.Price,
		})

		if err != nil {
			fmt.Println("error Client Update:", err.Error())
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Order godoc
// @ID create-order
// @Router /v1/order/create [POST]
// @Summary create order
// @Description create order
// @Tags order
// @Accept json
// @Produce json
// @Param order body order_service.CreateOrderRequest true "order"
// @Success 200 {object} models.ResponseModel{data=order_service.Response}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) CreateOrder(c *gin.Context) {
	var order order_service.CreateOrderRequest

	if err := c.BindJSON(&order); err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding json", err)
		return
	}

	resp, err := h.services.Order().Create(c.Request.Context(), &order)
	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error while creating order", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusCreated, "ok", resp)
}

// GetAllOrder godoc
// @ID get-order
// @Router /v1/order/list [GET]
// @Summary get list order
// @Description get order
// @Tags order
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query integer false "limit"
// @Param page query integer false "page"
// @Param page query integer false "page"
// @Param order_id query integer false "order_id"
// @Param client_id query integer false "client_id"
// @Param branch_id query integer false "branch_id"
// @Param delivery_type query string false "delivery_type"
// @Param courier_id query integer false "courier_id"
// @Param price_from query integer false "price_from"
// @Param price_to query integer false "price_to"
// @Param payment_type query string false "payment_type"
// @Success 200 {object} models.ResponseModel{data=order_service.ListOrderResponse} "desc"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetAllOrder(c *gin.Context) {
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

	client_id, err := h.ParseQueryParam(c, "client_id", "0")
	if err != nil {
		fmt.Println(err)
		return
	}
	branch_id, err := h.ParseQueryParam(c, "branch_id", "0")
	if err != nil {
		fmt.Println(err)
		return
	}
	courier_id, err := h.ParseQueryParam(c, "courier_id", "0")
	if err != nil {
		fmt.Println(err)
		return
	}
	price_from, err := h.ParseQueryParam(c, "price_from", "0")
	if err != nil {
		fmt.Println(err)
		return
	}
	price_to, err := h.ParseQueryParam(c, "price_to", "0")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := h.services.Order().List(
		c.Request.Context(),
		&order_service.ListOrderRequest{
			Limit:        int32(limit),
			Page:         int32(page),
			OrderId:      c.Query("order_id"),
			ClientId:     int32(client_id),
			BranchId:     int32(branch_id),
			CourierId:    int32(courier_id),
			DeliveryType: c.Query("delivery_type"),
			PriceFrom:    float64(price_from),
			PriceTo:      float64(price_to),
			PaymentType:  c.Query("payment_type"),
		},
	)

	if err != nil {
		h.handleErrorResponse(c, http.StatusInternalServerError, "error get list order", err)
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// Get-OrderByID godoc
// @ID get-order-byID
// @Router /v1/order/get/{order_id} [GET]
// @Summary get order by ID
// @Description get order
// @Tags order
// @Accept json
// @Produce json
// @Param order_id path string true "order_id"
// @Success 200 {object} models.ResponseModel{data=order_service.Order}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) GetOrder(c *gin.Context) {

	order_id := c.Param("order_id")

	resp, err := h.services.Order().Get(
		context.Background(),
		&order_service.IdRequest{
			Id: order_id,
		},
	)

	if !handleError(h.log, c, err, "error get order by id") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// Update Order godoc
// @ID update_order
// @Router /v1/order/update/{order_id} [PUT]
// @Summary Update Order
// @Description Update Order by ID
// @Tags order
// @Accept json
// @Produce json
// @Param order body order_service.UpdateOrderRequest true "order"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) UpdateOrder(c *gin.Context) {

	var order = order_service.UpdateOrderRequest{}

	err := c.ShouldBindJSON(&order)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while binding", err.Error())
		return
	}

	resp, err := h.services.Order().Update(c.Request.Context(), &order)
	if !handleError(h.log, c, err, "error while getting order") {
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// / Delete Order godoc
// @ID delete-order
// @Router /v1/order/delete/{order_id} [DELETE]
// @Summary delete order
// @Description Delete Order
// @Tags order
// @Accept json
// @Produce json
// @Param order_id path string true "order_id"
// @Success 200 {object} models.ResponseModel{data=models.Status}
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Response 400 {object} models.ResponseModel{error=string} "Bad Request"
// @Failure 500 {object} models.ResponseModel{error=string} "Server Error"
func (h *handlerV1) DeleteOrder(c *gin.Context) {

	var id order_service.IdRequest

	id.Id = c.Param("order_id")

	resp, err := h.services.Order().Delete(c.Request.Context(), &id)
	if err != nil {
		h.handleErrorResponse(c, http.StatusBadRequest, "error while deleting order", err.Error())
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}
