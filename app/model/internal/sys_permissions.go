// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// SysPermissions is the golang structure for table sys_permissions.
type SysPermissions struct {
	Id       int         `orm:"id,primary" json:"id"`       //
	ParentId int         `orm:"parent_id"  json:"parentId"` // 上级权限ID
	Method   string      `orm:"method"     json:"method"`   // HTTP请求方法 大写
	Path     string      `orm:"path"       json:"path"`     // 访问资源路径
	Name     string      `orm:"name"       json:"name"`     //
	Vtype    int         `orm:"vtype"      json:"vtype"`    // 前端路由类型  1 菜单 2 按钮
	Vpath    string      `orm:"vpath"      json:"vpath"`    // Vue 路由
	CreateAt *gtime.Time `orm:"create_at"  json:"createAt"` //
	UpdateAt *gtime.Time `orm:"update_at"  json:"updateAt"` //
}
