// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOmsOrderDelivery = "oms_order_delivery"

// OmsOrderDelivery 订单收货地址表
type OmsOrderDelivery struct {
	ID               int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	OrderID          int64      `gorm:"column:order_id;not null;comment:订单ID" json:"order_id"`                                 // 订单ID
	OrderNo          string     `gorm:"column:order_no;not null;comment:订单编号" json:"order_no"`                                 // 订单编号
	ReceiverName     string     `gorm:"column:receiver_name;not null;comment:收货人姓名" json:"receiver_name"`                      // 收货人姓名
	ReceiverPhone    string     `gorm:"column:receiver_phone;not null;comment:收货人电话" json:"receiver_phone"`                    // 收货人电话
	ReceiverProvince string     `gorm:"column:receiver_province;not null;comment:省份" json:"receiver_province"`                 // 省份
	ReceiverCity     string     `gorm:"column:receiver_city;not null;comment:城市" json:"receiver_city"`                         // 城市
	ReceiverDistrict string     `gorm:"column:receiver_district;not null;comment:区县" json:"receiver_district"`                 // 区县
	ReceiverAddress  string     `gorm:"column:receiver_address;not null;comment:详细地址" json:"receiver_address"`                 // 详细地址
	DeliveryCompany  string     `gorm:"column:delivery_company;not null;comment:物流公司" json:"delivery_company"`                 // 物流公司
	DeliveryNo       string     `gorm:"column:delivery_no;not null;comment:物流单号" json:"delivery_no"`                           // 物流单号
	CreateTime       time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime       *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                    // 更新时间
	IsDeleted        int32      `gorm:"column:is_deleted;not null;comment:是否删除" json:"is_deleted"`                             // 是否删除
}

// TableName OmsOrderDelivery's table name
func (*OmsOrderDelivery) TableName() string {
	return TableNameOmsOrderDelivery
}
