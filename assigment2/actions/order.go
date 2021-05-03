package actions

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ipan97/hactiv8-assigment2/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	db *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{db: db}
}

func (c *OrderHandler) Create(ctx echo.Context) error {
	var order models.Order
	if err := ctx.Bind(&order); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	err := c.db.Model(&models.Order{}).Create(&order).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, &order)
}

func (c *OrderHandler) Update(ctx echo.Context) error {
	var order models.Order
	if err := ctx.Bind(&order); err != nil {
		return ctx.JSON(http.StatusBadRequest, &models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	err := c.db.Model(&models.Order{}).Update(&order).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, &order)
}

func (c *OrderHandler) GetAll(ctx echo.Context) error {
	var orders []models.Order
	err := c.db.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, &orders)
}

func (c *OrderHandler) Delete(ctx echo.Context) error {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	var order models.Order
	err = c.db.Model(&models.Order{}).First(&order, "order_id=?", id).Error
	if gorm.IsRecordNotFoundError(err) {
		return ctx.JSON(http.StatusBadRequest, &models.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	err = c.db.Model(&models.Order{}).Delete(&models.Order{OrderID: id}).Error
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, &models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, &models.APIResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("success deleted order_id %d", id),
	})
}
