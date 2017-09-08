package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type LevelZ struct {
	Id            int       `orm:"column(LevelId);auto"`
	LevelName     string    `orm:"column(LevelName);size(100)"`
	LevelNumber   int       `orm:"column(LevelNumber)"`
	LevelUpNumber int       `orm:"column(LevelUpNumber)"`
	LevelDateTime time.Time `orm:"column(LevelDateTime);type(datetime)"`
	LevelDate     time.Time `orm:"column(LevelDate);type(date)"`
}

func (t *LevelZ) TableName() string {
	return "Level_Z"
}

func init() {
	orm.RegisterModel(new(LevelZ))
}

// AddLevelZ insert a new LevelZ into database and returns
// last inserted Id on success.
func AddLevelZ(m *LevelZ) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLevelZById retrieves LevelZ by Id. Returns error if
// Id doesn't exist
func GetLevelZById(levelname string) (v LevelZ,err error) {
	o := orm.NewOrm()
	var level LevelZ
	err = o.QueryTable(new(LevelZ)).Filter("LevelName", levelname).One(&level)
	if err == nil {
		return level, nil
	}
	return level,err
}

// GetAllLevelZ retrieves all LevelZ matches certain condition. Returns empty list if
// no records exist
func GetAllLevelZ(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LevelZ))
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

	var l []LevelZ
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

// UpdateLevelZ updates LevelZ by Id and returns error if
// the record to be updated doesn't exist
func UpdateLevelZNumberByName(level string,levelnumber int,levelupnumber int) (id int64,err error) {
	o := orm.NewOrm()
	id, err = o.QueryTable(new(LevelZ)).Filter("LevelName", level).Update(orm.Params{
		"LevelNumber": levelnumber,
		"LevelUpNumber": levelupnumber,
	})
	return
}

func UpdateLevelZByIdName(newName string,level string,levelnumber int,levelupnumber int) (id int64,err error) {
	o := orm.NewOrm()
	id, err = o.QueryTable(new(LevelZ)).Filter("LevelName", level).Update(orm.Params{
		"LevelNumber": levelnumber,
		"LevelUpNumber": levelupnumber,
		"LevelName": newName,
	})
	return
}

// DeleteLevelZ deletes LevelZ by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLevelZ(levelName string) (id int64,err error) {
	o := orm.NewOrm()
	id ,err = o.QueryTable(new(LevelZ)).Filter("UserName", levelName).Delete()
	return
}
