// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package refund

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table refund.
type Entity struct {
    Id               uint        `orm:"id,primary"              json:"id"`                 //                          
    UserId           uint64      `orm:"userId"                  json:"user_id"`            //                          
    State            uint        `orm:"state"                   json:"state"`              // 状态: 0处理中 1处理完成  
    No               string      `orm:"no,unique"               json:"no"`                 // 退款编号                 
    RelationType     uint        `orm:"relationType"            json:"relation_type"`      // 关联类型: 1套餐订单      
    RelationId       uint64      `orm:"relationId"              json:"relation_id"`        // 关联ID                   
    Reason           string      `orm:"reason"                  json:"reason"`             // 退款原因                 
    Amount           float64     `orm:"amount"                  json:"amount"`             // 退款金额                 
    PlatformRefundNo string      `orm:"platformRefundNo,unique" json:"platform_refund_no"` // 第三方流水号             
    CreatedAt        *gtime.Time `orm:"createdAt"               json:"created_at"`         //                          
    UpdatedAt        *gtime.Time `orm:"updatedAt"               json:"updated_at"`         //                          
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