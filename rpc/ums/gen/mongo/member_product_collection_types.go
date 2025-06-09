package mymongo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MemberProductCollection 用户收藏的商品
/*
Author: LiuFeiHua
Date: 2025/06/10 11:44:11
*/
type MemberProductCollection struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                          // 主键
	MemberId        int64              `bson:"memberId,omitempty" json:"memberId,omitempty"`               // 会员id
	MemberNickName  string             `bson:"memberNickName,omitempty" json:"memberNickName,omitempty"`   // 会员姓名
	MemberIcon      string             `bson:"memberIcon,omitempty" json:"memberIcon,omitempty"`           // 会员头像
	ProductId       int64              `bson:"productId,omitempty" json:"productId,omitempty"`             // 商品id
	ProductName     string             `bson:"productName,omitempty" json:"productName,omitempty"`         // 商品名称
	ProductPic      string             `bson:"productPic,omitempty" json:"productPic,omitempty"`           // 商品图片
	ProductSubTitle string             `bson:"productSubTitle,omitempty" json:"productSubTitle,omitempty"` // 商品标题
	ProductPrice    int64              `bson:"productPrice,omitempty" json:"productPrice,omitempty"`       // 商品价格
	UpdateAt        time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt        time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
