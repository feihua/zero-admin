package commentreplayservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// CommentReplayDeleteLogic 删除评价回复
/*
Author: LiuFeiHua
Date: 2024/5/8 9:49
*/
type CommentReplayDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentReplayDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentReplayDeleteLogic {
	return &CommentReplayDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CommentReplayDelete 删除评价回复
func (l *CommentReplayDeleteLogic) CommentReplayDelete(in *pmsclient.CommentReplayDeleteReq) (*pmsclient.CommentReplayDeleteResp, error) {
	q := query.PmsCommentReplay
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.CommentReplayDeleteResp{}, nil
}
