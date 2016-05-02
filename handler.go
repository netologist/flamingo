package vodka

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) AppContext(c *gin.Context) *AppContext {
	return NewAppContext(c)
}
