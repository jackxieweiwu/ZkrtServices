package controllers

import (
	"ZkrtServices/models"
	"errors"
	"strconv"
	"strings"
	"github.com/astaxie/beego"
	"time"
)

// UserZController operations for UserZ
type UserZController struct {
	beego.Controller
}

// URLMapping ...
func (c *UserZController) URLMapping() {
	c.Mapping("AddUser", c.AddUser)
	c.Mapping("GetUserOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Sign", c.SignUser)
	c.Mapping("UserUpdate", c.GetUpdateUserAdminBool)
	c.Mapping("UserDel", c.Delete)
}

// Post ...
// @Title Post
// @Description create UserZ
// @Param	UserName		query 	string	true		"User Name"
// @Param	UserPwd			query 	string	true		"User Password"
// @Param	UserAdminBool	query 	int	true		    "User AdminBool"
// @Success 201 {int} models.UserZ
// @Failure 403 body is empty
// @router /AddUser/:userNamePwd [post,get,put]
func (c *UserZController) AddUser() {
	userNamePwd := c.Ctx.Input.Param(":userNamePwd")
	struser := strings.Split(userNamePwd,",")
	num, _ := strconv.Atoi(struser[2])
	var user models.UserZ
	user.UserName = struser[0]
	user.UserPwd = struser[1]
	user.UserAdminBool = num
	user.LoginTime = time.Now()
	user.UserData = time.Now()

	if v,err := models.AddUserZ(&user); err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get UserZ by userName
// @Param	userName	path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserZ
// @Failure 403 :id is empty
// @router /GetUserOne/:name [get,post,put]
func (c *UserZController) GetOne() {
	idStr := c.Ctx.Input.Param(":name")
	/*id, _ := strconv.Atoi(idStr)*/
	v, err := models.GetUserZById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetUpdateUserAdminBool ...
// @Title GetUpdateUserAdminBool
// @Description get UserZ by userupdate
// @Param	UserName		query 	string	true		"User Name"
// @Param	UserAdminBool	query 	int	true		    "User AdminBool"
// @Success 200 {object} models.UserZ
// @Failure 403 :userupdate is empty
// @router /UserUpdate/:userupdate [get,post,put]
func (c *UserZController) GetUpdateUserAdminBool()  {
	updatecam := c.Ctx.Input.Param(":userupdate")
	user := strings.Split(updatecam,",")
	v, err := models.UpdateAdminBool(user[0],user[1])
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// SetUser ...
// @Title SignUser
// @Description get UserZ by userNamePwd
// @Param	UserName		query 	string	true		"User Password"
// @Param	UserPwd			query 	string	true		"User Password"
// @Success 200 {object} models.UserZ
// @Failure 403 :userNamePwd is empty
// @router /Sign/:userNamePwd [get,post,put]
func (c * UserZController) SignUser()  {
	userNamePwd := c.Ctx.Input.Param(":userNamePwd")
	struser := strings.Split(userNamePwd,",")

	v, err := models.SetUserZSign(struser[0],struser[1])
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get UserZ
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.UserZ
// @Failure 403
// @router / [get]
func (c *UserZController) GetAll() {
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

	l, err := models.GetAllUserZ(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the UserZ
// @Param	userName	path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 userName is empty
// @router /UserDel/:name [delete]
func (c *UserZController) Delete() {
	idStr := c.Ctx.Input.Param(":name")
	//id, _ := strconv.Atoi(idStr)
	if id,err := models.DeleteUserZ(idStr); err == nil {
		c.Data["json"] = id
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
