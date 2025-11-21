package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductVertifyRecord struct {
	ID        bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ProductId int64         `bson:"productId,omitempty" json:"productId,omitempty"` // 商品id
	ReviewMan string        `bson:"reviewMan,omitempty" json:"reviewMan,omitempty"` // 审核人
	Status    int32         `bson:"status,omitempty" json:"status,omitempty"`       // 审核状态：0->未通过；1->通过
	Detail    string        `bson:"detail,omitempty" json:"detail,omitempty"`       // 反馈详情
	UpdateAt  time.Time     `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt  time.Time     `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
