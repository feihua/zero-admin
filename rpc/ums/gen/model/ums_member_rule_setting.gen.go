// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUmsMemberRuleSetting = "ums_member_rule_setting"

// UmsMemberRuleSetting 会员积分成长规则表
type UmsMemberRuleSetting struct {
	ID               int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ConsumePerPoint  int64      `gorm:"column:consume_per_point;not null;comment:每消费多少元获取1个点" json:"consume_per_point"`        // 每消费多少元获取1个点
	LowOrderAmount   int64      `gorm:"column:low_order_amount;not null;comment:最低获取点数的订单金额" json:"low_order_amount"`          // 最低获取点数的订单金额
	MaxPointPerOrder int32      `gorm:"column:max_point_per_order;not null;comment:每笔订单最高获取点数" json:"max_point_per_order"`     // 每笔订单最高获取点数
	RuleType         int32      `gorm:"column:rule_type;not null;comment:类型：0->积分规则；1->成长值规则" json:"rule_type"`                // 类型：0->积分规则；1->成长值规则
	Status           int32      `gorm:"column:status;not null;default:1;comment:状态：0->禁用；1->启用" json:"status"`                 // 状态：0->禁用；1->启用
	CreateBy         int64      `gorm:"column:create_by;not null;comment:创建人ID" json:"create_by"`                              // 创建人ID
	CreateTime       time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy         *int64     `gorm:"column:update_by;comment:更新人ID" json:"update_by"`                                       // 更新人ID
	UpdateTime       *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName UmsMemberRuleSetting's table name
func (*UmsMemberRuleSetting) TableName() string {
	return TableNameUmsMemberRuleSetting
}
