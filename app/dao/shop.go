// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"battery/app/dao/internal"
)

// shopDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type shopDao struct {
	*internal.ShopDao
}

var (
	// Shop is globally public accessible object for table shop operations.
	Shop shopDao
)

func init() {
	Shop = shopDao{
		internal.NewShopDao(),
	}
}

// Fill with you ideas below.
