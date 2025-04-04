// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysDictType = "sys_dict_type"

// SysDictType 字典类型表
type SysDictType struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement:true;comment:字典id" json:"id"`                        // 字典id
	DictName   string     `gorm:"column:dict_name;not null;comment:字典名称" json:"dict_name"`                               // 字典名称
	DictType   string     `gorm:"column:dict_type;not null;comment:字典类型" json:"dict_type"`                               // 字典类型
	Status     int32      `gorm:"column:status;not null;comment:状态（0：停用，1:正常）" json:"status"`                            // 状态（0：停用，1:正常）
	Remark     string     `gorm:"column:remark;not null;comment:备注" json:"remark"`                                       // 备注
	CreateBy   string     `gorm:"column:create_by;not null;default:admin;comment:创建者" json:"create_by"`                  // 创建者
	CreateTime time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateBy   string     `gorm:"column:update_by;not null;comment:更新者" json:"update_by"`                                // 更新者
	UpdateTime *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
}

// TableName SysDictType's table name
func (*SysDictType) TableName() string {
	return TableNameSysDictType
}
