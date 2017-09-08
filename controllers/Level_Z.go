package controllers

import (
	"ZkrtServices/models"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"time"
)

// LevelZController operations for LevelZ
type LevelZController struct {
	beego.Controller
}

// URLMapping ...
func (c *LevelZController) URLMapping() {
	c.Mapping("AddLevel", c.AddLevelOne)
	c.Mapping("GetLevelOne", c.GetOne)
	c.Mapping("GetLevelAll", c.GetAll)
	c.Mapping("UpdateLevel", c.UpdateLevel)
	c.Mapping("UpdateLevelId", c.UpdateLevelId)
	c.Mapping("levelDel", c.Delete)
}

// AddLevelOne ...
// @Title AddLevelOne
// @Description create LevelZ
// @Param	LevelName		string 	models.LevelZ	true
// @Param	LevelNumber		int 	models.LevelZ	true
// @Param	LevelUpNumber	int 	models.LevelZ	true
// @Success 201 {int} models.LevelZ
// @Failure 403 body is empty
// @router /AddLevel/:levelName [post,get]
func (c *LevelZController) AddLevelOne() {
	var v models.LevelZ
	levelName := c.Ctx.Input.Param(":levelName")
	strName := strings.Split(levelName,",")
	LevelNumber, _ := strconv.Atoi(strName[1])
	LevelUpNumber, _ := strconv.Atoi(strName[2])

	v.LevelName = strName[0]
	v.LevelNumber = LevelNumber
	v.LevelUpNumber = LevelUpNumber
	v.LevelDateTime = time.Now()
	v.LevelDate = time.Now()

	if id,err := models.AddLevelZ(&v); err != nil{
		c.Data["json"] = err.Error()
	}else{
		c.Data["json"] = id
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get LevelZ by id
// @Param	LevelName		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.LevelZ
// @Failure 403 :id is empty
// @router /GetLevelOne/:LevelName [get,post]
func (c *LevelZController) GetOne() {
	idStr := c.Ctx.Input.Param(":LevelName")
	//id, _ := strconv.Atoi(idStr)
	v, err := models.GetLevelZById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get LevelZ
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.LevelZ
// @Failure 403
// @router /GetLevelAll [get,post]
func (c *LevelZController) GetAll() {
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

	l, err := models.GetAllLevelZ(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// UpdateLevel ...
// @Title UpdateLevel
// @Description update the LevelZ
// @Param	LevelName		string 	models.LevelZ	true
// @Param	LevelNumber		int 	models.LevelZ	true
// @Param	LevelUpNumber	int 	models.LevelZ	true
// @Success 200 {object} models.LevelZ
// @Failure 403 :id is not int
// @router /UpdateLevel/:levelUpdateName [post,get]
func (c *LevelZController) UpdateLevel() {
	levelName := c.Ctx.Input.Param(":levelUpdateName")
	strName := strings.Split(levelName,",")
	LevelNumber, _ := strconv.Atoi(strName[1])
	LevelUpNumber, _ := strconv.Atoi(strName[2])

	if id,err := models.UpdateLevelZNumberByName(strName[0],LevelNumber,LevelUpNumber); err !=nil{
		c.Data["json"] = err.Error()
	}else {
		c.Data["json"] = id
	}
	c.ServeJSON()
}
// UpdateLevel ...
// @Title UpdateLevel
// @Description update the LevelZ
// @Param	LevelNewName    string 	models.LevelZ	true
// @Param	LevelName		string 	models.LevelZ	true
// @Param	LevelNumber		int 	models.LevelZ	true
// @Param	LevelUpNumber	int 	models.LevelZ	true
// @Success 200 {object} models.LevelZ
// @Failure 403 :id is not int
// @router /UpdateLevelId/:levelUpdateNameId [post,get,put]
func (c *LevelZController) UpdateLevelId() {
	levelName := c.Ctx.Input.Param(":levelUpdateNameId")
	strName := strings.Split(levelName,",")
	LevelNumber, _ := strconv.Atoi(strName[2])
	LevelUpNumber, _ := strconv.Atoi(strName[3])

	if id,err := models.UpdateLevelZByIdName(strName[0],strName[1],LevelNumber,LevelUpNumber); err !=nil{
		c.Data["json"] = err.Error()
	}else {
		c.Data["json"] = id
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the LevelZ
// @Param	LevelName		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /levelDel/:LevelName [delete]
func (c *LevelZController) Delete() {
	idStr := c.Ctx.Input.Param(":LevelName")
	//id, _ := strconv.Atoi(idStr)
	if id, err := models.DeleteLevelZ(idStr); err == nil {
		c.Data["json"] = id
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
