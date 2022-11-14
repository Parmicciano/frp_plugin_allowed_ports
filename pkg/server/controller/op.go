package controller

import (
	"fmt"
	"net/http"
	"strings"

	plugin "github.com/fatedier/frp/pkg/plugin/server"
	"github.com/gin-gonic/gin"

	"strconv"
)

type OpController struct {
	ports map[string][]string
}

func NewOpController(ports map[string][]string) *OpController {
	return &OpController{
		ports: ports,
	}
}

func (c *OpController) Register(engine *gin.Engine) {
	engine.POST("/handler", MakeGinHandlerFunc(c.HandleLogin))
}
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (c *OpController) HandleLogin(ctx *gin.Context) (interface{}, error) {
	var r plugin.Request
	var content plugin.NewProxyContent
	var res plugin.Response

	r.Content = &content
	if err := ctx.BindJSON(&r); err != nil {
		return nil, &HTTPError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	fmt.Println("-------------Plugin: Allowed Ports--------------------")
	fmt.Printf("ProxyName: %s\tProxyType%s\t", content.ProxyName, content.ProxyType)
	if strings.ToLower(content.ProxyType) == "tcp" || strings.ToLower(content.ProxyType) == "udp" {
		fmt.Printf("RemotePort: %d\r\n", content.RemotePort)
	} else if strings.HasPrefix(content.ProxyType, "http") {
		fmt.Printf("CustomDomains%s\r\n", content.CustomDomains)
	} else {
		fmt.Println("Won't do validation for this type")
		res.Unchange = true
		return res, nil
	}

	subdomain := content.SubDomain
	remoteport := strconv.Itoa(content.RemotePort)
	username := content.User.User

	if subdomain == "" && remoteport == "0" && len(content.CustomDomains) == 0 {
		res.Reject = true
		res.RejectReason = "Rejected due to misconfiguration of the client"
	}

	find := false

	for _, port_allowed := range c.ports[username] {
		if port_allowed == remoteport || port_allowed == subdomain {
			find = true
		}

		if contains(content.CustomDomains, port_allowed) {
			find = true
		}
	}
	if !find {
		res.Reject = true
		res.RejectReason = "Client is not allowed => Port or subdomain false"
	}

	if !res.Reject {
		res.Unchange = true
	}
	return res, nil
}
