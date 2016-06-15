package handlers

import (
	"github.com/hasanozgan/flamingo/examples/basic/context"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) AppContext(c *gin.Context) context.AppContext {
	return context.NewAppContext(c).(context.AppContext)
}
