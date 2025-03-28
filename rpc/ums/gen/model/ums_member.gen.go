// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUmsMember = "ums_member"

// UmsMember 会员信息
type UmsMember struct {
	ID                    int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	MemberLevelID         int64      `gorm:"column:member_level_id;not null;comment:会员等级id" json:"member_level_id"`                 // 会员等级id
	MemberName            string     `gorm:"column:member_name;not null;comment:用户名" json:"member_name"`                            // 用户名
	Password              string     `gorm:"column:password;not null;comment:密码" json:"password"`                                   // 密码
	Nickname              string     `gorm:"column:nickname;not null;comment:昵称" json:"nickname"`                                   // 昵称
	Phone                 string     `gorm:"column:phone;not null;comment:手机号码" json:"phone"`                                       // 手机号码
	MemberStatus          int32      `gorm:"column:member_status;not null;comment:帐号启用状态:0->禁用；1->启用" json:"member_status"`         // 帐号启用状态:0->禁用；1->启用
	Icon                  string     `gorm:"column:icon;not null;comment:头像" json:"icon"`                                           // 头像
	Gender                int32      `gorm:"column:gender;not null;comment:性别：0->未知；1->男；2->女" json:"gender"`                       // 性别：0->未知；1->男；2->女
	Birthday              *time.Time `gorm:"column:birthday;comment:生日" json:"birthday"`                                            // 生日
	City                  string     `gorm:"column:city;not null;comment:所做城市" json:"city"`                                         // 所做城市
	Job                   string     `gorm:"column:job;not null;comment:职业" json:"job"`                                             // 职业
	PersonalizedSignature string     `gorm:"column:personalized_signature;not null;comment:个性签名" json:"personalized_signature"`     // 个性签名
	SourceType            int32      `gorm:"column:source_type;not null;comment:用户来源：1->移动端; 2->pc端" json:"source_type"`            // 用户来源：1->移动端; 2->pc端
	Integration           int32      `gorm:"column:integration;not null;comment:积分" json:"integration"`                             // 积分
	Growth                int32      `gorm:"column:growth;not null;comment:成长值" json:"growth"`                                      // 成长值
	LotteryCount          int32      `gorm:"column:lottery_count;not null;comment:剩余抽奖次数" json:"lottery_count"`                     // 剩余抽奖次数
	HistoryIntegration    int32      `gorm:"column:history_integration;not null;comment:历史积分数量" json:"history_integration"`         // 历史积分数量
	CreateTime            time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime            *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName UmsMember's table name
func (*UmsMember) TableName() string {
	return TableNameUmsMember
}
