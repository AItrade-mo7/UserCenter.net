package router

import (
	"os"
	"time"

	"UserCenter.net/server/global"
	"UserCenter.net/server/global/config"
	"UserCenter.net/server/router/api"
	"UserCenter.net/server/router/middle"
	"UserCenter.net/server/router/private"
	"UserCenter.net/server/router/public"
	"UserCenter.net/server/router/wss"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Start() {
	// 加载日志文件
	fileName := config.Dir.Log + "/HTTP-T" + time.Now().Format("06年1月02日15时") + ".log"
	logFile, _ := os.Create(fileName)
	/*
		加载模板
		https://www.gouguoyin.cn/posts/10103.html
	*/

	// 创建服务
	app := fiber.New(fiber.Config{
		ServerHeader: config.SysName,
	})

	// 日志中间件
	app.Use(
		limiter.New(limiter.Config{
			Max:        200,
			Expiration: 1 * time.Second,
		}), // 限流
		logger.New(logger.Config{
			Format:     "[${time}] [${ip}:${port}] ${status} - ${method} ${latency} ${path} \n",
			TimeFormat: "2006-01-02 - 15:04:05",
			Output:     logFile,
		}), // 日志
		cors.New(),     // 允许跨域
		compress.New(), // 压缩
		middle.Public,  // 授权验证
	)

	// AItrade_net
	app.All("/CoinAI/*", api.AIServeProxy)
	app.All("/StockAI/*", api.AIServeProxy)

	// api
	r_api := app.Group("/api")
	r_api.Get("/wss", wss.WsServer())

	// /api/public
	public.Router(r_api)

	// /api/private
	private.Router(r_api)

	// 默认返回 && 文件服务器
	app.Use(api.Ping)

	listenHost := mStr.Join(":", config.AppInfo.Port)
	global.Log.Println(mStr.Join(`启动服务: http://127.0.0.1`, listenHost))
	app.Listen(listenHost)
}
