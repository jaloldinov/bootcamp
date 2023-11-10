package v1

import (
	"api_gateway/config"
	"api_gateway/genproto/order_service"
	"api_gateway/genproto/user_service"
	"api_gateway/pkg/helper"
	"api_gateway/pkg/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Router       /v1/courier/active-orders/list [get]
// @Summary      List Order Active orders
// @Description  get Order
// @Tags         courier
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Success 200 {object} models.ResponseModel{data=order_service.ListOrderResponse} "desc"
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) GetListActiveOrders(c *gin.Context) {

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

	resp, err := h.services.Order().GetListActiveOrders(c.Request.Context(), &order_service.ActiveOrderReq{
		Page:  int32(page),
		Limit: int32(limit),
	})

	if err != nil {
		h.log.Error("error Order GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)
}

// @Router       /v1/courier/delete-order/{id} [PUT]
// @Summary      Update Order
// @Description  api for update order
// @Tags         courier
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of order"
// @Success      200  {string}   string
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) DeleteCourierInOrder(c *gin.Context) {
	id := c.Param("id")

	respOrder, err := h.services.Order().Get(c.Request.Context(), &order_service.IdRequest{Id: id})

	if err != nil {
		fmt.Println("error while getting order:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")

		return
	}
	ids, _ := strconv.Atoi(id)
	resp, err := h.services.Order().Update(c.Request.Context(), &order_service.UpdateOrderRequest{
		Id:            int32(ids),
		OrderId:       respOrder.OrderId,
		ClientId:      respOrder.ClientId,
		BranchId:      respOrder.BranchId,
		CourierId:     0,
		Type:          respOrder.Type,
		Address:       respOrder.Address,
		DeliveryPrice: float64(respOrder.DeliveryPrice),
		Price:         float64(respOrder.Price),
		Discount:      float64(respOrder.Discount),
		PaymentType:   respOrder.PaymentType,
		Status:        "accepted",
	})

	if err != nil {
		fmt.Println("error updating order:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}

// @Router       /v1/courier/get-order/{id} [put]
// @Summary      Update Order
// @Description  api for update order
// @Tags         courier
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of order"
// @Success      200  {string}   string
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) CourierGetOrder(c *gin.Context) {

	tokenString := c.Request.Header.Get("Authorization")

	token, err := helper.ExtractToken(tokenString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := helper.ParseClaims(token, config.JWTSecretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respCourier, err := h.services.Courier().Get(c.Request.Context(), &user_service.IdRequest{Id: claims.UserID})

	if err != nil {
		fmt.Println("error Order Get:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	id := c.Param("id")

	respOrder, err := h.services.Order().Get(c.Request.Context(), &order_service.IdRequest{Id: id})

	if err != nil {
		fmt.Println("error Order Get:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	if respOrder.CourierId != 0 {
		c.JSON(http.StatusBadRequest, "order already received")
		return
	}

	idstr := strconv.Itoa(int(respOrder.CourierId))
	respCouriersOrder, err := h.services.Order().GetListByCourierId(c.Request.Context(), &order_service.IdRequest{
		Id: idstr})

	if err != nil {
		fmt.Println("error Order Get:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	if respCouriersOrder.Count == respCourier.MaxOrderCount {
		c.JSON(http.StatusBadRequest, "you are reached max order!")
		return
	}

	resp, err := h.services.Order().UpdateStatus(c.Request.Context(), &order_service.UpdateOrderStatusRequest{
		Id:      respOrder.Id,
		OrderId: respOrder.OrderId,
		Status:  "courier_accepted",
	})

	if err != nil {
		fmt.Println("error Order Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	h.handleSuccessResponse(c, http.StatusOK, "OK", resp)

}
