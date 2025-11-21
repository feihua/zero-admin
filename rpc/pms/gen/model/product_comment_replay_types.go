package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type ProductCommentReplay struct {
	ID             bson.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CommentId      int64         `bson:"commentId,omitempty" json:"commentId,omitempty"`           // 评论id
	MemberNickName string        `bson:"memberNickName,omitempty" json:"memberNickName,omitempty"` // 评论人员昵称
	MemberIcon     string        `bson:"memberIcon,omitempty" json:"memberIcon,omitempty"`         // 评论人员头像
	Content        string        `bson:"content,omitempty" json:"content,omitempty"`               // 内容
	Type           int32         `bson:"type,omitempty" json:"type,omitempty"`                     // 评论人员类型；0->会员；1->管理员
	UpdateAt       time.Time     `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt       time.Time     `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
