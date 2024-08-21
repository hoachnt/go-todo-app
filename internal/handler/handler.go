package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/hoachnt/go-todo-app/internal/service"
	"github.com/hoachnt/go-todo-app/pkg/auth"
)

type Handler struct {
	services *service.Service
	auth     *auth.User
}

func NewHandler(services *service.Service, auth *auth.User) *Handler {
	return &Handler{services: services, auth: auth}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// Swagger setup
	// Authentication routes
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// API routes with authentication middleware
	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group("/:id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
