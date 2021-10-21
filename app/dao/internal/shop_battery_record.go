// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// ShopBatteryRecordDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type ShopBatteryRecordDao struct {
	gmvc.M                           // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                   // DB is the raw underlying database management object.
	Table   string                   // Table is the table name of the DAO.
	Columns shopBatteryRecordColumns // Columns contains all the columns of Table that for convenient usage.
}

// ShopBatteryRecordColumns defines and stores column names for table shop_battery_record.
type shopBatteryRecordColumns struct {
	Id          string //
	ShopId      string // 门店ID
	BatteryType string // 电池型号
	Type        string // 类别: 1入 2出
	BizType     string // 业务类别: 0平台调拨, 其他详细见表user_biz中bizType解释
	BizId       string // 业务ID: 0为平台调拨
	Num         string // 数量
	Date        string // 日期
	UserId      string // 骑手ID
	UserName    string // 骑手名字
	SysUserId   string // 系统管理员ID
	SysUserName string // 操作员名字
	CreatedAt   string //
}

func NewShopBatteryRecordDao() *ShopBatteryRecordDao {
	return &ShopBatteryRecordDao{
		M:     g.DB("default").Model("shop_battery_record").Safe(),
		DB:    g.DB("default"),
		Table: "shop_battery_record",
		Columns: shopBatteryRecordColumns{
			Id:          "id",
			ShopId:      "shopId",
			BatteryType: "batteryType",
			Type:        "type",
			BizType:     "bizType",
			BizId:       "bizId",
			Num:         "num",
			Date:        "date",
			UserId:      "userId",
			UserName:    "userName",
			SysUserId:   "sysUserId",
			SysUserName: "sysUserName",
			CreatedAt:   "createdAt",
		},
	}
}
