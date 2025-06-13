package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductOperateLog struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductId        int64              `bson:"productId,omitempty" json:"productId,omitempty"`               // 产品id
	PriceOld         int64              `bson:"priceOld,omitempty" json:"priceOld,omitempty"`                 // 原价
	PriceNew         int64              `bson:"priceNew,omitempty" json:"priceNew,omitempty"`                 // 新价格
	SalePriceOld     int64              `bson:"salePriceOld,omitempty" json:"salePriceOld,omitempty"`         // 原售价
	SalePriceNew     int64              `bson:"salePriceNew,omitempty" json:"salePriceNew,omitempty"`         // 新售价
	GiftPointOld     int32              `bson:"giftPointOld,omitempty" json:"giftPointOld,omitempty"`         // 赠送的积分
	GiftPointNew     int32              `bson:"giftPointNew,omitempty" json:"giftPointNew,omitempty"`         // 新的积分
	UsePointLimitOld int32              `bson:"usePointLimitOld,omitempty" json:"usePointLimitOld,omitempty"` // 使用积分限制
	UsePointLimitNew int32              `bson:"usePointLimitNew,omitempty" json:"usePointLimitNew,omitempty"` // 新的使用积分限制
	OperateMan       string             `bson:"operateMan,omitempty" json:"operateMan,omitempty"`             // 操作人
	UpdateAt         time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt         time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
