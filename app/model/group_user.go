// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package model

import (
    "battery/app/model/internal"
)

// GroupUser is the golang structure for table group_user.
type GroupUser internal.GroupUser

// Fill with you ideas below.

type UserInfo struct {
    Id        uint   `json:"id"`        // 用户ID
    GroupName string `json:"groupName"` // 公司名称
    Name      string `json:"name"`      // 姓名
    State     int    `json:"state"`     // 状态

}
