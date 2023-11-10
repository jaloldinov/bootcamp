package v1

import (
	"fmt"
	"net/http"
	"strconv"

	user_service "api_gateway/genproto/user_service"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Router       /v1/user/create [post]
// @Summary      Create a new user
// @Description  Create a new user with the provuser_ided details
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user     body  user_service.CreateUsersRequest true  "data of the user"
// @Success      201  {object}  user_service.Response
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) CreateUser(ctx *gin.Context) {
	var user = user_service.CreateUsersRequest{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		h.handlerResponse(ctx, "CreateUser", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.services.User().Create(ctx, &user_service.CreateUsersRequest{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		BranchId:  user.BranchId,
		Phone:     user.Phone,
		Login:     user.Login,
		Password:  user.Password,
	})

	if err != nil {
		h.handlerResponse(ctx, "CatgeoryService().Create", http.StatusBadRequest, err.Error())

		return
	}

	h.handlerResponse(ctx, "create user response", http.StatusOK, resp)
}

// GetAllUser godoc
// @Router       /v1/user/list [get]
// @Summary      GetAll User
// @Description  get user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "limit for response"  Default(10)
// @Param		 page     query     int  false  "page for response"   Default(1)
// @Param        firstname     query     string false "firstname"
// @Param        lastname     query     string false "lastname"
// @Param        phone     query     string false "phone"
// @Param        created_at_from     query     string false "search by created_at_from"
// @Param        created_at_to     query     string false "search by created_at_to"
// @Success      200  {array}   user_service.ListUsersResponse
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) GetListUser(ctx *gin.Context) {
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

	resp, err := h.services.User().List(ctx.Request.Context(), &user_service.ListUsersRequest{
		Page:          int32(page),
		Limit:         int32(limit),
		Firstname:     ctx.Query("firstname"),
		Lastname:      ctx.Query("lastname"),
		Phone:         ctx.Query("phone"),
		CreatedAtFrom: ctx.Query("created_at_from"),
		CreatedAtTo:   ctx.Query("created_at_to"),
	})

	if err != nil {
		h.handlerResponse(ctx, "error GetListUser", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get AllUser response", http.StatusOK, resp)
}

// GetUser godoc
// @Router       /v1/user/get/{user_id} [get]
// @Summary      Get a user by ID
// @Description  Retrieve a user by its unique user_identifier
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user_id   path    string     true    "User ID to retrieve"
// @Success      200  {object}  user_service.Users
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) GetUser(ctx *gin.Context) {
	user_id := ctx.Param("user_id")

	resp, err := h.services.User().Get(ctx.Request.Context(), &user_service.IdRequest{Id: user_id})
	if err != nil {
		h.handlerResponse(ctx, "error user GetById", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "get user response", http.StatusOK, resp)
}

// UpdateProduct godoc
// @Router       /v1/user/update/{user_id} [put]
// @Summary      Update an existing user
// @Description  Update an existing user with the provuser_ided details
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user_id       path    int     true    "User ID to update"
// @Param        user   body    user_service.UpdateUsersRequest  true    "Updated data for the user"
// @Success      200  {object}  user_service.UpdateUsersRequest
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) UpdateUser(ctx *gin.Context) {
	var user = user_service.UpdateUsersRequest{}
	user_id, _ := strconv.Atoi(ctx.Param("user_id"))

	err := ctx.ShouldBind(&user)
	if err != nil {
		h.handlerResponse(ctx, "error while binding", http.StatusBadRequest, err.Error())
		return
	}

	user.Id = int32(user_id)

	resp, err := h.services.User().Update(ctx.Request.Context(), &user)
	fmt.Println("before  send bind", resp)
	if err != nil {
		h.handlerResponse(ctx, "error user Update", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "update user response", http.StatusOK, resp)
}

// DeleteUser godoc
// @Router       /v1/user/delete/{user_id} [delete]
// @Summary      Delete a Catgory
// @Description  delete a user by its unique user_identifier
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user_id   path    string     true    "Catgeory ID to retrieve"
// @Success      200  {object}  user_service.Response
// @Failure      400  {object}  Response{data=string}
// @Failure      404  {object}  Response{data=string}
// @Failure      500  {object}  Response{data=string}
func (h *handlerV1) DeleteUser(ctx *gin.Context) {
	user_id := ctx.Param("user_id")

	resp, err := h.services.User().Delete(ctx.Request.Context(), &user_service.IdRequest{Id: user_id})
	if err != nil {
		h.handlerResponse(ctx, "error user Delete", http.StatusBadRequest, err.Error())
		return
	}

	h.handlerResponse(ctx, "delete user response", http.StatusOK, resp)
}
