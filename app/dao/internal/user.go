// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// UserDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserDao struct {
	gmvc.M              // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB      // DB is the raw underlying database management object.
	Table   string      // Table is the table name of the DAO.
	Columns userColumns // Columns contains all the columns of Table that for convenient usage.
}

// UserColumns defines and stores column names for table user.
type userColumns struct {
	Id                           string //
	GroupId                      string // 团签ID
	ComboId                      string // 套餐ID
	ComboOrderId                 string // 套餐订单ID
	Mobile                       string // 手机号码
	Type                         string // 用户类型: 1个签骑手 2团签骑手 3团签BOSS
	Qr                           string // 骑手二维码数据
	RealName                     string // 真实姓名
	IdCardNo                     string // 身份证号码
	IdCardImg1                   string // 身份证人像面
	IdCardImg2                   string // 身份证国徽面
	IdCardImg3                   string // 手持身份证
	AuthState                    string // 实名认证状态: 0未提交 1待审核 2审核通过 3审核拒绝
	BatteryState                 string // 个签换电状态: 0 未开通 1租借中 2寄存中 3已退租
	BatteryType                  string // 电池型号
	BatteryReturnAt              string // 个签应归还电池时间, 小于当前时间即逾期
	BatterySaveAt                string // 个签用户电池寄存时间
	DeviceType                   string // 设备类型: 0未上报 1安卓 2iOS
	DeviceToken                  string // 用户推送token
	EsignAccountId               string // 易签账户ID
	BizBatteryRenewalCnt         string // 积累换次数
	BizBatteryRenewalDays        string // 累计换电自然天数
	BizBatteryRenewalDaysStartAt string // 需要统计使用天数的开始时间, 为空则无需计算
	Salt                         string //
	AccessToken                  string //
	CreatedAt                    string //
	UpdatedAt                    string //
}

func NewUserDao() *UserDao {
	return &UserDao{
		M:     g.DB("default").Model("user").Safe(),
		DB:    g.DB("default"),
		Table: "user",
		Columns: userColumns{
			Id:                           "id",
			GroupId:                      "groupId",
			ComboId:                      "comboId",
			ComboOrderId:                 "comboOrderId",
			Mobile:                       "mobile",
			Type:                         "type",
			Qr:                           "qr",
			RealName:                     "realName",
			IdCardNo:                     "idCardNo",
			IdCardImg1:                   "idCardImg1",
			IdCardImg2:                   "idCardImg2",
			IdCardImg3:                   "idCardImg3",
			AuthState:                    "authState",
			BatteryState:                 "batteryState",
			BatteryType:                  "batteryType",
			BatteryReturnAt:              "batteryReturnAt",
			BatterySaveAt:                "batterySaveAt",
			DeviceType:                   "deviceType",
			DeviceToken:                  "deviceToken",
			EsignAccountId:               "esignAccountId",
			BizBatteryRenewalCnt:         "bizBatteryRenewalCnt",
			BizBatteryRenewalDays:        "bizBatteryRenewalDays",
			BizBatteryRenewalDaysStartAt: "bizBatteryRenewalDaysStartAt",
			Salt:                         "salt",
			AccessToken:                  "accessToken",
			CreatedAt:                    "createdAt",
			UpdatedAt:                    "updatedAt",
		},
	}
}
