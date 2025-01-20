package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	app "github.com/rommomm123321/go-test-backend"
)

func (h *Handler) createSportsHall(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	var input app.SportsHall

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	data, err := h.services.SportsHall.Create(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":          data.Id,
		"name":        data.Name,
		"description": data.Description,
		"owner_id":    data.OwnerId,
		"image_url":   data.Image_url,
	})
}

type getAllResponse struct {
	Data []app.SportsHall `json:"data"`
}

func (h *Handler) getSportsHall(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	data, err := h.services.SportsHall.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllResponse{Data: data})
}

func (h *Handler) getSportsHallById(c *gin.Context) {

}

func (h *Handler) updateSportsHall(c *gin.Context) {

}

func (h *Handler) deleteSportsHall(c *gin.Context) {

}
