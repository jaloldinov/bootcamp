package handler

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTopStaff godoc
// @Router       /get_top_staff [GET]
// @Summary      GET top staffes
// @Description  Top ishchilarni chiqarish: berilgan vaqt oralig'ida type dynamic (cashier, shopAssistant)
// @Tags         BIZNESS
// @Accept       json
// @Produce      json
// @Param        from_date   query      string  true  "from_date"
// @Param        to_date   query      string  true  "to_date"
// @Success      200  {object}  models.Branch
// @Failure      400  {object}  response.ErrorResp
// @Failure      404  {object}  response.ErrorResp
// @Failure      500  {object}  response.ErrorResp
func (h *Handler) GetTopStaff(c *gin.Context) {
	FromDate := c.Query("from_date")
	ToDate := c.Query("to_date")
	fmt.Println(FromDate, ToDate)

	resp, err := h.storage.BiznesLoggic().GetTopStaff(&models.TopStaffRequest{FromDate: FromDate, ToDate: ToDate})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		fmt.Println("error get top staff:", err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
