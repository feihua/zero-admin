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

// DeleteTopicCommentLogic 删除专题评论
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type DeleteTopicCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTopicCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTopicCommentLogic {
	return &DeleteTopicCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteTopicComment 删除专题评论
func (l *DeleteTopicCommentLogic) DeleteTopicComment(in *cmsclient.DeleteTopicCommentReq) (*cmsclient.DeleteTopicCommentResp, error) {
	q := query.CmsTopicComment

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		logc.Errorf(l.ctx, "删除专题评论失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除专题评论失败")
	}

	logc.Infof(l.ctx, "删除专题评论成功,参数：%+v", in)
	return &cmsclient.DeleteTopicCommentResp{}, nil
}
