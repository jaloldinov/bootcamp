package v1

import (
	"fmt"
	"net/http"
	"strconv"

	user_service "api_gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

// CreateCourier godoc
// @Router       /v1/courier/create [post]
// @Summary      Create a new courier
// @Description  Create a new courier with the provcourier_ided details
// @Tags         courier
// @Accept       json
// @Produce      json
// @Param        courier     body  user_service.CreateCouriersRequest true  "data of the courier"
// @Success      201  {object}  user_service.Response
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) CreateCourier(ctx *gin.Context) {
	var courier = user_service.CreateCouriersRequest{}

	err := ctx.ShouldBindJSON(&courier)
	if err != nil {
		h.handlerResponse(ctx, "CreateCourier", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.Courier().Create(ctx, &user_service.CreateCouriersRequest{
		Firstname:     courier.Firstname,
		Lastname:      courier.Lastname,
		BranchId:      courier.BranchId,
		Phone:         courier.Phone,
		Login:         courier.Login,
		Password:      courier.Password,
		MaxOrderCount: courier.MaxOrderCount,
	})

	if err != nil {
		h.handlerResponse(ctx, "CatgeoryService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create courier response", http.StatusOK, resp)
}

// GetAllCourier godoc
// @Router       /v1/courier/list [get]
// @Summary      GetAll Courier
// @Description  get courier
// @Tags         courier
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        firstname     query     string false "firstname"
// @Param        lastname     query     string false "lastname"
// @Param        phone     query     string false "phone"
// @Param        created_at_from     query     string false "search by created_at_from"
// @Param        created_at_to     query     string false "search by created_at_to"
// @Success      200  {array}   user_service.ListCouriersResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) GetListCourier(ctx *gin.Context) {
	limit, err := h.ParseQueryParam(ctx, "limit", "10")
	if err != nil {
		fmt.Println(err)
		return
	}

	page, err := h.ParseQueryParam(ctx, "page", "1")
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := h.services.Courier().List(ctx.Request.Context(), &user_service.ListCouriersRequest{
		Page:          int32(page),
		Limit:         int32(limit),
		Firstname:     ctx.Query("firstname"),
		Lastname:      ctx.Query("lastname"),
		Phone:         ctx.Query("phone"),
		CreatedAtFrom: ctx.Query("created_at_from"),
		CreatedAtTo:   ctx.Query("created_at_to"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListCourier", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get AllCourier response", http.StatusOK, resp)
}

// GetCourier godoc
// @Router       /v1/courier/get/{courier_id} [get]
// @Summary      Get a courier by ID
// @Description  Retrieve a courier by its unique courier_identifier
// @Tags         courier
// @Accept       json
// @Produce      json
// @Param        courier_id   path    string     true    "Courier ID to retrieve"
// @Success      200  {object}  user_service.Couriers
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) GetCourier(ctx *gin.Context) {
	courier_id := ctx.Param("courier_id")

	resp, err := h.services.Courier().Get(ctx.Request.Context(), &user_service.IdRequest{Id: courier_id})
	if err != nil {
		h.handlerResponse(ctx, "error courier GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get courier response", http.StatusOK, resp)
}

// UpdateProduct godoc
// @Router       /v1/courier/update/{courier_id} [put]
// @Summary      Update an existing courier
// @Description  Update an existing courier with the provcourier_ided details
// @Tags         courier
// @Accept       json
// @Produce      json
// @Param        courier_id       path    int     true    "Courier ID to update"
// @Param        courier   body    user_service.UpdateCouriersRequest  true    "Updated data for the courier"
// @Success      200  {object}  user_service.UpdateCouriersRequest
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) UpdateCourier(ctx *gin.Context) {
	var courier = user_service.UpdateCouriersRequest{}
	courier_id, _ := strconv.Atoi(ctx.Param("courier_id"))

	err := ctx.ShouldBind(&courier)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	courier.Id = int32(courier_id)

	resp, err := h.services.Courier().Update(ctx.Request.Context(), &courier)
	fmt.Println("before  send bind", resp)
	if err != nil {
		h.handlerResponse(ctx, "error courier Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update courier response", http.StatusOK, resp)
}

// DeleteCourier godoc
// @Router       /v1/courier/delete/{courier_id} [delete]
// @Summary      Delete a Catgory
// @Description  delete a courier by its unique courier_identifier
// @Tags         courier
// @Accept       json
// @Produce      json
// @Param        courier_id   path    string     true    "Catgeory ID to retrieve"
// @Success      200  {object}  user_service.Response
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) DeleteCourier(ctx *gin.Context) {
	courier_id := ctx.Param("courier_id")

	resp, err := h.services.Courier().Delete(ctx.Request.Context(), &user_service.IdRequest{Id: courier_id})
	if err != nil {
		h.handlerResponse(ctx, "error courier Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete courier response", http.StatusOK, resp)
}
