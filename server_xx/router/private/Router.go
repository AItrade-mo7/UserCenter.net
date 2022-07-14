package private

import (
	"DataCenter.net/server/router/api/account"
	"DataCenter.net/server/router/api/coinServe"
	"DataCenter.net/server/router/api/okxKey"
	"github.com/gofiber/fiber/v2"
)

/*

/api/private

*/

func Router(router fiber.Router) {
	r := router.Group("/private", MiddleWare)

	r.Get("/okx_list", okxKey.GetOkxList)
	r.Post("/add_okx_key", okxKey.AddOkxKey)

	r.Get("/server_list", coinServe.GetList)
	r.Post("/add_server", coinServe.AddServer)

	r.Get("/get_user_info", account.GetUserInfo)
	r.Post("/edit_profile", account.EditProfile)

	r.Post("/get_deploy_shell", coinServe.DeployShell)

	r.Post("/del_server", coinServe.Del)
	r.Post("/del_okxkey", okxKey.Del)
}
