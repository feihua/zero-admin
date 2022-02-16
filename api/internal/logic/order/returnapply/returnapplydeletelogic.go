package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/core/logx"
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
	_, err := l.svcCtx.Oms.OrderReturnApplyDelete(l.ctx, &omsclient.OrderReturnApplyDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除退货申请异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除退货申请失败")
	}

	return &types.DeleteReturnApplyResp{
		Code:    "000000",
		Message: "",
	}, nil
}
