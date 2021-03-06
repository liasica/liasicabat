// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package combo_order

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table combo_order.
type Entity struct {
    Id            uint64      `orm:"id,primary"    json:"id"`              //                                    
    ParentId      uint64      `orm:"parentId"      json:"parent_id"`       // 关联订单ID(续签或违约)             
    CityId        uint        `orm:"cityId"        json:"city_id"`         // 城市ID                             
    ShopId        uint        `orm:"shopId"        json:"shop_id"`         // 门店ID                             
    UserId        uint64      `orm:"userId"        json:"user_id"`         // 用户ID                             
    ComboId       uint        `orm:"comboId"       json:"combo_id"`        // 套餐ID                             
    GroupId       uint        `orm:"groupId"       json:"group_id"`        // 团签ID                             
    No            string      `orm:"no,unique"     json:"no"`              // 订单编号                           
    Type          uint        `orm:"type"          json:"type"`            // 订单类别: 1新签 2续费 3违约金      
    Amount        float64     `orm:"amount"        json:"amount"`          // 总金额(包含押金)                   
    Deposit       float64     `orm:"deposit"       json:"deposit"`         // 押金                               
    PayType       uint        `orm:"payType"       json:"pay_type"`        // 支付方式: 1支付宝 2微信            
    PayPlatformNo string      `orm:"payPlatformNo" json:"pay_platform_no"` // 支付平台单号                       
    PayAt         *gtime.Time `orm:"payAt"         json:"pay_at"`          // 支付时间                           
    PayState      uint        `orm:"payState"      json:"pay_state"`       // 支付状态:0未支付  1待支付 2已支付  
    FirstUseAt    *gtime.Time `orm:"firstUseAt"    json:"first_use_at"`    // 启用时间, 即首次领取电池时间       
    Month         int         `orm:"month"         json:"month"`           //                                    
    CreatedAt     *gtime.Time `orm:"createdAt"     json:"created_at"`      //                                    
    UpdatedAt     *gtime.Time `orm:"updatedAt"     json:"updated_at"`      //                                    
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