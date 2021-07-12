// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"battery/app/dao/internal"
)

// exceptionDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type exceptionDao struct {
	*internal.ExceptionDao
}

var (
	// Exception is globally public accessible object for table exception operations.
	Exception exceptionDao
)

func init() {
	Exception = exceptionDao{
		internal.NewExceptionDao(),
	}
}

// Fill with you ideas below.
