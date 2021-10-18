// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// SignDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type SignDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB      // DB is the raw underlying database management object.
	Table   string      // Table is the table name of the DAO.
	Columns signColumns // Columns contains all the columns of Table that for convenient usage.
}

// SignColumns defines and stores column names for table sign.
type signColumns struct {
	Id           string //
	UserId       string //
	GroupId      string //
	ComboOrderId string //
	BatteryType  string // 电池型号
	FlowId       string // 易签签约流程ID
	FileId       string // 易签签约文件ID
	State        string // 签约状态 0 未签约 1 已签约
	CreatedAt    string //
	UpdatedAt    string //
}

func NewSignDao() *SignDao {
	return &SignDao{
		M:     g.DB("default").Model("sign").Safe(),
		DB:    g.DB("default"),
		Table: "sign",
		Columns: signColumns{
			Id:           "id",
			UserId:       "userId",
			GroupId:      "groupId",
			ComboOrderId: "comboOrderId",
			BatteryType:  "batteryType",
			FlowId:       "flowId",
			FileId:       "fileId",
			State:        "state",
			CreatedAt:    "createdAt",
			UpdatedAt:    "updatedAt",
		},
	}
}
