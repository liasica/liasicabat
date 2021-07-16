package model

import (
	"battery/app/model/internal"
	"github.com/gogf/gf/os/gtime"
)

const (
	UserBizNew            = 1 //新签
	UserBizBatteryRenewal = 2 //换电池
	UserBizBatterySave    = 3 //寄存
	UserBizBatteryUnSave  = 4 //恢复计费
	UserBizClose          = 5 //退租
)

// UserBiz is the golang structure for table user_biz.
type UserBiz internal.UserBiz

// UserBizReq 业务办理请求数据
type UserBizReq struct {
	Code string `validate:"required" v:"required|string"`   //用户二维码code
	Type uint   `validate:"required" v:"required|in:3,4,5"` //业务类型:  3 寄存(仅个签用户使用)， 4 恢复计费，5 退租
}

// UserBizRecordListRep 骑手获取换电记录响应
type UserBizRecordListRep struct {
	ShopName string      `validate:"required" json:"shopName"` //店铺名称
	ScanAt   *gtime.Time `validate:"required" json:"scanAt"`   //扫码时间
	CityName string      `validate:"required" json:"cityName"` //城市名称
}

// UserBizBatteryRenewalReq 骑手扫码换电请求
type UserBizBatteryRenewalReq struct {
	Code string `validate:"required" v:"required" json:"code"` //店铺二维码扫码获得
}

// UserBizBatteryRenewalRep 骑手扫码换电响应
type UserBizBatteryRenewalRep struct {
	ShopName    string      `validate:"required" json:"shopName"`    //店铺名称
	BatteryType uint        `validate:"required" json:"batteryType"` //电池型号
	At          *gtime.Time `validate:"required" json:"at"`          //时间
}

// UserBizRecordStatRep 骑手换电统计
type UserBizRecordStatRep struct {
	Count uint `validate:"required" json:"count"` //累计换电次数
	Days  uint `validate:"required" json:"days"`  //累计使用天数
}

// UserBizSignReq 个签骑手新签套餐请求数据
type UserBizSignReq struct {
	PackagesId uint `validate:"required" json:"packagesId" v:"required|integer"` //套餐ID
}

// UserBizNewReq 个签骑手签约之后获取支付信息
type UserBizNewReq struct {
	PayType uint `validate:"required" json:"payType" v:"required|integer|in:1,2"` //支付方式 1 支付宝 2 微信支付
}

// UserBizNewRep 个签骑手新签套餐响应数据
type UserBizNewRep struct {
	PayOrderInfo string `validate:"required" json:"PayOrderInfo"` //发起支付使用数据
}

// UserBizRenewalReq 个签骑手续约请求数据
type UserBizRenewalReq struct {
	PaymentType uint `validate:"required" json:"paymentType" v:"required|in:1,2"` //支付类型 1 支付宝 2 微信
}

// UserBizRenewalRep 个签骑手续约响应数据
type UserBizRenewalRep struct {
	PayOrderInfo string `validate:"required" json:"PayOrderInfo"` //发起支付使用数据
}

// UserBizPenaltyReq 个签骑手支付违约金请求数据
type UserBizPenaltyReq struct {
	PaymentType uint `validate:"required" json:"paymentType" v:"required|in:1,2"` //支付类型 1 支付宝 2 微信
}

// UserBizPenaltyRep 个签骑手支付违约响应数据
type UserBizPenaltyRep struct {
	PayOrderInfo string `validate:"required" json:"PayOrderInfo"` //发起支付使用数据
}

// UserBizGroupNewReq 团签骑手新领电池
type UserBizGroupNewReq struct {
	BatteryType uint `validate:"required" json:"batteryType" v:"required|in:60,72"` //电池类型 60 / 72
}

// UserBizShopRecordReq 店长获取业务记录请求
type UserBizShopRecordReq struct {
	Page
	Month    uint `validate:"required" json:"batteryType" v:"required"`     //月份数字 如： 20210705
	UserType uint `validate:"required" json:"userTpe" v:"required|in:1,2"`  //用户类型 1 个签 2 团签
	BizType  uint `validate:"required" json:"bizTpe" v:"required|in:3,2,5"` //用户类型 2 换电 3 寄存(仅个签可用)，5 退租
}

// UserBizShopRecordRep 店长获取业务记录响应
type UserBizShopRecordRep struct {
	UserName     string      `json:"userName"`            //用户姓名
	PackagesName string      `json:"packagesName"`        //套餐名称
	GroupName    string      `json:"groupName,omitempty"` //团体名称， 名称为空即为 个签用户
	UserMobile   string      `json:"userMobile"`          //手机号
	BizType      uint        `json:"bizType"`             //业务类型  2 换电 3 寄存(仅个签可用)，5 退租
	At           *gtime.Time `json:"At"`                  //时间
}

// UserBizShopRecordMonthTotalReq 店长获取业务记录按月统计请求
type UserBizShopRecordMonthTotalReq struct {
	Month    uint `validate:"required" json:"batteryType" v:"required"`    // "月份数字，如 202106"
	UserType uint `validate:"required" json:"userTpe" v:"required|in:1,2"` //  "业务类型 1 个签  2 团签"
}

// UserBizShopRecordMonthTotalRep 店长获取业务记录按月统计响应
type UserBizShopRecordMonthTotalRep struct {
	Cnt int `json:"cnt"` //总条数
}
