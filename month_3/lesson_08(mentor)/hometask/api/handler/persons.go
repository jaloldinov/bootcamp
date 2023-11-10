package handler

import (
	"errors"
	"fmt"
	"market/api/response"
	pb "market/genproto"
	"market/models"
	"market/pkg/helper"
	"market/pkg/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// create person handler
// @Router       /person [post]
// @Summary      create person
// @Description  api for create persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        person    body     models.CreatePerson  true  "data of person"
// @Success      200  {object}  response.CreateResponse
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) CreatePerson(c *gin.Context) {
	var person models.CreatePerson
	err := c.ShouldBindBodyWith(&person, binding.JSON)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "BAD REQUEST", Message: "invalid fields in body"}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	passwd, err := helper.GeneratePasswordHash(person.Password)
	person.Password = string(passwd)
	resp, err := h.grpcClient.PersonService().Create(c.Request.Context(), &pb.CreatePerson{
		Name:    person.Name,
		Address: person.Job,
	})
	if err != nil {
		fmt.Println("error Person Create:", err.Error())
		res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	c.JSON(http.StatusCreated, response.CreateResponse{Id: resp.GetId()})
}

// @Router       /person/{id} [put]
// @Summary      update person
// @Description  api for update persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of person"
// @Param        person    body     models.CreatePerson  true  "data of person"
// @Success      200  {string}   string
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) UpdatePerson(c *gin.Context) {
	var person models.Person
	err := c.ShouldBindJSON(&person)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}
	person.Id = c.Param("id")
	resp, err := h.psqlStrg.Person().Update(person)
	if err != nil {
		fmt.Println("error Person Update:", err.Error())
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	err = h.redisStrg.Cache().Delete(c.Request.Context(), person.Id)
	if err != nil {
		fmt.Println("error deleting cache Person Create:", err.Error())
		res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Router       /person/{id} [get]
// @Summary      List persons
// @Description  get persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of person"  Format(uuid)
// @Success      200  {object}   models.Person
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetPerson(c *gin.Context) {
	var resp = &models.Person{}
	id := c.Param("id")
	found, err := h.redisStrg.Cache().Get(c.Request.Context(), id, resp)
	if err != nil {
		fmt.Println("error Person Get:", err.Error())
	}
	if found {
		c.JSON(http.StatusOK, resp)
		return
	}
	resp, err = h.psqlStrg.Person().Get(models.RequestByID{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		fmt.Println("error Person Get:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)

	err = h.redisStrg.Cache().Create(c.Request.Context(), id, resp, 0)
	if err != nil {
		fmt.Println("error deleting cache Person Create:", err.Error())
		res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

}

// @Security ApiKeyAuth
// @Router       /person [get]
// @Summary      List persons
// @Description  get persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        limit    query     integer  true  "limit for response"  Default(10)
// @Param        page    query     integer  true  "page of req"  Default(1)
// @Param        job    query     string  false  "filter by job"  Enums(dev,backend,frontend)
// @Success      200  {array}   models.Person
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetAllPersons(c *gin.Context) {
	// h.log.Info("request GetAllPersons")
	resp := &models.GetAllPersonsResponse{}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		h.log.Error("error get page:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	age, err := strconv.Atoi(c.DefaultQuery("age", "0"))
	if err != nil {
		h.log.Error("error get limit:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid page param")
		return
	}
	// found, err := h.redisStrg.Cache().Get(c.Request.Context(), fmt.Sprintf("persons@%d@%d", page, limit), resp)
	// if err != nil {
	// 	fmt.Println("error Person Get:", err.Error())
	// }
	// if found {
	// 	c.JSON(http.StatusOK, resp)
	// 	return
	// }
	resp, err = h.psqlStrg.Person().GetAll(models.GetAllPersonsRequest{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
		Job:    c.Query("job"),
		Age:    age,
	})
	if err != nil {
		h.log.Error("error Person GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	// err = h.redisStrg.Cache().Create(c.Request.Context(), "persons", resp, 0)
	// if err != nil {
	// 	fmt.Println("error deleting cache Person Create:", err.Error())
	// 	res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
	// 	c.JSON(http.StatusInternalServerError, res)
	// 	return
	// }
	// h.log.Warn("response to GetAllPersons")
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	if !helper.IsValidUUID(id) {
		h.log.Error("error Person GetAll:", logger.Error(errors.New("invalid id")))
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}
	resp, err := h.psqlStrg.Person().Delete(models.RequestByID{ID: id})
	if err != nil {
		h.log.Error("error Person GetAll:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	err = h.redisStrg.Cache().Delete(c.Request.Context(), id)
	if err != nil {
		fmt.Println("error deleting cache Person Create:", err.Error())
		res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// create person handler
// @Router       /person/{id} [post]
// @Summary      create person
// @Description  api for create persons
// @Tags         persons
// @Accept       json
// @Produce      json
// @Param        id    path     string  true  "id of person"
// @Param        person    body     models.CreatePerson  true  "data of person"
// @Success      200  {object}  response.ChangePassword
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) ChangePassword(c *gin.Context) {
	var req models.ChangePassword
	userInfo, _ := c.Get("user_info")
	user := userInfo.(helper.TokenInfo)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		res := response.ErrorResp{Code: "BAD REQUEST", Message: "invalid fields in body"}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	person, err := h.psqlStrg.Person().Get(models.RequestByID{ID: user.UserID})
	err = helper.ComparePasswords([]byte(person.Password), []byte(req.OldPassword))
	passwd, err := helper.GeneratePasswordHash(req.NewPassword)
	fmt.Println(passwd)
	// resp, err := h.storage.Person().Create(req)
	if err != nil {
		fmt.Println("error Person Create:", err.Error())
		res := response.ErrorResp{Code: "INTERNAL ERROR", Message: "internal server error"}
		c.JSON(http.StatusInternalServerError, res)
		return
	}
	c.JSON(http.StatusCreated, response.CreateResponse{Id: "resp"})
}
