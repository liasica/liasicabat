// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package model

import (
    "battery/app/model/internal"
    "github.com/gogf/gf/os/gtime"
    "github.com/gogf/gf/util/gmeta"
)

// Group is the golang structure for table group.
type Group internal.Group

// Fill with you ideas below.

// GroupListItem 团签详情
type GroupListItem struct {
    Id               uint   `json:"id"`
    Name             string `json:"name"`             // 名称
    ProvinceId       uint   `json:"provinceId"`       // 省份ID
    CityId           uint   `json:"cityId"`           // 城市ID
    UserCnt          uint   `json:"userCnt"`          // 用户数量
    DaysCnt60        uint   `json:"daysCnt60"`        // 60电池天数
    DaysCnt72        uint   `json:"daysCnt72"`        // 72电池天数
    ArrearsDaysCnt60 uint   `json:"arrearsDaysCnt60"` // 60电池欠款天数
    ArrearsDaysCnt72 uint   `json:"arrearsDaysCnt72"` // 72电池欠款天数
    ContactName      string `json:"contactName"`      // 联系人
    ContactMobile    string `json:"contactMobile"`    // 联系电话
}

// GroupListReq 列表请求参数
type GroupListReq struct {
    GroupListAdminReq
    StartDate *gtime.Time `json:"startDate"` // 开始日期
    EndDate   *gtime.Time `json:"endDate"`   // 结束日期
}

// GroupFormReq 团签请求表单
type GroupFormReq struct {
    Name          string               `v:"required" json:"name"`                     // 名称
    ContactName   string               `v:"required" json:"contactName"`              // 联系人
    ContactMobile string               `v:"required" json:"contactMobile"`            // 联系人手机号
    ProvinceId    uint                 `v:"required|integer|min:1" json:"provinceId"` // 省份ID
    CityId        uint                 `v:"required|integer|min:1" json:"cityId"`     // 城市ID
    UserList      []GroupCreateUserReq `json:"userList"`                              // 用户列表
    ContractFile  string               `v:"required" json:"contractFile"`             // 合同文件 (使用接口`/api/upload/file`上传)
}

// GroupCreateUserReq 团签用户表单
type GroupCreateUserReq struct {
    Name     string `v:"required|length:2,30" json:"name" validate:"required"`     // 姓名
    Mobile   string `v:"required|phone-loose" json:"mobile" validate:"required"`   // 电话
    IdCardNo string `v:"required|resident-id" json:"idCardNo" validate:"required"` // 身份证
}

// UserGroupStatRep 我的团队统计
type UserGroupStatRep struct {
    MemberCnt uint `json:"memberCnt"` // 团队成员人数
    BillDays  uint `json:"billDays"`  // 未结算天数
}

// UserGroupListRep 我的团队详情列表
type UserGroupListRep struct {
    RealName string `json:"realName"` // 姓名
    Days     uint   `json:"days"`     // 总使用天数
    BillDays uint   `json:"billDays"` // 未结算天数
}

// GroupListAdminReq 管理列表请求数据
type GroupListAdminReq struct {
    Page
    Keywords string `json:"keywords"`
}

// GroupEntity 团签详情实体
type GroupEntity struct {
    gmeta.Meta `orm:"table:group_settlement_detail" swaggerignore:"true"`

    Id            uint   `json:"id"`
    Name          string `json:"name"`          // 名称
    ProvinceId    uint   `json:"provinceId"`    // 省份ID
    CityId        uint   `json:"cityId"`        // 城市ID
    ContactName   string `json:"contactName"`   // 联系人
    ContactMobile string `json:"contactMobile"` // 联系电话

    MemberCnt uint   `json:"memberCnt"` // 骑手数量
    CityName  string `json:"cityName"`  // 城市名称
    Days      uint   `json:"days"`      // 累积天数
    BillDays  uint   `json:"billDays"`  // 未结天数

    SettlementDetails []*SettlementDetailEntity `json:"-" orm:"with:groupId=id, where:ignorance=0, order:startDate desc"`
    City              *Districts                `json:"-" orm:"with:id=cityId"`
}

// GroupUsageDays 团队使用天数
type GroupUsageDays struct {
    Days     uint `json:"days"`     // 总使用天数
    BillDays uint `json:"billDays"` // 未结算天数
}
