// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ShopBatteryRecord is the golang structure for table shop_battery_record.
type ShopBatteryRecord struct {
	Id          uint64      `orm:"id,primary"  json:"id"`          //
	ShopId      uint        `orm:"shopId"      json:"shopId"`      // 门店ID
	BatteryType string      `orm:"batteryType" json:"batteryType"` // 电池型号
	Type        uint        `orm:"type"        json:"type"`        // 类别: 1入 2出
	BizType     uint        `orm:"bizType"     json:"bizType"`     // 业务类别: 0平台调拨, 其他详细见表user_biz中bizType解释
	BizId       uint64      `orm:"bizId"       json:"bizId"`       // 业务ID: 0为平台调拨
	Num         int         `orm:"num"         json:"num"`         // 数量
	Date        *gtime.Time `orm:"date"        json:"date"`        // 日期
	UserId      uint64      `orm:"userId"      json:"userId"`      // 骑手ID
	UserName    string      `orm:"userName"    json:"userName"`    // 骑手名字
	SysUserId   uint        `orm:"sysUserId"   json:"sysUserId"`   // 系统管理员ID
	SysUserName string      `orm:"sysUserName" json:"sysUserName"` // 操作员名字
	CreatedAt   *gtime.Time `orm:"createdAt"   json:"createdAt"`   //
}
