package controller

import (
	"busuanzi/core"
	"github.com/gin-gonic/gin"
	"net/url"
)

func ApiHandler(c *gin.Context) { // test redisHelper
	var referer = c.Request.Header.Get("x-bsz-referer")
	if referer == "" {
		c.JSON(200, gin.H{
			"success": false,
			"message": "empty referer",
			"data":    gin.H{},
		})
		return
	}

	u, err := url.Parse(referer)
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "unable to parse referer",
			"data":    gin.H{},
		})
		return
	}

	var host = u.Host
	var path = u.Path
	var ip = c.ClientIP()

	// count
	sitePv, siteUv, pagePv, pageUv := core.Count(host, path, ip)

	// json
	c.JSON(200, gin.H{
		"success": true,
		"message": "ok",
		"data": gin.H{
			"site_pv": sitePv,
			"site_uv": siteUv,
			"page_pv": pagePv,
			"page_uv": pageUv,
		},
	})
}