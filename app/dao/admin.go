// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"battery/app/dao/internal"
)

// adminDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type adminDao struct {
	*internal.AdminDao
}

var (
	// Admin is globally public accessible object for table admin operations.
	Admin adminDao
)

func init() {
	Admin = adminDao{
		internal.NewAdminDao(),
	}
}

// Fill with you ideas below.
