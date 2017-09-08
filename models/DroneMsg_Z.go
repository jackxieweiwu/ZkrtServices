package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type DroneMsgZ struct {
	Id            int       `orm:"column(DId);auto"`
	DroneID       string    `orm:"column(DroneID);size(100)"`
	DroneAlt      float32   `orm:"column(DroneAlt)"`
	DroneYaw      float32   `orm:"column(DroneYaw)"`
	DronePitch    float32   `orm:"column(DronePitch)"`
	DroneSpeed    float32   `orm:"column(DroneSpeed)"`
	DroneBool     int       `orm:"column(DroneBool)"`
	DroneDateTime time.Time `orm:"column(DroneDateTime);type(datetime)"`
	DroneDate     time.Time `orm:"column(DroneDate);type(time)"`
}

func (t *DroneMsgZ) TableName() string {
	return "DroneMsg_Z"
}

func init() {
	orm.RegisterModel(new(DroneMsgZ))
}

// AddDroneMsgZ insert a new DroneMsgZ into database and returns
// last inserted Id on success.
func AddDroneMsgZ(m *DroneMsgZ) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDroneMsgZById retrieves DroneMsgZ by Id. Returns error if
// Id doesn't exist
func GetDroneMsgZById(id int) (v *DroneMsgZ, err error) {
	o := orm.NewOrm()
	v = &DroneMsgZ{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDroneMsgZ retrieves all DroneMsgZ matches certain condition. Returns empty list if
// no records exist
func GetAllDroneMsgZ(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DroneMsgZ))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []DroneMsgZ
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateDroneMsgZ updates DroneMsgZ by Id and returns error if
// the record to be updated doesn't exist
func UpdateDroneMsgZByDroneId(droneId string,droneAlt float32,
	droneYaw float32,dronePitch float32,droneSpeed float32,droneBool float32) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.QueryTable(new(DroneMsgZ)).Filter("DroneID", droneId).Update(orm.Params{
		"DroneAlt": droneAlt,"DroneYaw": droneYaw,"DronePitch": dronePitch,
		"DroneSpeed": droneSpeed,"DroneBool": droneBool,
	})
	return
}

// UpdateDroneMsgZ updates DroneMsgZ by Id and returns error if
// the record to be updated doesn't exist
func UpdateDroneMsgZByDroneIdBool(droneId string,droneBool float32) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.QueryTable(new(DroneMsgZ)).Filter("DroneID", droneId).Update(orm.Params{
		"DroneBool": droneBool,
	})
	return
}

// DeleteDroneMsgZ deletes DroneMsgZ by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDroneMsgZ(id int) (err error) {
	o := orm.NewOrm()
	v := DroneMsgZ{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DroneMsgZ{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
