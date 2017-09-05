package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DMoudleZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:DroneMsgZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GasresZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:GroupZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LevelZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:LogZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:MoudleZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:TraceZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:UserZController"],
		beego.ControllerComments{
			Method: "SignUser",
			Router: `/Sign`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"] = append(beego.GlobalControllerRouter["ZkrtServices/controllers:VideoZController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
