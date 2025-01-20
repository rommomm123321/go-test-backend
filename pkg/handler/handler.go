package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rommomm123321/go-test-backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		sportsHall := api.Group("/sports-hall")
		{
			sportsHall.POST("/", h.createSportsHall)
			sportsHall.GET("/", h.getSportsHall)
			sportsHall.GET("/:id", h.getSportsHallById)
			sportsHall.PUT("/:id", h.updateSportsHall)
			sportsHall.DELETE("/:id", h.deleteSportsHall)
		}

	}
	return router
}
