package handlers

import "github.com/hasanozgan/vodka"

//import "github.com/hasanozgan/vodka/examples/basic/context"

type UserHandler struct {
	Handler
}

func (h *UserHandler) Index(c vodka.Context) {

	//a := h.AppContext(c)

	//	logger := a.NewLogger("User.Index Handler")
	//	logger.Info("hede hodo")

	//	user := a.Services().UserService().GetUserById(123)
	//c.JSON(200, "user")
}
