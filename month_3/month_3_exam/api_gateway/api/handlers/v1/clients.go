package v1

import (
	"fmt"
	"net/http"
	"strconv"

	user_service "api_gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

// CreateClients godoc
// @Router       /v1/client/create [post]
// @Summary      Create a new client
// @Description  Create a new client with the provclient_ided details
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        client     body  user_service.CreateClientsRequest true  "data of the client"
// @Success      201  {object}  user_service.Response
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) CreateClients(ctx *gin.Context) {
	var client = user_service.CreateClientsRequest{}

	err := ctx.ShouldBindJSON(&client)
	if err != nil {
		h.handlerResponse(ctx, "CreateClients", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.Client().Create(ctx, &user_service.CreateClientsRequest{
		Firstname:      client.Firstname,
		Lastname:       client.Lastname,
		Phone:          client.Phone,
		Photo:          client.Photo,
		BirthDate:      client.BirthDate,
		DiscountType:   client.DiscountType,
		DiscountAmount: client.DiscountAmount,
	})

	if err != nil {
		h.handlerResponse(ctx, "ClientsService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create client response", http.StatusOK, resp)
}

// GetAllClients godoc
// @Router       /v1/client/list [get]
// @Summary      GetAll Clients
// @Description  get client
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        firstname     query     string false "firstname"
// @Param        lastname     query     string false "lastname"
// @Param        phone     query     string false "phone"
// @Param        created_at_from     query     string false "search by created_at_from"
// @Param        created_at_to     query     string false "search by created_at_to"
// @Param        last_order_date_from     query     string false "search by last_order_date_from"
// @Param        last_order_date_to    query     string false "search by last_order_date_to"
// @Param        total_orders_count_from    query     int false "search by total_orders_count_from"
// @Param        total_orders_count_to    query     int false "search by total_orders_sum_to"
// @Param        total_orders_sum_from    query     int false "search by total_orders_sum_from"
// @Param        total_orders_sum_to    query     int false "search by total_orders_count_to"
// @Param        discount_type    query     string false "search by discount_type"
// @Param        discount_from    query     string false "search by discount_from"
// @Param        discount_to    query     string false "search by discount_to"
// @Success      200  {array}   user_service.ListClientsResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) GetListClients(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		h.handlerResponse(ctx, "error get page", http.StatusBadRequest, err.Error())
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		h.handlerResponse(ctx, "error get limit", http.StatusBadRequest, err.Error())
		return
	}

	total_orders_count_from, err := h.ParseQueryParam(ctx, "total_orders_count_from", "0")
	if err != nil {
		fmt.Println(err)
		return
	}

	total_orders_count_to, err := h.ParseQueryParam(ctx, "total_orders_count_to", "0")
	if err != nil {
		fmt.Println(err)
		return
	}

	total_orders_sum_from, err := h.ParseQueryParam(ctx, "total_orders_sum_from", "0")
	if err != nil {
		fmt.Println(err)
		return
	}
	total_orders_sum_to, err := h.ParseQueryParam(ctx, "total_orders_sum_to", "0")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := h.services.Client().List(ctx.Request.Context(), &user_service.ListClientsRequest{
		Page:                 int32(page),
		Limit:                int32(limit),
		Firstname:            ctx.Query("firstname"),
		Lastname:             ctx.Query("lastname"),
		Phone:                ctx.Query("phone"),
		CreatedAtFrom:        ctx.Query("created_at_from"),
		CreatedAtTo:          ctx.Query("created_at_to"),
		LastOrderedDateFrom:  ctx.Query("last_order_date_from"),
		LastOrderedDateTo:    ctx.Query("last_order_date_to"),
		TotalOrdersCountFrom: int64(total_orders_count_from),
		TotalOrdersCountTo:   int64(total_orders_count_to),
		TotalOrdersSumFrom:   int64(total_orders_sum_from),
		TotalOrdersSumTo:     int64(total_orders_sum_to),
		DiscountType:         ctx.Query("dicount_type"),
		DiscountAmountFrom:   ctx.Query("discount_amount_from"),
		DiscountAmountTo:     ctx.Query("discount_amount_to"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListClients", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get AllClients response", http.StatusOK, resp)
}

// GetClients godoc
// @Router       /v1/client/get/{client_id} [get]
// @Summary      Get a client by ID
// @Description  Retrieve a client by its unique client_identifier
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        client_id   path    string     true    "Clients ID to retrieve"
// @Success      200  {object}  user_service.Clients
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) GetClients(ctx *gin.Context) {
	client_id := ctx.Param("client_id")

	resp, err := h.services.Client().Get(ctx.Request.Context(), &user_service.IdRequest{Id: client_id})
	if err != nil {
		h.handlerResponse(ctx, "error client GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get client response", http.StatusOK, resp)
}

// UpdateProduct godoc
// @Router       /v1/client/update/{client_id} [put]
// @Summary      Update an existing client
// @Description  Update an existing client with the provclient_ided details
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        client_id       path    int     true    "Clients ID to update"
// @Param        client   body    user_service.UpdateClientsRequest  true    "Updated data for the client"
// @Success      200  {object}  user_service.UpdateClientsRequest
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) UpdateClients(ctx *gin.Context) {
	var client = user_service.UpdateClientsRequest{}
	client_id, _ := strconv.Atoi(ctx.Param("client_id"))

	err := ctx.ShouldBind(&client)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	client.Id = int32(client_id)

	resp, err := h.services.Client().Update(ctx.Request.Context(), &client)
	fmt.Println("before  send bind", resp)
	if err != nil {
		h.handlerResponse(ctx, "error user Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update user response", http.StatusOK, resp)
}

// DeleteClients godoc
// @Router       /v1/client/delete/{client_id} [delete]
// @Summary      Delete a Catgory
// @Description  delete a client by its unique client_identifier
// @Tags         client
// @Accept       json
// @Produce      json
// @Param        client_id   path    string     true    "Catgeory ID to retrieve"
// @Success      200  {object}  user_service.Response
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) DeleteClients(ctx *gin.Context) {
	client_id := ctx.Param("client_id")

	resp, err := h.services.Client().Delete(ctx.Request.Context(), &user_service.IdRequest{Id: client_id})
	if err != nil {
		h.handlerResponse(ctx, "error client Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete client response", http.StatusOK, resp)
}
