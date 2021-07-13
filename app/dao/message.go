// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"battery/app/dao/internal"
)

// messageDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type messageDao struct {
	*internal.MessageDao
}

var (
	// Message is globally public accessible object for table message operations.
	Message messageDao
)

func init() {
	Message = messageDao{
		internal.NewMessageDao(),
	}
}

// Fill with you ideas below.
