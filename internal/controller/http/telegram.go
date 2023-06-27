package v1

import (
	"MateMind/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"

	"MateMind/internal/usecase"
	"MateMind/pkg/logger"
)

type orderRoutes struct {
	o usecase.User
	l logger.Interface
}

func newOrderRoutes(handler *gin.RouterGroup, order usecase.User, log logger.Interface) {
	r := &orderRoutes{order, log}

	h := handler.Group("/order")
	{
		h.GET("", r.getOrder)
		h.POST("", r.createOrder)
		h.PUT("", r.updateOrder)
		h.DELETE("", r.cancelOrder)
		h.GET("/status", r.getOrderStatus)
		h.GET("/all", r.getAllOrders)
	}
}

type orderResponse struct {
	Order entity.User `json:"order"`
}

func (r *orderRoutes) getOrder(c *gin.Context) {
	order, err := r.o.CreateUser(c.Request.Context())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, orderResponse{order})
}

func (r *orderRoutes) createOrder(c *gin.Context) {

}
func (r *orderRoutes) cancelOrder(c *gin.Context) {

}

func (r *orderRoutes) updateOrder(c *gin.Context) {

}

func (r *orderRoutes) getOrderStatus(c *gin.Context) {

}

func (r *orderRoutes) getAllOrders(c *gin.Context) {

}
