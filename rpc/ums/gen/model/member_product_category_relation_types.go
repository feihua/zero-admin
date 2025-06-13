package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// MemberProductCategoryRelation 会员与产品分类关系（用户喜欢的分类）
/*
Author: LiuFeiHua
Date: 2025/06/10 11:44:11
*/
type MemberProductCategoryRelation struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`                              // 主键
	MemberId          int64              `bson:"memberId,omitempty" json:"memberId,omitempty"`                   // 会员id
	ProductCategoryId int64              `bson:"productCategoryId,omitempty" json:"productCategoryId,omitempty"` // 商品分类id
	UpdateAt          time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt          time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
