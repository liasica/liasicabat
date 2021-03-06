package model

import (
    "github.com/gogf/gf/os/gtime"
    "github.com/gogf/gf/util/gmeta"

    "battery/app/model/internal"
)

const (
    UserBizNew              = 1 // 新签
    UserBizBatteryRenewal   = 2 // 换电池
    UserBizBatteryPause     = 3 // 寄存暂停
    UserBizBatteryRetrieval = 4 // 恢复计费
    UserBizCancel           = 5 // 退租
)

// UserBiz is the golang structure for table user_biz.
type UserBiz internal.UserBiz

// UserBizReq 业务办理请求数据
type UserBizReq struct {
    Code string `validate:"required" v:"required"`            // 用户二维码code
    Type uint   `validate:"required" v:"required|in:2,3,4,5"` // 业务类型:  2 换电 3 寄存(仅个签用户使用)， 4 恢复计费，5 退租
}

// UserBizRecordListRep 骑手获取换电记录响应
type UserBizRecordListRep struct {
    ShopName string      `validate:"required" json:"shopName"` // 门店名称
    ScanAt   *gtime.Time `validate:"required" json:"scanAt"`   // 扫码时间
    CityName string      `validate:"required" json:"cityName"` // 城市名称
}

// UserBizBatteryRenewalReq 骑手扫码换电请求
type UserBizBatteryRenewalReq struct {
    Code string `validate:"required" v:"required" json:"code"` // 门店二维码扫码获得
}

// UserBizBatteryRenewalRep 骑手扫码换电响应
type UserBizBatteryRenewalRep struct {
    ShopName    string      `validate:"required" json:"shopName"`    // 门店名称
    BatteryType string      `validate:"required" json:"batteryType"` // 电池型号
    At          *gtime.Time `validate:"required" json:"at"`          // 时间
}

// UserBizRecordStatRep 骑手换电统计
type UserBizRecordStatRep struct {
    Count uint `validate:"required" json:"count"` // 累计换电次数
    Days  uint `validate:"required" json:"days"`  // 累计使用天数
}

// UserBizSignReq 个签骑手新签套餐请求数据
type UserBizSignReq struct {
    ComboId uint `validate:"required" json:"comboId" v:"required|integer"` // 套餐ID
}

// UserBizNewReq 个签骑手签约之后获取支付信息
type UserBizNewReq struct {
    PayType uint   `validate:"required" json:"payType" v:"required|integer|in:1,2"` // 支付方式 1 支付宝 2 微信支付
    FlowId  string `validate:"required" json:"flowId" v:"required"`                 // 签约流程ID
}

// UserBizNewRep 个签骑手新签套餐响应数据
type UserBizNewRep struct {
    OrderId      uint64 `validate:"required" json:"orderId"`      // 订单ID
    PayOrderInfo string `validate:"required" json:"PayOrderInfo"` // 发起支付使用数据

}

// UserBizNewComboOrderStateRep 个签骑手获取支付状态响应数据
type UserBizNewComboOrderStateRep struct {
    PayState uint `json:"payState"` // 1 待支付 2 已支付
}

// UserBizRenewalReq 个签骑手续约请求数据
type UserBizRenewalReq struct {
    PaymentType uint `validate:"required" json:"paymentType" v:"required|in:1,2"` // 支付类型 1 支付宝 2 微信
}

// UserBizRenewalRep 个签骑手续约响应数据
type UserBizRenewalRep struct {
    OrderId      uint64 `validate:"required" json:"orderId"`      // 订单ID
    PayOrderInfo string `validate:"required" json:"PayOrderInfo"` // 发起支付使用数据
}

// UserBizPenaltyReq 个签骑手支付违约金请求数据
type UserBizPenaltyReq struct {
    PaymentType uint `validate:"required" json:"paymentType" v:"required|in:1,2"` // 支付类型 1 支付宝 2 微信
}

// UserBizPenaltyRep 个签骑手支付违约响应数据
type UserBizPenaltyRep struct {
    OrderId      uint64 `validate:"required" json:"orderId"`      // 订单ID
    PayOrderInfo string `validate:"required" json:"PayOrderInfo"` // 发起支付使用数据
}

// UserBizPenaltyProfileRep 个签骑手获取违约金额响应数据
type UserBizPenaltyProfileRep struct {
    Amount    float64     `validate:"required" json:"amount"`    // 金额
    ComboName string      `validate:"required" json:"comboName"` // 套餐名称
    Days      int64       `validate:"required" json:"days"`      // 逾期天数
    StartAt   *gtime.Time `validate:"required" json:"startAt"`   // 开始时间
    EndAt     *gtime.Time `validate:"required" json:"endAt"`     // 结束时间
}

// UserBizGroupNewReq 团签骑手新领电池
type UserBizGroupNewReq struct {
    BatteryType string `validate:"required" json:"batteryType" v:"required|in:60,72"` // 电池类型 60 / 72
}

type BizEntity struct {
    gmeta.Meta `json:"-" orm:"table:user_biz" swaggerignore:"true"`

    Id          uint64      `orm:"id,primary"   json:"id"`          // ID
    UserId      uint64      `orm:"userId"       json:"userId"`      // 用户ID
    ComboId     uint        `orm:"comboId"      json:"comboId"`     // 套餐ID
    CityId      uint        `orm:"cityId"       json:"cityId"`      // 城市ID
    ShopId      uint        `orm:"shopId"       json:"shopId"`      // 门店ID
    GroupId     uint        `orm:"groupId"     json:"groupId"`      // 团签ID
    BizType     uint        `orm:"bizType"      json:"bizType"`     // 业务类型: 1新签 2换电 3寄存 4恢复计费 5退租
    BatteryType string      `orm:"batteryType"  json:"batteryType"` // 电池型号
    CreatedAt   *gtime.Time `orm:"createdAt"    json:"createdAt"`   // 扫码时间

    Mobile    string `json:"mobile"`    // 手机号
    RealName  string `json:"realName"`  // 姓名
    CityName  string `json:"cityName"`  // 城市
    ShopName  string `json:"shopName"`  // 门店名称
    GroupName string `json:"groupName"` // 团签名称
    ComboName string `json:"comboName"` // 套餐名称

    User        *User      `json:"-" orm:"with:id=UserId"`
    City        *Districts `json:"-" orm:"with:id=CityId"`
    Shop        *Shop      `json:"-" orm:"with:id=ShopId"`
    ComboDetail *Combo     `json:"-" orm:"with:id=ComboId"`
    Group       *Group     `json:"-" orm:"with:id=GroupId"`
}

// BizShopFilterReq 门店获取业务记录请求
type BizShopFilterReq struct {
    Page
    RealName string `json:"realName"`            // 骑手姓名
    Month    string `json:"month"`               // 月份, 如:2021-10-01
    UserType uint   `json:"userTpe" v:"in:1,2"`  // 用户类型: 1个签 2团签
    BizType  uint   `json:"bizTpe" v:"in:3,2,5"` // 业务类型: 1新签 2换电 3寄存 4恢复计费 5退租
    ShopId   uint   `json:"shopId"`              // 门店ID
}

// BizShopFilterResp 门店业务记录返回
type BizShopFilterResp struct {
    Month string               `json:"month"` // 月份
    Total int                  `json:"total"` // 记录数
    Items []*BizShopFilterItem `json:"items"` // 记录详细
}

// BizShopFilterItem 门店业务记录列表数据
type BizShopFilterItem struct {
    RealName  string      `json:"realName"`            // 用户姓名
    ComboName string      `json:"comboName,omitempty"` // 套餐名称
    GroupName string      `json:"groupName,omitempty"` // 团体名称， 名称为空即为 个签用户
    Mobile    string      `json:"mobile"`              // 手机号
    BizType   uint        `json:"bizType"`             // 业务类型: 1新签 2换电 3寄存 4恢复计费 5退租
    CreatedAt *gtime.Time `json:"createdAt"`           // 时间
}

// BizListReq 业务记录请求
type BizListReq struct {
    Page
    UserId    uint        `json:"userId"`                 // 骑手ID
    ShopId    uint        `json:"shopId"`                 // 店铺ID
    RealName  string      `json:"realName"`               // 骑手姓名
    Mobile    string      `json:"mobile" v:"phone-loose"` // 手机号
    StartDate *gtime.Time `json:"startDate"`              // 开始日期 eg: 2021-10-17
    EndDate   *gtime.Time `json:"endDate"`                // 结束日期 eg: 2021-10-19
    UserType  uint        `json:"userTpe" v:"in:1,2"`     // 用户类型: 1个签 2团签
    BizType   uint        `json:"bizTpe" v:"in:3,2,5"`    // 业务类型: 1新签 2换电 3寄存 4恢复计费 5退租
}

// BizSimpleItem 业务简单返回
type BizSimpleItem struct {
    gmeta.Meta `json:"-" orm:"table:user_biz" swaggerignore:"true"`

    Id          uint        `json:"id"`
    ShopId      uint        `json:"shopId"`      // 门店ID
    Type        uint        `json:"type"`        // 业务类型: 1新签 2换电 3寄存 4恢复计费 5退租
    BatteryType string      `json:"batteryType"` // 电池型号: 60 / 72
    CreatedAt   *gtime.Time `json:"createdAt"`   // 业务办理时间

    ShopName string `json:"shopName"` // 门店名称

    Shop *Shop `json:"-" orm:"with:id=shopId"`
}
