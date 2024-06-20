package handler

import (
	"net/http"
	"strconv"

	"github.com/ShawaDev/Dealer/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddCar(c *gin.Context) {
	var car model.Car

	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := h.services.AddCar(car); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"id": car.Id,
	})
}

func (h *Handler) GetCarById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	car, err := h.services.GetCarById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(200, car)
}

func (h *Handler) EditCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	var car model.Car
	car.Id = id
	if err := c.BindJSON(&car); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = h.services.EditCar(car, car.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"id": car.Id,
	})
}

func (h *Handler) DeleteCar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := h.services.DeleteCar(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, gin.H{
		"id": id,
	})
}

func (h *Handler) GetAllCars(c *gin.Context) {
	cars, err := h.services.GetAllCars()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cars)
}
