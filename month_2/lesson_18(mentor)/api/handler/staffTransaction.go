package handler

import (
	"app/models"
	"app/pkg/logger"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	var transaction models.CreateTransaction
	err := c.ShouldBind(&transaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, "invalid body")
		return
	}

	resp, err := h.storage.Transaction().CreateTransaction(&transaction)
	if err != nil {
		fmt.Println("error from storage create transaction:", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created", "id": resp})
}

func (h *Handler) GetTransaction(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Transaction().GetTransaction(&models.IdRequest{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println("error from storage get transaction:", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": resp})
}

func (h *Handler) GetAllTransaction(c *gin.Context) {
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

	resp, err := h.storage.Transaction().GetAllTransaction(&models.GetAllTransactionRequest{
		Page:  page,
		Limit: limit,
		Text:  c.Query("search"),
	})
	if err != nil {
		h.log.Error("error from storage getAll transaction:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	var transaction models.Transaction
	err := c.ShouldBind(&transaction)
	if err != nil {
		h.log.Error("error while binding:", logger.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	transaction.Id = c.Param("id")
	resp, err := h.storage.Transaction().UpdateTransaction(&transaction)
	if err != nil {
		h.log.Error("error transaction update:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated", "id": resp})
}

func (h *Handler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	resp, err := h.storage.Transaction().DeleteTransaction(&models.IdRequest{Id: id})
	if err != nil {
		h.log.Error("error deleting transaction:", logger.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transaction successfully deleted", "id": resp})
}

// func (h *Handler) GetTopStaffs(c *gin.Context) {

// 	resp, err := h.strg.Transaction().GetTopStaffs(models.TopWorkerRequest{
// 		Type:     Type,
// 		FromDate: fromData,
// 		ToDate:   ToData,
// 	})

// 	branchNamesMap, _ := h.strg.Branch().GetAllBranch(models.GetAllBranchRequest{})

// 	branchName := make(map[string]string)
// 	for _, b := range branchNamesMap.Branches {
// 		branchName[b.Id] = b.Name
// 	}
// 	for _, v := range resp {
// 		fmt.Printf("Branch: %s Staff: %s Earning: %d\n", branchName[v.BranchId], v.Name, v.Money)
// 	}

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
