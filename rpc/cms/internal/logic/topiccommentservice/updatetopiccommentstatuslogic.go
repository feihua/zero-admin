package topiccommentservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateTopicCommentStatusLogic 更新专题评论
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateTopicCommentStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTopicCommentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTopicCommentStatusLogic {
	return &UpdateTopicCommentStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateTopicCommentStatus 更新专题评论状态
func (l *UpdateTopicCommentStatusLogic) UpdateTopicCommentStatus(in *cmsclient.UpdateTopicCommentStatusReq) (*cmsclient.UpdateTopicCommentStatusResp, error) {
	q := query.CmsTopicComment

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.ShowStatus, in.ShowStatus)

	if err != nil {
		logc.Errorf(l.ctx, "更新专题评论状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新专题评论状态失败")
	}

	logc.Infof(l.ctx, "更新专题评论状态成功,参数：%+v", in)
	return &cmsclient.UpdateTopicCommentStatusResp{}, nil
}
