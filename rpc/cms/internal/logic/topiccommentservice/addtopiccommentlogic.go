package topiccommentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddTopicCommentLogic 添加专题评论
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type AddTopicCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTopicCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTopicCommentLogic {
	return &AddTopicCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddTopicComment 添加专题评论
func (l *AddTopicCommentLogic) AddTopicComment(in *cmsclient.AddTopicCommentReq) (*cmsclient.AddTopicCommentResp, error) {
	q := query.CmsTopicComment

	item := &model.CmsTopicComment{
		MemberNickName: in.MemberNickName, // 评论人员昵称
		TopicID:        in.TopicId,        // 专题ID
		MemberIcon:     in.MemberIcon,     // 评论人员头像
		Content:        in.Content,        // 评论内容
		ShowStatus:     in.ShowStatus,     // 是否显示，0->不显示；1->显示
	}

	err := q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加专题评论失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加专题评论失败")
	}

	return &cmsclient.AddTopicCommentResp{}, nil
}
