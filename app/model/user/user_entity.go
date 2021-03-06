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
    GroupId                      uint        `orm:"groupId"                      json:"group_id"`                          // 团签ID                                             
    ComboId                      uint        `orm:"comboId"                      json:"combo_id"`                          // 套餐ID                                             
    ComboOrderId                 uint64      `orm:"comboOrderId"                 json:"combo_order_id"`                    // 套餐订单ID                                         
    Mobile                       string      `orm:"mobile,unique"                json:"mobile"`                            // 手机号码                                           
    Type                         uint        `orm:"type"                         json:"type"`                              // 用户类型: 1个签骑手 2团签骑手 3团签BOSS            
    Qr                           string      `orm:"qr,unique"                    json:"qr"`                                // 骑手二维码数据                                     
    RealName                     string      `orm:"realName"                     json:"real_name"`                         // 真实姓名                                           
    IdCardNo                     string      `orm:"idCardNo"                     json:"id_card_no"`                        // 身份证号码                                         
    IdCardImg1                   string      `orm:"idCardImg1"                   json:"id_card_img_1"`                     // 身份证人像面                                       
    IdCardImg2                   string      `orm:"idCardImg2"                   json:"id_card_img_2"`                     // 身份证国徽面                                       
    IdCardImg3                   string      `orm:"idCardImg3"                   json:"id_card_img_3"`                     // 手持身份证                                         
    AuthState                    uint        `orm:"authState"                    json:"auth_state"`                        // 实名认证状态: 0未提交 1待审核 2审核通过 3审核拒绝  
    BatteryState                 uint        `orm:"batteryState"                 json:"battery_state"`                     // 个签换电状态: 0 未开通 1租借中 2寄存中 3已退租     
    BatteryType                  string      `orm:"batteryType"                  json:"battery_type"`                      // 电池型号                                           
    BatteryReturnAt              *gtime.Time `orm:"batteryReturnAt"              json:"battery_return_at"`                 // 个签应归还电池时间, 小于当前时间即逾期             
    BatterySaveAt                *gtime.Time `orm:"batterySaveAt"                json:"battery_save_at"`                   // 个签用户电池寄存时间                               
    DeviceType                   int         `orm:"deviceType"                   json:"device_type"`                       // 设备类型: 0未上报 1安卓 2iOS                       
    DeviceToken                  string      `orm:"deviceToken"                  json:"device_token"`                      // 用户推送token                                      
    EsignAccountId               string      `orm:"esignAccountId"               json:"esign_account_id"`                  // 易签账户ID                                         
    BizBatteryRenewalCnt         uint        `orm:"bizBatteryRenewalCnt"         json:"biz_battery_renewal_cnt"`           // 积累换次数                                         
    BizBatteryRenewalDays        uint        `orm:"bizBatteryRenewalDays"        json:"biz_battery_renewal_days"`          // 累计换电自然天数                                   
    BizBatteryRenewalDaysStartAt *gtime.Time `orm:"bizBatteryRenewalDaysStartAt" json:"biz_battery_renewal_days_start_at"` // 需要统计使用天数的开始时间, 为空则无需计算         
    Salt                         string      `orm:"salt"                         json:"salt"`                              //                                                    
    AccessToken                  string      `orm:"accessToken,unique"           json:"access_token"`                      //                                                    
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