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

func newTelegramRoutes(handler *gin.RouterGroup, user usecase.User, log logger.Interface) {
	tr := &telegramRoutes{user, log}

	h := handler.Group("/callback")
	{
		h.POST("", tr.handler)
	}
}

type orderResponse struct {
	User entity.User `json:"user"`
}

func (tr *telegramRoutes) handler(c *gin.Context) {
	user, err := tr.user.CreateUser(c.Request.Context())
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, orderResponse{user})
}
