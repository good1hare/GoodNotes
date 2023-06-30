package v1

import (
	"MateMind/internal/entity"
	"github.com/gin-gonic/gin"
	"net/http"

	"MateMind/internal/usecase"
	"MateMind/pkg/logger"
)

type telegramRoutes struct {
	user usecase.User
	l    logger.Interface
}

func newTelegramRoutes(handler *gin.RouterGroup, order usecase.User, log logger.Interface) {
	r := &telegramRoutes{order, log}

	h := handler.Group("/callback")
	{
		h.POST("", r.createUser)
	}
}

type orderResponse struct {
	User entity.User `json:"user"`
}

func (r *telegramRoutes) getUser(c *gin.Context) {
	user, err := r.user.CreateUser(c.Request.Context())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, orderResponse{user})
}

func (r *telegramRoutes) createUser(c *gin.Context) {

}
func (r *telegramRoutes) cancelOrder(c *gin.Context) {

}

func (r *telegramRoutes) updateOrder(c *gin.Context) {

}

func (r *telegramRoutes) getOrderStatus(c *gin.Context) {

}

func (r *telegramRoutes) getAllOrders(c *gin.Context) {

}
