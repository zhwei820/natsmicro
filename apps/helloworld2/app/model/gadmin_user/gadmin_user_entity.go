// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package gadmin_user

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table gadmin_user.
type Entity struct {
	Id           int         `orm:"id,primary"      json:"id"`           //
	Status       int         `orm:"status"          json:"status"`       //
	Username     string      `orm:"username,unique" json:"username"`     //
	Nickname     string      `orm:"nickname"        json:"nickname"`     //
	Password     string      `orm:"password"        json:"password"`     //
	Email        string      `orm:"email"           json:"email"`        //
	Phone        string      `orm:"phone"           json:"phone"`        //
	CreateTime   *gtime.Time `orm:"create_time"     json:"create_time"`  //
	UpdateTime   *gtime.Time `orm:"update_time"     json:"update_time"`  //
	AddUserId    int         `orm:"add_user_id"     json:"add_user_id"`  //
	Introduction string      `orm:"introduction"    json:"introduction"` //
	Avatar       string      `orm:"avatar"          json:"avatar"`       //
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}
