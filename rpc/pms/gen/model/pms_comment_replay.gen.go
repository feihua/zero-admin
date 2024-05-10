// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePmsCommentReplay = "pms_comment_replay"

// PmsCommentReplay 产品评价回复表
type PmsCommentReplay struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CommentID      int64     `gorm:"column:comment_id;not null" json:"comment_id"`
	MemberNickName string    `gorm:"column:member_nick_name;not null" json:"member_nick_name"`
	MemberIcon     string    `gorm:"column:member_icon;not null" json:"member_icon"`
	Content        string    `gorm:"column:content;not null" json:"content"`
	CreateTime     time.Time `gorm:"column:create_time;not null" json:"create_time"`
	Type           int32     `gorm:"column:type;not null;comment:评论人员类型；0->会员；1->管理员" json:"type"` // 评论人员类型；0->会员；1->管理员
}

// TableName PmsCommentReplay's table name
func (*PmsCommentReplay) TableName() string {
	return TableNamePmsCommentReplay
}