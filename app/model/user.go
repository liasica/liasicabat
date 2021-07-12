package model

import (
	"battery/app/model/internal"
	"github.com/gogf/gf/os/gtime"
)

const (
	//实名认证状态
	AuthStateVerifyDefault = 0 //未提交
	AuthStateVerifyWait    = 1 //待审核
	AuthStateVerifySuccess = 2 //审核通过
	AuthStateDefaultFailed = 3 //审核失败

	//换电状态
	BatteryStateDefault = 0 //未开通
	BatteryStateNew     = 1 //新购待领取 （团用户未领取）
	BatteryStateUse     = 2 //租借中
	BatteryStateSave    = 3 //寄存中
	BatteryStateExit    = 4 //已退租
	BatteryStateOverdue = 5 //已逾期

	//签约状态
	SignStateDefault = 0 //未签约
	SignStateDone    = 1 //已签约

	UserTypePersonal   = 1 //个签用户
	UserTypeGroupRider = 2 //团签骑手
	UserTypeGroupBoss  = 3 //团签BOSS
)

type User internal.User

// UserRegisterReq 用户注册请求数据
type UserRegisterReq struct {
	Mobile string `validate:"required" v:"required|phone-loose"` //手机号
	Sms    string `validate:"required" v:"required|length:6,6"`  //短信验证码
}

// UserLoginReq 用户登录请求数据
type UserLoginReq struct {
	Mobile string `validate:"required" v:"required|phone-loose"` //手机号
	Sms    string `validate:"required" v:"required|length:6,6"`  //短信验证码
}

// UserLoginRep 用户登录返回数据
type UserLoginRep struct {
	AccessToken string `validate:"required" json:"accessToken"` //请求 token
	Type        uint   `validate:"required" json:"type"`        //用户角色 1 个签骑手 2 团签骑手 3 团签BOSS
	AuthState   uint   `validate:"required" json:"authState"`   //实名认证状态 0 未提交 ，1 待审核， 2 审核通过，3 审核未通过
	SignState   uint   `validate:"required" json:"signState"`   //签约状态 0 未签约 1 已签约
}

// UserRealNameAuthReq 实名认证请求数据
type UserRealNameAuthReq struct {
	RealName   string `validate:"required" v:"required|length:2,10"`              //真实姓名
	IdCardNo   string `validate:"required" v:"required|length:15,18|resident-id"` //身份证号码
	IdCardImg1 string `validate:"required" v:"required"`                          //身份证正面照片
	IdCardImg2 string `validate:"required" v:"required"`                          //身份证反面照片
	IdCardImg3 string `validate:"required" v:"required"`                          //身份证手持照片
}

// UserRealNameAuthRep 实名认证响应数据
type UserRealNameAuthRep struct {
	FlowId    string `json:"flowId"`       //流程ID
	ShortLink string `json:"shortLink"`    //短地址
	Url       string `validate:"required"` // 骑手实名认证连接地址
}

// UserSignRep 获取签约URL
type UserSignRep struct {
	Url      string `json:"url"`
	ShortUrl string `json:"shortUrl"`
}

// RealNameAuthVerifyProfileRep 获取用户实名认证提交资料信息
type RealNameAuthVerifyProfileRep struct {
	Id         uint64 `orm:"id,primary"         json:"id"`         //
	Mobile     string `orm:"mobile,unique"      json:"mobile"`     //
	RealName   string `orm:"realName"           json:"realName"`   // 真实姓名
	IdCardNo   string `orm:"idCardNo"           json:"idCardNo"`   // 身份证号码
	IdCardImg1 string `orm:"idCardImg1"         json:"idCardImg1"` // 正面图
	IdCardImg2 string `orm:"idCardImg2"         json:"idCardImg2"` // 反面图
	IdCardImg3 string `orm:"idCardImg3"         json:"idCardImg3"` // 人像图
	AuthState  uint   `orm:"authState"          json:"authState"`  // 实名认证状态 0 未提交 ，1 待审核， 2 审核通过，3 审核未通过
}

// RealNameAuthVerifyReq 实名认证审核请求数据
type RealNameAuthVerifyReq struct {
	AuthState uint `v:"required|in:2,3"` //审核结果 2 通过 3 失败
}

// BizProfileRep 用户业务办理店长扫码获取用户信息
type BizProfileRep struct {
	Id           uint64 `json:"id"`
	Mobile       string `json:"mobile"`
	RealName     string `json:"realName"`               //真实姓名
	IdCardNo     string `json:"idCardNo"`               //身份证号码
	AuthState    uint   `json:"authState"`              //实名认证状态 0 未提交 ，1 待审核， 2 审核通过，3 审核未通过
	BatteryState uint   `json:"batteryState"`           //换电状态：0 未开通，1 租借中，2 寄存中，3 已退租 4 已逾期
	BatteryType  uint   `json:"batteryType"`            //电池类型 60 / 72
	PackagesName string `json:"PackagesName,omitempty"` //套餐名称
	GroupId      uint   `json:"groupId"`                //团体Id，个签用户为 0
	GroupName    string `json:"groupName,omitempty"`    //团体名称
}

// PushTokenReq 上报用户的推送token
type PushTokenReq struct {
	DeviceType  int    `validate:"required" v:"required|in:1,2" json:"deviceType"` // 1 android  2 ios
	DeviceToken string `validate:"required" v:"required" json:"deviceToken"`       //token
}

// UserProfileRep 骑手端用户信息概况
type UserProfileRep struct {
	Name      string `json:"name"`                          //姓名
	Type      uint   `validate:"required" json:"type"`      //用户角色 1 个签骑手 2 团签骑手 3 团签BOSS
	AuthState uint   `validate:"required" json:"authState"` //实名认证状态 0 未提交 ，1 待审核， 2 审核通过，3 审核未通过
	SignState uint   `validate:"required" json:"signState"` //签约状态 0 未签约 1 已签约
	Qr        string `json:"qr"`                            //用户二维码数据，需要本地生成图片

	User struct {
		BatteryState    uint        `json:"batteryState"`    //个签骑手换电状态：0 未开通， 1 新签未领 ，2 租借中，3 寄存中，4 已退租， 5 已逾期
		PackagesId      uint        `json:"packagesId"`      //个签骑手所购套餐ID
		PackagesName    string      `json:"packagesName"`    //个签骑手所购套餐名称
		BatteryReturnAt *gtime.Time `json:"batteryReturnAt"` //个签骑手套餐到期时间
	} `json:"user"` //个签用户套餐信息， 其它类型用户忽略

	GroupUser struct {
		BatteryState uint `json:"batteryState"` //团签骑手换电状态：0 未开通，2 租借中，3 寄存中，4 已退租
		BatteryType  uint `json:"batteryType"`  //电池型号 60 / 72  未开通为 0
	} `json:"groupUser"` // 团签用户骑手信息， 其它类型用户忽略

	GroupBoos UserGroupStatRep `json:"groupBoos"` //团队BOSS信息， 其它类型用户忽略
}
