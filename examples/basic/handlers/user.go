package handlers

import "github.com/gin-gonic/gin"

type UserHandler struct {
	Handler
}

func (h *UserHandler) Index(c *gin.Context) {
	a := h.AppContext(c)

	logger := a.NewLogger("User.Index Handler")
	logger.Info("hede hodo")

	user := a.Services().UserService().GetUserById(123)
	c.JSON(200, user)
}
