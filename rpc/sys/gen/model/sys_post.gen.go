// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysPost = "sys_post"

// SysPost 岗位信息表
type SysPost struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:岗位id" json:"id"`                        // 岗位id
	PostCode   string     `gorm:"column:post_code;not null;comment:岗位编码" json:"post_code"`                               // 岗位编码
	PostName   string     `gorm:"column:post_name;not null;comment:岗位名称" json:"post_name"`                               // 岗位名称
	Sort       int32      `gorm:"column:sort;not null;comment:显示顺序" json:"sort"`                                         // 显示顺序
	Status     int32      `gorm:"column:status;not null;comment:岗位状态（0：停用，1:正常）" json:"status"`                          // 岗位状态（0：停用，1:正常）
	Remark     string     `gorm:"column:remark;not null;comment:备注" json:"remark"`                                       // 备注
	CreateBy   string     `gorm:"column:create_by;not null;default:admin;comment:创建者" json:"create_by"`                  // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by;not null;comment:更新者" json:"update_by"`                                // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName SysPost's table name
func (*SysPost) TableName() string {
	return TableNameSysPost
}
