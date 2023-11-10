package handler

import (
	"app/models"
	"app/pkg/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStaffTarif(c *gin.Context) {
	var tariff models.CreateStaffTarif
	err := c.ShouldBind(&tariff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Tariff().CreateStaffTarif(&tariff)
	if err != nil {
		fmt.Println("error Tariff Create:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Tariff successfully created", "id": resp})
}

func (h *Handler) GetStaffTarif(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Tariff().GetStaffTarif(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error tariff Get:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})
}

func (h *Handler) GetAllStaffTarif(c *gin.Context) {
	h.log.Info("request GetAllTariff")
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

	resp, err := h.storage.Tariff().GetAllStaffTarif(&models.GetAllStaffTarifRequest{
		Page:  page,
		Limit: limit,
		Name:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error Branch GetAllTariff:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	h.log.Warn("response to GetAllTariff")
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateStaffTarif(c *gin.Context) {
	var tariff models.StaffTarif
	err := c.ShouldBind(&tariff)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tariff.Id = c.Param("id")
	resp, err := h.storage.Tariff().UpdateStaffTarif(&tariff)
	if err != nil {
		h.log.Error("error Tariff Update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tariff successfully updated", "id": resp})
}

func (h *Handler) DeleteStaffTarif(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Tariff().DeleteStaffTarif(&models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting tarif:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "tariff successfully deleted", "id": resp})
}
