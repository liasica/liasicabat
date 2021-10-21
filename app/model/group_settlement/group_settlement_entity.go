// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package group_settlement

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table group_settlement.
type Entity struct {
    Id        uint64      `orm:"id,primary" json:"id"`          //             
    Hash      string      `orm:"hash"       json:"hash"`        // 结算单hash  
    GroupId   uint        `orm:"groupId"    json:"group_id"`    // 团队ID      
    Amount    float64     `orm:"amount"     json:"amount"`      // 总结算金额  
    SysUserId uint        `orm:"sysUserId"  json:"sys_user_id"` // 结算人ID    
    SysName   string      `orm:"sysName"    json:"sys_name"`    // 结算人姓名  
    Date      *gtime.Time `orm:"date"       json:"date"`        // 结算日      
    Remark    string      `orm:"remark"     json:"remark"`      // 备注        
    Detail    string      `orm:"detail"     json:"detail"`      // 结算详情    
    CreatedAt *gtime.Time `orm:"createdAt"  json:"created_at"`  //             
    UpdatedAt *gtime.Time `orm:"updatedAt"  json:"updated_at"`  //             
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