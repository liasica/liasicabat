// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"battery/app/dao/internal"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// sysUsersDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type sysUsersDao struct {
	*internal.SysUsersDao
}

var (
	// SysUsers is globally public accessible object for table sys_users operations.
	SysUsers sysUsersDao
)

func init() {
	SysUsers = sysUsersDao{
		internal.NewSysUsersDao(),
	}
}

// GenerateAccessToken 生成accessToken
func (d *sysUsersDao) GenerateAccessToken(id uint, salt string) string {
	m := md5.New()
	m.Write([]byte(strconv.Itoa(int(id))))
	m.Write([]byte(time.Now().String()))
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}

// GenerateSalt 生成加密盐
func (d *sysUsersDao) GenerateSalt() string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	randBytes := make([]byte, 6)
	for i := 0; i < 6; i++ {
		randBytes[i] = byte(r.Intn(26) + 65)
	}
	return string(randBytes)
}

// EncryptPassword 加密密码
func (d *sysUsersDao) EncryptPassword(password, salt string) string {
	m := md5.New()
	if len(password) != 32 {
		mp := md5.New()
		mp.Write([]byte(password))
		m.Write([]byte(hex.EncodeToString(mp.Sum(nil))))
	} else {
		m.Write([]byte(password))
	}
	m.Write([]byte(salt))
	return hex.EncodeToString(m.Sum(nil))
}
