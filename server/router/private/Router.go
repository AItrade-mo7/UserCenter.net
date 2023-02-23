package private

import (
	"UserCenter.net/server/router/api/account"
	"UserCenter.net/server/router/api/coinAI"
	"github.com/gofiber/fiber/v2"
)

func Router(router fiber.Router) {
	r := router.Group("/private", MiddleWare)

	/*
		/api/private
	*/
	r.Post("/loginOut", account.LoginOut)
	r.Get("/get_user_info", account.GetUserInfo)
	r.Post("/edit_profile", account.EditProfile)

	r.Post("/coinAI/List", coinAI.List)
	r.Post("/coinAI/PublicList", coinAI.PublicList)
	r.Post("/coinAI/Remove", coinAI.Remove)
}
