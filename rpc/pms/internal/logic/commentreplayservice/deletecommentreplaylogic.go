package commentreplayservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteCommentReplayLogic 删除产品评价回复
/*
Author: LiuFeiHua
Date: 2024/6/12 16:33
*/
type DeleteCommentReplayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentReplayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentReplayLogic {
	return &DeleteCommentReplayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteCommentReplay 删除产品评价回复
func (l *DeleteCommentReplayLogic) DeleteCommentReplay(in *pmsclient.DeleteCommentReplayReq) (*pmsclient.DeleteCommentReplayResp, error) {
	_, err := l.svcCtx.ProductCommentReplayModel.Delete(l.ctx, in.Id)

	if err != nil {
		logc.Errorf(l.ctx, "删除产品评价回复失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除产品评价回复失败")
	}

	return &pmsclient.DeleteCommentReplayResp{}, nil
}
