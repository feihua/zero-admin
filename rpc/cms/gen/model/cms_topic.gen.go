// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCmsTopic = "cms_topic"

// CmsTopic 话题表
type CmsTopic struct {
	ID             int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`                        // 主键ID
	CategoryID     int64      `gorm:"column:category_id;not null;comment:关联分类id" json:"category_id"`                         // 关联分类id
	Name           string     `gorm:"column:name;not null;comment:话题名称" json:"name"`                                         // 话题名称
	StartTime      time.Time  `gorm:"column:start_time;not null;comment:话题开始时间" json:"start_time"`                           // 话题开始时间
	EndTime        time.Time  `gorm:"column:end_time;not null;comment:话题结束时间" json:"end_time"`                               // 话题结束时间
	AttendCount    int32      `gorm:"column:attend_count;not null;comment:参与人数" json:"attend_count"`                         // 参与人数
	AttentionCount int32      `gorm:"column:attention_count;not null;comment:关注人数" json:"attention_count"`                   // 关注人数
	ReadCount      int32      `gorm:"column:read_count;not null;comment:阅读数" json:"read_count"`                              // 阅读数
	AwardName      string     `gorm:"column:award_name;not null;comment:奖品名称" json:"award_name"`                             // 奖品名称
	AttendType     string     `gorm:"column:attend_type;not null;comment:参与方式" json:"attend_type"`                           // 参与方式
	Content        string     `gorm:"column:content;not null;comment:话题内容" json:"content"`                                   // 话题内容
	CreateBy       string     `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`                                // 创建者
	CreateTime     time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy       string     `gorm:"column:update_by;not null;comment:更新者" json:"update_by"`                                // 更新者
	UpdateTime     *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName CmsTopic's table name
func (*CmsTopic) TableName() string {
	return TableNameCmsTopic
}
