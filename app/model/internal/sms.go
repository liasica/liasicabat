// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Sms is the golang structure for table sms.
type Sms struct {
	Id        int64       `orm:"id,primary" json:"id"`        //
	Mobile    string      `orm:"mobile"     json:"mobile"`    //
	Code      string      `orm:"code"       json:"code"`      // 验证码
	CreatedAt *gtime.Time `orm:"createdAt"  json:"createdAt"` //
}
