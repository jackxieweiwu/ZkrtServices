package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type MoudleZ struct {
	Id             int       `orm:"column(MoudleID);auto"`
	DroneID        string    `orm:"column(DroneID);size(100)"`
	ModuleName     string    `orm:"column(ModuleName);size(225)"`
	MoudleBool     int       `orm:"column(MoudleBool)"`
	MoudleDateTime time.Time `orm:"column(MoudleDateTime);type(datetime)"`
	MoudleDate     time.Time `orm:"column(MoudleDate);type(date)"`
}

func (t *MoudleZ) TableName() string {
	return "Moudle_Z"
}

func init() {
	orm.RegisterModel(new(MoudleZ))
}

// AddMoudleZ insert a new MoudleZ into database and returns
// last inserted Id on success.
func AddMoudleZ(m *MoudleZ) (id int64, err error) {
	o := orm.NewOrm()
	user := MoudleZ{DroneID:m.DroneID,ModuleName:m.ModuleName,MoudleBool:m.MoudleBool,MoudleDateTime:m.MoudleDateTime,MoudleDate: m.MoudleDate}
	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	if created, id, err := o.ReadOrCreate(&user, "MoudleDate"); err == nil {
		if created {
			return id,nil
		} else {
			return 0,err
		}
	}
	//id, err = o.Insert(m)
	return 0,err
}

// GetMoudleZById retrieves MoudleZ by Id. Returns error if
// Id doesn't exist
func GetMoudleZById(id int) (v *MoudleZ, err error) {
	o := orm.NewOrm()
	v = &MoudleZ{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMoudleZ retrieves all MoudleZ matches certain condition. Returns empty list if
// no records exist
func GetAllMoudleZ(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MoudleZ))
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

	var l []MoudleZ
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

// UpdateMoudleZ updates MoudleZ by Id and returns error if
// the record to be updated doesn't exist
func UpdateMoudleZById(m *MoudleZ) (err error) {
	o := orm.NewOrm()
	v := MoudleZ{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMoudleZ deletes MoudleZ by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMoudleZ(id int) (err error) {
	o := orm.NewOrm()
	v := MoudleZ{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MoudleZ{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
