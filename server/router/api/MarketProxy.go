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

const testMarketHost = "127.0.0.1:8998"

func MarketProxy(c *fiber.Ctx) error {
	// 代理 wss
	findWss := strings.Contains(c.Path(), "/wss")
	if findWss {
		return MarketProxy_wss(c)
	}
	host := "trade.mo7.cc"
	if config.SysEnv.RunMod == 1 {
		host = testMarketHost
	}

	if len(host) < 6 {
		return c.JSON(result.Fail.WithData("缺少代理地址"))
	}
	proxyServer := fastProxy.NewReverseProxy(host)
	proxyServer.ServeHTTP(c.Context())
	return nil
}

func MarketProxy_wss(c *fiber.Ctx) error {
	host := "trade.mo7.cc"
	path := c.Path()
	if config.SysEnv.RunMod == 1 {
		host = testMarketHost
	}

	if len(host) < 6 {
		return c.JSON(result.Fail.WithData("缺少代理地址"))
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
