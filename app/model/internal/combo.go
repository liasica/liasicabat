// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Combo is the golang structure for table combo.
type Combo struct {
	Id          uint        `orm:"id,primary"  json:"id"`          //
	Type        uint        `orm:"type"        json:"type"`        // 套餐类型: 1个人 2团体
	BatteryType string      `orm:"batteryType" json:"batteryType"` // 电池型号
	Name        string      `orm:"name"        json:"name"`        // 名称
	Days        uint        `orm:"days"        json:"days"`        // 套餐时长天数
	Amount      float64     `orm:"amount"      json:"amount"`      // 套餐总价(包含押金)
	Price       float64     `orm:"price"       json:"price"`       // 计费价格
	Deposit     float64     `orm:"deposit"     json:"deposit"`     // 押金
	ProvinceId  uint        `orm:"provinceId"  json:"provinceId"`  // 省级行政编码
	CityId      uint        `orm:"cityId"      json:"cityId"`      // 市级行政编码
	Desc        string      `orm:"desc"        json:"desc"`        // 描述
	CreatedAt   *gtime.Time `orm:"createdAt"   json:"createdAt"`   //
	UpdatedAt   *gtime.Time `orm:"updatedAt"   json:"updatedAt"`   //
	DeletedAt   *gtime.Time `orm:"deletedAt"   json:"deletedAt"`   // 停用时间
}
