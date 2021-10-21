// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id                           uint64      `orm:"id,primary"                   json:"id"`                           //
	GroupId                      uint        `orm:"groupId"                      json:"groupId"`                      // 团签ID
	ComboId                      uint        `orm:"comboId"                      json:"comboId"`                      // 套餐ID
	ComboOrderId                 uint64      `orm:"comboOrderId"                 json:"comboOrderId"`                 // 套餐订单ID
	SettlementDetailId           uint64      `orm:"settlementDetailId"           json:"settlementDetailId"`           // 团签当前账单ID
	Mobile                       string      `orm:"mobile,unique"                json:"mobile"`                       // 手机号码
	Type                         uint        `orm:"type"                         json:"type"`                         // 用户类型: 1个签骑手 2团签骑手 3团签BOSS
	Qr                           string      `orm:"qr,unique"                    json:"qr"`                           // 骑手二维码数据
	RealName                     string      `orm:"realName"                     json:"realName"`                     // 真实姓名
	IdCardNo                     string      `orm:"idCardNo"                     json:"idCardNo"`                     // 身份证号码
	IdCardImg1                   string      `orm:"idCardImg1"                   json:"idCardImg1"`                   // 身份证人像面
	IdCardImg2                   string      `orm:"idCardImg2"                   json:"idCardImg2"`                   // 身份证国徽面
	IdCardImg3                   string      `orm:"idCardImg3"                   json:"idCardImg3"`                   // 手持身份证
	AuthState                    uint        `orm:"authState"                    json:"authState"`                    // 实名认证状态: 0未提交 1待审核 2审核通过 3审核拒绝
	BatteryState                 uint        `orm:"batteryState"                 json:"batteryState"`                 // 个签换电状态: 0 未开通 1租借中 2寄存中 3已退租
	BatteryType                  string      `orm:"batteryType"                  json:"batteryType"`                  // 电池型号
	BatteryReturnAt              *gtime.Time `orm:"batteryReturnAt"              json:"batteryReturnAt"`              // 个签应归还电池时间, 小于当前时间即逾期
	BatterySaveAt                *gtime.Time `orm:"batterySaveAt"                json:"batterySaveAt"`                // 个签用户电池寄存时间
	DeviceType                   int         `orm:"deviceType"                   json:"deviceType"`                   // 设备类型: 0未上报 1安卓 2iOS
	DeviceToken                  string      `orm:"deviceToken"                  json:"deviceToken"`                  // 用户推送token
	EsignAccountId               string      `orm:"esignAccountId"               json:"esignAccountId"`               // 易签账户ID
	BizBatteryRenewalCnt         uint        `orm:"bizBatteryRenewalCnt"         json:"bizBatteryRenewalCnt"`         // 积累换次数
	BizBatteryRenewalDays        uint        `orm:"bizBatteryRenewalDays"        json:"bizBatteryRenewalDays"`        // 累计换电自然天数
	BizBatteryRenewalDaysStartAt *gtime.Time `orm:"bizBatteryRenewalDaysStartAt" json:"bizBatteryRenewalDaysStartAt"` // 需要统计使用天数的开始时间, 为空则无需计算
	Salt                         string      `orm:"salt"                         json:"salt"`                         //
	AccessToken                  string      `orm:"accessToken,unique"           json:"accessToken"`                  //
	CreatedAt                    *gtime.Time `orm:"createdAt"                    json:"createdAt"`                    //
	UpdatedAt                    *gtime.Time `orm:"updatedAt"                    json:"updatedAt"`                    //
}
