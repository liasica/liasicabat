// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// GroupSettlementDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type GroupSettlementDao struct {
	gmvc.M                         // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                 // DB is the raw underlying database management object.
	Table   string                 // Table is the table name of the DAO.
	Columns groupSettlementColumns // Columns contains all the columns of Table that for convenient usage.
}

// GroupSettlementColumns defines and stores column names for table group_settlement.
type groupSettlementColumns struct {
	Id        string //
	Hash      string // 结算单hash
	GroupId   string // 团队ID
	Amount    string // 总结算金额
	SysUserId string // 结算人ID
	SysName   string // 结算人姓名
	Date      string // 结算日
	Remark    string // 备注
	Detail    string // 结算详情
	CreatedAt string //
	UpdatedAt string //
}

func NewGroupSettlementDao() *GroupSettlementDao {
	return &GroupSettlementDao{
		M:     g.DB("default").Model("group_settlement").Safe(),
		DB:    g.DB("default"),
		Table: "group_settlement",
		Columns: groupSettlementColumns{
			Id:        "id",
			Hash:      "hash",
			GroupId:   "groupId",
			Amount:    "amount",
			SysUserId: "sysUserId",
			SysName:   "sysName",
			Date:      "date",
			Remark:    "remark",
			Detail:    "detail",
			CreatedAt: "createdAt",
			UpdatedAt: "updatedAt",
		},
	}
}
