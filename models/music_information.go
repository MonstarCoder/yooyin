package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MusicInformation struct {
	Id              int    `orm:"column(id);auto"`
	Name            string `orm:"column(name);size(512);null" description:"歌名或者歌手名称或者标签"`
	Type            int    `orm:"column(type);null" description:"类型，具体自己定义"`
	AddtionalFields string `orm:"column(addtional_fields);null" description:"json字符串，此处需要前端定义不同类型需要展现的字段以及爬虫能入库的信息"`
}

func (t *MusicInformation) TableName() string {
	return "music_information"
}

func init() {
	orm.RegisterModel(new(MusicInformation))
}

// AddMusicInformation insert a new MusicInformation into database and returns
// last inserted Id on success.
func AddMusicInformation(m *MusicInformation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMusicInformationById retrieves MusicInformation by Id. Returns error if
// Id doesn't exist
func GetMusicInformationById(id int) (v *MusicInformation, err error) {
	o := orm.NewOrm()
	v = &MusicInformation{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetMusicInformationsCountByNameType(name string, info_type int) (int64, error) {
	o := orm.NewOrm()
	cnt, err := o.QueryTable(new(MusicInformation).TableName()).Filter("type", info_type).Filter("name__icontains", name).OrderBy("-id").Count()
	return cnt, err
}

func GetMusicInformationsByNameType(musicInfo *[]MusicInformation, name string, info_type int, offset int64, limit int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable(new(MusicInformation).TableName()).Filter("type", info_type).Filter("name__icontains", name).OrderBy("-id").Offset(offset).Limit(limit).All(musicInfo)
	return err
}

// GetAllMusicInformation retrieves all MusicInformation matches certain condition. Returns empty list if
// no records exist
func GetAllMusicInformation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MusicInformation))
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

	var l []MusicInformation
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