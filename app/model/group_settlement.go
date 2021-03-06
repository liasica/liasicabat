// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package model

import (
    "battery/app/model/internal"
    "github.com/qiniu/qmgo/field"
)

// GroupSettlement is the golang structure for table group_settlement.
type GroupSettlement internal.GroupSettlement

// Fill with you ideas below.

// SettlementCache 结算单缓存
type SettlementCache struct {
    GroupId  uint                  `json:"groupId" bson:"groupId"`   // 团签ID
    Hash     string                `json:"hash" bson:"-"`            // 结算单hash
    Amount   float64               `json:"amount" bson:"amount"`     // 总金额
    FromDate string                `json:"fromDate" bson:"fromDate"` // 上次结算日
    ToDate   string                `json:"toDate" bson:"toDate"`     // 结算日
    Items    []*SettlementListItem `json:"items" bson:"items"`       // 结算单详情

    field.DefaultField `json:"-" bson:",inline" swaggerignore:"true"`
}

// GroupSettlementCheckoutReq 结算请求
type GroupSettlementCheckoutReq struct {
    Hash   string `json:"hash" validate:"required" v:"required|length:24,24"` // 结算单hash
    Remark string `json:"remark"`                                             // 结算单备注
}
