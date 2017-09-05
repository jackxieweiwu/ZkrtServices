// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ZkrtServices/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/DMoudle_Z",
			beego.NSInclude(
				&controllers.DMoudleZController{},
			),
		),

		beego.NSNamespace("/DroneMsg_Z",
			beego.NSInclude(
				&controllers.DroneMsgZController{},
			),
		),

		beego.NSNamespace("/Gas_Z",
			beego.NSInclude(
				&controllers.GasZController{},
			),
		),

		beego.NSNamespace("/Gasres_Z",
			beego.NSInclude(
				&controllers.GasresZController{},
			),
		),

		beego.NSNamespace("/Group_Z",
			beego.NSInclude(
				&controllers.GroupZController{},
			),
		),

		beego.NSNamespace("/Level_Z",
			beego.NSInclude(
				&controllers.LevelZController{},
			),
		),

		beego.NSNamespace("/Log_Z",
			beego.NSInclude(
				&controllers.LogZController{},
			),
		),

		beego.NSNamespace("/Moudle_Z",
			beego.NSInclude(
				&controllers.MoudleZController{},
			),
		),

		beego.NSNamespace("/Trace_Z",
			beego.NSInclude(
				&controllers.TraceZController{},
			),
		),

		beego.NSNamespace("/User_Z",
			beego.NSInclude(
				&controllers.UserZController{},
			),
		),

		beego.NSNamespace("/Video_Z",
			beego.NSInclude(
				&controllers.VideoZController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
