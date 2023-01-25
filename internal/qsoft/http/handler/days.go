package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qsoft_test/internal/qsoft/models"
)

func (h *Handler) GetDays(c *gin.Context) {
	year, ok := c.Params.Get("year")
	if !ok {
		c.JSON(http.StatusBadRequest, models.Response{
			Error:     true,
			ErrorText: models.ErrBadParam.Error(),
			Code:      http.StatusBadRequest,
		})
	}

	days, err := h.service.Days(c.Request.Context(), year)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error:     true,
			ErrorText: err.Error(),
			Code:      http.StatusBadRequest,
		})
	}

	c.JSON(http.StatusOK, models.Response{Data: days})
}
