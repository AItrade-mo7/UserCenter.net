package api

import (
	"strings"

	"UserCenter.net/server/global/middle"
	"UserCenter.net/server/router/result"
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	fastProxy "github.com/yeqown/fasthttp-reverse-proxy"
)

func AIServeProxy(c *fiber.Ctx) error {
	isCrawler := middle.CrawlerIS(c)
	if isCrawler {
		return c.JSON(result.ErrLogin.With("访问失败", "设备异常"))
	}

	fastProxy.SetProduction()
	// 代理 wss
	findWss := strings.Contains(c.Path(), "/wss")
	if findWss {
		return AIServeProxy_wss(c)
	}
	host := c.Get("Satellite-Serve")

	if len([]rune(host)) < 6 {
		return c.JSON(result.Fail.WithMsg("缺少代理地址"))
	}

	proxyServer := fastProxy.NewReverseProxy(host)
	proxyServer.ServeHTTP(c.Context())
	return nil
}

func AIServeProxy_wss(c *fiber.Ctx) error {
	host := c.Query("host")
	path := c.Path()
	if len([]rune(host)) < 6 {
		return c.JSON(result.Fail.WithMsg("缺少代理地址"))
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
