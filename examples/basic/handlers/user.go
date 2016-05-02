package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Handler
}

func (h *UserHandler) Index(c *gin.Context) {
	a := h.AppContext(c)

	user := a.Services().UserService().GetUserById(123)
	//logger := a.NewLogger("User.Index Handler")
	//logger.Info()
	fmt.Println("user %s", a)
	//log.Infof("username: %s", "ekin")
	//content := gin.H{"Hello": "World"}
	c.JSON(200, user)
}
