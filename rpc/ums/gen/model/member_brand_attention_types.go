package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// MemberBrandAttention 会员关注品牌管理
/*
Author: LiuFeiHua
Date: 2025/06/10 11:44:10
*/
type MemberBrandAttention struct {
	ID             bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                        // 主键
	MemberId       int64         `bson:"memberId,omitempty" json:"memberId,omitempty"`             // 会员id
	MemberNickName string        `bson:"memberNickName,omitempty" json:"memberNickName,omitempty"` // 会员姓名
	MemberIcon     string        `bson:"memberIcon,omitempty" json:"memberIcon,omitempty"`         // 会员头像
	BrandId        int64         `bson:"brandId,omitempty" json:"brandId,omitempty"`               // 品牌id
	BrandName      string        `bson:"brandName,omitempty" json:"brandName,omitempty"`           // 品牌名称
	BrandLogo      string        `bson:"brandLogo,omitempty" json:"brandLogo,omitempty"`           // 品牌Logo
	BrandCity      string        `bson:"brandCity,omitempty" json:"brandCity,omitempty"`           // 品牌所在城市
	UpdateAt       time.Time     `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt       time.Time     `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
