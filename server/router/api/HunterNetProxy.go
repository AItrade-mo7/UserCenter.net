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

func AIFundNetProxy(c *fiber.Ctx) error {
	fastProxy.SetProduction() // 关闭 debug
	// 代理 wss
	findWss := strings.Contains(c.Path(), "/wss")
	if findWss {
		return AIFundNetProxy_WSS(c)
	}
	host := c.Get("AIFund-Net-Host")

	if config.SysEnv.RunMod == 1 && host == "50.18.29.218:9010" {
		host = "127.0.0.1:9010"
	}

	proxyServer := fastProxy.NewReverseProxy(host)

	if len(host) < 6 {
		return c.JSON(result.Fail.WithData("缺少代理地址"))
	}
	proxyServer.ServeHTTP(c.Context())
	return nil
}

func AIFundNetProxy_WSS(c *fiber.Ctx) error {
	host := c.Query("host")
	path := c.Path()
	if len(host) < 6 {
		return c.JSON(result.Fail.WithData("缺少代理地址"))
	}

	if config.SysEnv.RunMod == 1 && host == "50.18.29.218:9010" {
		host = "127.0.0.1:9010"
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
