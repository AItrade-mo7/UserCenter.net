package private

import (
	"DataCenter.net/server/router/api/account"
	"DataCenter.net/server/router/api/coinAI"
	"github.com/gofiber/fiber/v2"
)

/*

/api/private

*/

func Router(router fiber.Router) {
	r := router.Group("/private", MiddleWare)

	r.Get("/get_user_info", account.GetUserInfo)
	r.Post("/edit_profile", account.EditProfile)

	r.Post("/coinAI/List", coinAI.List)
	r.Post("/coinAI/PublicList", coinAI.PublicList)
	r.Post("/coinAI/Remove", coinAI.Remove)

	r.Post("/genshin_check", account.GenshinCheck)
}
