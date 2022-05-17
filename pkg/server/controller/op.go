package controller

import (
	"net/http"
	"fmt"
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

func (c *OpController) HandleLogin(ctx *gin.Context) (interface{}, error) {

	var r plugin.Request 

	var content plugin.NewProxyContent
	
	var res plugin.Response

	var find bool

	r.Content = &content
	if err := ctx.BindJSON(&r); err != nil {
		return nil, &HTTPError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}
	

	//fmt.Println("---------------------------------")
	//fmt.Println(content.ProxyName)
	//fmt.Println(content.ProxyType)
	//fmt.Println(content.RemotePort)
	//fmt.Println(content.CustomDomains)
	//fmt.Println(content.SubDomain)
	//fmt.Println("---------------------------------")

	subdomain  := content.SubDomain
	remoteport := strconv.Itoa(content.RemotePort)
	username := content.User.User

	
	if subdomain == "" && remoteport == "0"{
		fmt.Println("Rejected")
		res.Reject = true
		res.RejectReason = "Misconfiguration of the client"
	}

	find = false

	for _, port_allowed := range c.ports[username] {
	  if port_allowed == remoteport || port_allowed == subdomain {
		find = true

	}

    }
	if find == false {
	
		fmt.Println("Client is not allowed to use this port")
		res.Reject = true
		res.RejectReason = "Not allowed => Port or subdomain false"
		
	}
	
	if res.Reject != true {
		res.Unchange = true
	}
	return res, nil
}

