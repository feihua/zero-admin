package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	_, err := l.svcCtx.OrderReturnApplyService.OrderReturnApplyDelete(l.ctx, &omsclient.OrderReturnApplyDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除退货申请异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除退货申请失败")
	}

	return &types.DeleteReturnApplyResp{
		Code:    "000000",
		Message: "",
	}, nil
}
