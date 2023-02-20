package public

import (
	"UserCenter.net/server/router/api"
	"UserCenter.net/server/router/api/account"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	r := router.Group("/public", MiddleWare)

	/*
		/api/public
	*/

	r.Post("/send_code", account.SendEmailCode)
	r.Post("/register", account.Register)
	r.Post("/change_password", account.ChangePassword)
	r.Post("/login", account.Login)

	r.Get("/InstallCoinAI.sh", api.InstallCoinAI)
}
