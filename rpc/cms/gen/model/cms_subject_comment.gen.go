// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCmsSubjectComment = "cms_subject_comment"

// CmsSubjectComment 专题评论表
type CmsSubjectComment struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`             // 编号
	SubjectID      int64     `gorm:"column:subject_id;not null;comment:关联专题id" json:"subject_id"`              // 关联专题id
	MemberNickName string    `gorm:"column:member_nick_name;not null;comment:关联会员昵称" json:"member_nick_name"`  // 关联会员昵称
	MemberIcon     string    `gorm:"column:member_icon;not null;comment:会员头像" json:"member_icon"`              // 会员头像
	Content        string    `gorm:"column:content;not null;comment:评论内容" json:"content"`                      // 评论内容
	CreateTime     time.Time `gorm:"column:create_time;not null;comment:创建时间" json:"create_time"`              // 创建时间
	ShowStatus     int32     `gorm:"column:show_status;not null;comment:是否显示，0->不显示；1->显示" json:"show_status"` // 是否显示，0->不显示；1->显示
}

// TableName CmsSubjectComment's table name
func (*CmsSubjectComment) TableName() string {
	return TableNameCmsSubjectComment
}
