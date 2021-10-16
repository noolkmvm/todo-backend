package handler

import (
	"github.com/gin-gonic/gin"
	"todo/pkg/service"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "todo/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.singUp)
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		list := api.Group("/lists")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllList)
			list.GET("/:id", h.getListById)
			list.PUT("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)

			item := list.Group(":id/items")
			{
				item.POST("/", h.createItem)
				item.GET("/", h.getAllItem)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
