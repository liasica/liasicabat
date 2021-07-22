// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package user

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table user.
type Entity struct {
    Id                           uint64      `orm:"id,primary"                   json:"id"`                                //                                                              
    GroupId                      uint        `orm:"groupId"                      json:"group_id"`                          // 团签用户，团体ID;  个签用户为 0                              
    Mobile                       string      `orm:"mobile,unique"                json:"mobile"`                            // 手机号码                                                     
    Type                         uint        `orm:"type"                         json:"type"`                              // 用户类型 	1 个签骑手 2 团签骑手 3 团签BOSS                    
    Qr                           string      `orm:"qr"                           json:"qr"`                                // 骑手二维码数据                                               
    RealName                     string      `orm:"realName"                     json:"real_name"`                         // 真实姓名                                                     
    IdCardNo                     string      `orm:"idCardNo"                     json:"id_card_no"`                        // 身份证号码                                                   
    IdCardImg1                   string      `orm:"idCardImg1"                   json:"id_card_img_1"`                     // 正面图                                                       
    IdCardImg2                   string      `orm:"idCardImg2"                   json:"id_card_img_2"`                     // 反面图                                                       
    IdCardImg3                   string      `orm:"idCardImg3"                   json:"id_card_img_3"`                     // 人像图                                                       
    AuthState                    uint        `orm:"authState"                    json:"auth_state"`                        // 实名认证状态 0 未提交 ，1 待审核， 2 审核通过，3 审核未通过  
    BatteryState                 uint        `orm:"batteryState"                 json:"battery_state"`                     // 个人用户换点状态：0 未开通，1 租借中，2 寄存中，3 已退租     
    BatteryType                  uint        `orm:"batteryType"                  json:"battery_type"`                      // 套餐电池型号 60 、 72                                        
    PackagesId                   uint        `orm:"packagesId"                   json:"packages_id"`                       // 套餐ID                                                       
    PackagesOrderId              uint64      `orm:"packagesOrderId"              json:"packages_order_id"`                 // 办理套餐订单ID                                               
    BatteryReturnAt              *gtime.Time `orm:"batteryReturnAt"              json:"battery_return_at"`                 // 个人用户应归还电池时间， 小于当前时间即逾期                  
    BatterySaveAt                *gtime.Time `orm:"batterySaveAt"                json:"battery_save_at"`                   // 个签用户电池寄存时间                                         
    AccessToken                  string      `orm:"accessToken,unique"           json:"access_token"`                      //                                                              
    Salt                         string      `orm:"salt"                         json:"salt"`                              //                                                              
    DeviceType                   int         `orm:"deviceType"                   json:"device_type"`                       // 1 安卓  2  iOS                                               
    DeviceToken                  string      `orm:"deviceToken"                  json:"device_token"`                      // 用户推送消息唯一ID                                           
    EsignAccountId               string      `orm:"esignAccountId"               json:"esign_account_id"`                  // 易签账户ID                                                   
    BizBatteryRenewalCnt         uint        `orm:"bizBatteryRenewalCnt"         json:"biz_battery_renewal_cnt"`           // 积累换次数                                                   
    BizBatteryRenewalDays        uint        `orm:"bizBatteryRenewalDays"        json:"biz_battery_renewal_days"`          // 累计换电自然天数                                             
    BizBatteryRenewalDaysStartAt *gtime.Time `orm:"bizBatteryRenewalDaysStartAt" json:"biz_battery_renewal_days_start_at"` // 需要统计使用天数的开始时间，为空需要统计所有时间             
    CreatedAt                    *gtime.Time `orm:"createdAt"                    json:"created_at"`                        //                                                              
    UpdatedAt                    *gtime.Time `orm:"updatedAt"                    json:"updated_at"`                        //                                                              
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
// Deprecated.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
// Deprecated.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
// Deprecated.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
// Deprecated.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
// Deprecated.
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
// Deprecated.
func (r *Entity) Update() (result sql.Result, err error) {
	where, args, err := gdb.GetWhereConditionOfStruct(r)
	if err != nil {
		return nil, err
	}
	return Model.Data(r).Where(where, args).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
// Deprecated.
func (r *Entity) Delete() (result sql.Result, err error) {
	where, args, err := gdb.GetWhereConditionOfStruct(r)
	if err != nil {
		return nil, err
	}
	return Model.Where(where, args).Delete()
}