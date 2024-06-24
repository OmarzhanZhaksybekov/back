package handler

import (
	"github.com/ShawaDev/Dealer/pkg/middleware"
	"github.com/ShawaDev/Dealer/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(func(c *gin.Context) {
		// Разрешаем доступ с любого источника
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
	cars := router.Group("/cars")
	{
		cars.GET("/:id", h.GetCarById)
		cars.GET("/", h.GetAllCars)
	}

	admin := router.Group("/admin")
	admin.Use(middleware.RoleMiddleware)
	{
		admin.POST("/add", h.AddCar)
		admin.PUT("/:id", h.EditCar)
		admin.DELETE("/:id", h.DeleteCar)
	}
	return router
}
