package commentreplayservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *CommentReplayDeleteLogic) CommentReplayDelete(in *pmsclient.CommentReplayDeleteReq) (*pmsclient.CommentReplayDeleteResp, error) {
	err := l.svcCtx.PmsCommentReplayModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &pmsclient.CommentReplayDeleteResp{}, nil
}
