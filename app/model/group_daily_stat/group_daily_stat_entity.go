// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package group_daily_stat

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table group_daily_stat.
type Entity struct {
    Id          uint64      `orm:"id,primary"  json:"id"`           //                         
    GroupId     uint        `orm:"groupId"     json:"group_id"`     //                         
    BatteryType uint        `orm:"batteryType" json:"battery_type"` // 电池型号 60 / 72        
    IsArrears   uint        `orm:"isArrears"   json:"is_arrears"`   // 是否未付款 1 是 0 不是  
    Date        uint        `orm:"date"        json:"date"`         // 日期 如 20210705        
    Total       uint        `orm:"total"       json:"total"`        // 使用人数                
    CreatedAt   *gtime.Time `orm:"createdAt"   json:"created_at"`   // 创建时间                
    UpdatedAt   *gtime.Time `orm:"updatedAt"   json:"updated_at"`   //                         
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