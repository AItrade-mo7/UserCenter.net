package api

import (
	"strings"

	"DataCenter.net/server/global/config"
	"DataCenter.net/server/router/result"
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	fastProxy "github.com/yeqown/fasthttp-reverse-proxy"
)

const testAIServeHost = "127.0.0.1:9010"

func AIServeProxy(c *fiber.Ctx) error {
	// 代理 wss
	findWss := strings.Contains(c.Path(), "/wss")
	if findWss {
		return AIServeProxy_wss(c)
	}
	host := c.Get("Coin-Serve-ID")

	if config.SysEnv.RunMod == 1 {
		host = testAIServeHost
	}

	if len(host) < 6 {
		return c.JSON(result.Fail.WithData("缺少代理地址"))
	}

	proxyServer := fastProxy.NewReverseProxy(host)
	proxyServer.ServeHTTP(c.Context())
	return nil
}

func AIServeProxy_wss(c *fiber.Ctx) error {
	host := c.Query("host")
	path := c.Path()
	if len(host) < 6 {
		return c.JSON(result.Fail.WithData("缺少代理地址"))
	}
	if config.SysEnv.RunMod == 1 {
		host = testAIServeHost
	}

	proxyServer := fastProxy.NewWSReverseProxy(host, path)
	proxyServer.Upgrader = &websocket.FastHTTPUpgrader{
		CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
			return true
		},
	}
	proxyServer.ServeHTTP(c.Context())
	return nil
}
