// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysRole = "sys_role"

// SysRole 角色管理
type SysRole struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:编号" json:"id"`                          // 编号
	Name       string     `gorm:"column:name;not null;comment:角色名称" json:"name"`                                         // 角色名称
	Remark     *string    `gorm:"column:remark;comment:备注" json:"remark"`                                                // 备注
	CreateBy   string     `gorm:"column:create_by;not null;comment:创建者" json:"create_by"`                                // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy   *string    `gorm:"column:update_by;comment:更新者" json:"update_by"`                                         // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;comment:更新时间" json:"update_time"`          // 更新时间
	DelFlag    int32      `gorm:"column:del_flag;not null;default:1;comment:是否删除  0：已删除  1：正常" json:"del_flag"`          // 是否删除  0：已删除  1：正常
	Status     int32      `gorm:"column:status;not null;default:1;comment:状态  1:启用,0:禁用" json:"status"`                  // 状态  1:启用,0:禁用
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}
