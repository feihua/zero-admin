package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductCollect struct {
	ID          bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MemberId    int64         `bson:"memberId,omitempty" json:"userId,omitempty"`         // 用户表的用户ID
	ValueId     int32         `bson:"valueId,omitempty" json:"valueId,omitempty"`         // 如果type=0，则是商品ID；如果type=1，则是专题ID
	CollectType int32         `bson:"collectType,omitempty" json:"collectType,omitempty"` // 收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID
	Deleted     int32         `bson:"deleted,omitempty" json:"deleted,omitempty"`         // 逻辑删除
	UpdateAt    time.Time     `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time     `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
