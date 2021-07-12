// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// GroupUserDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type GroupUserDao struct {
	gmvc.M                   // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB           // DB is the raw underlying database management object.
	Table   string           // Table is the table name of the DAO.
	Columns groupUserColumns // Columns contains all the columns of Table that for convenient usage.
}

// GroupUserColumns defines and stores column names for table group_user.
type groupUserColumns struct {
	Id        string //
	GroupId   string // 团队ID
	UserId    string // 用户ID
	DeletedAt string // 注销时间
	CreatedAt string //
	UpdatedAt string //
}

func NewGroupUserDao() *GroupUserDao {
	return &GroupUserDao{
		M:     g.DB("default").Model("group_user").Safe(),
		DB:    g.DB("default"),
		Table: "group_user",
		Columns: groupUserColumns{
			Id:        "id",
			GroupId:   "groupId",
			UserId:    "userId",
			DeletedAt: "deletedAt",
			CreatedAt: "createdAt",
			UpdatedAt: "updatedAt",
		},
	}
}
