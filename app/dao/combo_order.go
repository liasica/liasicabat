// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"battery/app/dao/internal"
)

// comboOrderDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type comboOrderDao struct {
	*internal.ComboOrderDao
}

var (
	// ComboOrder is globally public accessible object for table combo_order operations.
	ComboOrder comboOrderDao
)

func init() {
	ComboOrder = comboOrderDao{
		internal.NewComboOrderDao(),
	}
}

// Fill with you ideas below.
