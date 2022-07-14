package public

import (
	"DataCenter.net/server/router/api/account"
	"github.com/gofiber/fiber/v2"
)

/*
/api/public
*/

func Router(router fiber.Router) {
	r := router.Group("/public", MiddleWare)

	r.Post("/send_code", account.SendEmailCode)
	r.Post("/register", account.Register)
	r.Post("/change_password", account.ChangePassword)

	r.Post("/login", account.Login)
}
