package ip

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
)

// 获取ip

func GetIP(c *gin.Context) string {
	// var r *http.Request
	ip := exnet.ClientPublicIP(c.Request)
	if ip == "" {
		ip = exnet.ClientIP(c.Request)
	}
	return ip
}
