package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/oms/omsclient"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnApplyDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnApplyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnApplyDeleteLogic {
	return ReturnApplyDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnApplyDeleteLogic) ReturnApplyDelete(req types.DeleteReturnApplyReq) (*types.DeleteReturnApplyResp, error) {
	_, _ = l.svcCtx.Oms.OrderReturnApplyDelete(l.ctx, &omsclient.OrderReturnApplyDeleteReq{
		Id: req.Id,
	})

	return &types.DeleteReturnApplyResp{
		Code:    "000000",
		Message: "",
	}, nil
}
