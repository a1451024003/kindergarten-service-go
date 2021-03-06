package admin

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/hprose/hprose-golang/rpc"
)

type PingController struct {
	beego.Controller
}

// Ping ...
// @Title Ping
// @Description Ping
// @Success 200		success
// @Failure 403 :没有该服务
// @router / [get]
func (c *PingController) Ping() {
	var balance [2]string

	balance[0] = beego.AppConfig.String("appname")
	balance[1] = time.Now().Format("2006-01-02 15:04:05")

	c.Data["json"] = JSONStruct{"success", 0, balance, "获取成功"}
	c.ServeJSON()
}

// PingUserRpc ...
// @Title user服务是否调通
// @Description user服务是否调通
// @Success 200		success
// @Failure 403 :没有该服务
// @router /rpc/user [get]
func (c *PingController) PingUserRpc() {
	var User *UserService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_USER_SERVER"))
	client.UseService(&User)
	if User.Test() != "" {
		c.Data["json"] = JSONStruct{"success", 0, User.Test(), "获取成功"}
		c.ServeJSON()
	} else {
		c.Data["json"] = JSONStruct{"error", 1, nil, "获取失败"}
		c.ServeJSON()
	}
}

// PingOnemoreRpc ...
// @Title onemore服务是否调通
// @Description onemore服务是否调通
// @Success 200		success
// @Failure 403 :没有该服务
// @router /rpc/onemore [get]
func (c *PingController) PingOnemoreRpc() {
	var Onemore *OnemoreService
	client := rpc.NewHTTPClient(beego.AppConfig.String("ONE_MORE_SMS_SERVER"))
	client.UseService(&Onemore)
	if Onemore.Test() != "" {
		c.Data["json"] = JSONStruct{"success", 0, Onemore.Test(), "获取成功"}
		c.ServeJSON()
	} else {
		c.Data["json"] = JSONStruct{"error", 1, nil, "获取失败"}
		c.ServeJSON()
	}
}
