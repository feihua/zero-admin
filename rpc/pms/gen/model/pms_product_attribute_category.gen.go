// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePmsProductAttributeCategory = "pms_product_attribute_category"

// PmsProductAttributeCategory 产品属性分类表
type PmsProductAttributeCategory struct {
	ID             int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name           string `gorm:"column:name;not null;comment:商品属性分类名称" json:"name"`                   // 商品属性分类名称
	AttributeCount int32  `gorm:"column:attribute_count;not null;comment:属性数量" json:"attribute_count"` // 属性数量
	ParamCount     int32  `gorm:"column:param_count;not null;comment:参数数量" json:"param_count"`         // 参数数量
}

// TableName PmsProductAttributeCategory's table name
func (*PmsProductAttributeCategory) TableName() string {
	return TableNamePmsProductAttributeCategory
}
