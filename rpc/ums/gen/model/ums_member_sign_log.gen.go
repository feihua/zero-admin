// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUmsMemberSignLog = "ums_member_sign_log"

// UmsMemberSignLog 会员签到记录表
type UmsMemberSignLog struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MemberID     int64     `gorm:"column:member_id;not null;comment:会员ID" json:"member_id"`                     // 会员ID
	SignDate     time.Time `gorm:"column:sign_date;not null;comment:签到日期" json:"sign_date"`                     // 签到日期
	ContinueDays int32     `gorm:"column:continue_days;not null;default:1;comment:连续签到天数" json:"continue_days"` // 连续签到天数
	Points       int32     `gorm:"column:points;not null;comment:获得积分" json:"points"`                           // 获得积分
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
}

// TableName UmsMemberSignLog's table name
func (*UmsMemberSignLog) TableName() string {
	return TableNameUmsMemberSignLog
}
