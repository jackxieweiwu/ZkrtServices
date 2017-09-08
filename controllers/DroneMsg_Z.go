package controllers

import (
	"ZkrtServices/models"
	"errors"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"time"
)

// DroneMsgZController operations for DroneMsgZ
type DroneMsgZController struct {
	beego.Controller
}

// URLMapping ...
func (c *DroneMsgZController) URLMapping() {
	c.Mapping("AddDrone", c.AddDroneMessage)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("UpdateDrone", c.DroneMsgUpdate)
	c.Mapping("Delete", c.Delete)
}

// AddDroneMessage ...
// @Title AddDroneMessage
// @Description create DroneMsgZ
// @Param	DroneId	   query   strinng false	   "drone id"
// @Param   DroneAlt   query   float   false       "drone alt"
// @Param   DroneYaw   query   float   false       "drone yaw"
// @Param   DronePitch query   float   false       "drone pitch"
// @Param   DroneSpeed query   float   false       "drone speed"
// @Param   DroneBool  query   int     false       "drone bool"
// @Success 201 {int} models.DroneMsgZ.id
// @Failure 403 body is empty
// @router /AddDrone/:dronemsg [post,get]
func (c *DroneMsgZController) AddDroneMessage() {
	var v models.DroneMsgZ
	strDrone := c.Ctx.Input.Param(":dronemsg")
	droneMessage := strings.Split(strDrone,",")

	droneAlt,_ := strconv.ParseFloat(droneMessage[1],32)
	droneYaw,_ := strconv.ParseFloat(droneMessage[2],32)
	dronePitch,_ := strconv.ParseFloat(droneMessage[3],32)
	droneSpeed,_ := strconv.ParseFloat(droneMessage[4],32)
	droneBool,_ := strconv.Atoi(droneMessage[5])
	v.DroneID = droneMessage[0]
	v.DroneAlt = float32(droneAlt)
	v.DroneYaw = float32(droneYaw)
	v.DronePitch = float32(dronePitch)
	v.DroneSpeed = float32(droneSpeed)
	v.DroneBool = droneBool
	v.DroneDateTime = time.Now()
	v.DroneDate = time.Now()
	if id, err := models.AddDroneMsgZ(&v); err == nil {
		c.Data["json"] = id
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get DroneMsgZ by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.DroneMsgZ
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DroneMsgZController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetDroneMsgZById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get DroneMsgZ
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.DroneMsgZ
// @Failure 403
// @router / [get,post,put]
func (c *DroneMsgZController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllDroneMsgZ(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the DroneMsgZ
// @Param	DroneId	   query   strinng false	   "drone id"
// @Param   DroneAlt   query   float   false       "drone alt"
// @Param   DroneYaw   query   float   false       "drone yaw"
// @Param   DronePitch query   float   false       "drone pitch"
// @Param   DroneSpeed query   float   false       "drone speed"
// @Param   DroneBool  query   int     false       "drone bool"
// @Success 200 {object} models.DroneMsgZ
// @Failure 403 :id is not int
// @router /updateDrone/:updateDroneMessage [get,post,put]
func (c *DroneMsgZController) DroneMsgUpdate() {
	strDrone := c.Ctx.Input.Param(":updateDroneMessage")
	droneMessage := strings.Split(strDrone,",")

	var id int64
	var err error

	if len(droneMessage)==2 {
		droneBool,_ := strconv.Atoi(droneMessage[1])
		id,err =models.UpdateDroneMsgZByDroneIdBool(droneMessage[0],float32(droneBool))
	}else{
		droneAlt,_ := strconv.ParseFloat(droneMessage[1],32)
		droneYaw,_ := strconv.ParseFloat(droneMessage[2],32)
		dronePitch,_ := strconv.ParseFloat(droneMessage[3],32)
		droneSpeed,_ := strconv.ParseFloat(droneMessage[4],32)
		droneBool,_ := strconv.Atoi(droneMessage[5])
		id,err =models.UpdateDroneMsgZByDroneId(droneMessage[0],float32(droneAlt),
			float32(droneYaw),float32(dronePitch),float32(droneSpeed),float32(droneBool))
	}

	if err != nil{
		c.Data["json"] = err.Error()
	}else{
		c.Data["json"] = id
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the DroneMsgZ
// @Param	DroneId	   query   strinng false	   "drone id"
// @Param   DroneBool  query   int     false       "drone bool"
// @Success 200 {object} models.DroneMsgZ
// @Failure 403 :id is not int
// @router /updateDrone/:updateDroneMessage [get,post,put]
func (c *DroneMsgZController) DroneMsgUpdateBool() {
	strDrone := c.Ctx.Input.Param(":updateDroneMessage")
	droneMessage := strings.Split(strDrone,",")
	droneBool,_ := strconv.Atoi(droneMessage[1])

	if id,err :=models.UpdateDroneMsgZByDroneIdBool(droneMessage[0],float32(droneBool)); err != nil{
		c.Data["json"] = err.Error()
	}else{
		c.Data["json"] = id
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the DroneMsgZ
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DroneMsgZController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteDroneMsgZ(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
