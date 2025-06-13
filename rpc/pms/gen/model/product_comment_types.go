package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductComment struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductId        int64              `bson:"productId,omitempty" json:"productId,omitempty"`               // 商品id
	MemberNickName   string             `bson:"memberNickName,omitempty" json:"memberNickName,omitempty"`     // 评价者昵称
	ProductName      string             `bson:"productName,omitempty" json:"productName,omitempty"`           // 商品名称
	Star             int32              `bson:"star,omitempty" json:"star,omitempty"`                         // 评价星数：0->5
	MemberIp         string             `bson:"memberIp,omitempty" json:"memberIp,omitempty"`                 // 评价的ip
	ShowStatus       int32              `bson:"showStatus,omitempty" json:"showStatus,omitempty"`             // 是否显示，0-不显示，1-显示
	ProductAttribute string             `bson:"productAttribute,omitempty" json:"productAttribute,omitempty"` // 购买时的商品属性
	CollectCount     int32              `bson:"collectCount,omitempty" json:"collectCount,omitempty"`         // 点赞数
	ReadCount        int32              `bson:"readCount,omitempty" json:"readCount,omitempty"`               // 阅读数
	Content          string             `bson:"content,omitempty" json:"content,omitempty"`                   // 内容
	Pics             string             `bson:"pics,omitempty" json:"pics,omitempty"`                         // 上传图片地址，以逗号隔开
	MemberIcon       string             `bson:"memberIcon,omitempty" json:"memberIcon,omitempty"`             // 评论用户头像
	ReplayCount      int32              `bson:"replayCount,omitempty" json:"replayCount,omitempty"`           // 回复数量
	UpdateAt         time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt         time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
