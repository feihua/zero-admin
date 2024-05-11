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

type ReturnResonDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonDeleteLogic {
	return ReturnResonDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonDeleteLogic) ReturnResonDelete(req types.DeleteReturnResonReq) (*types.DeleteReturnResonResp, error) {
	_, err := l.svcCtx.OrderReturnReasonService.OrderReturnReasonDelete(l.ctx, &omsclient.OrderReturnReasonDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除退货原因异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除退货原因失败")
	}
	return &types.DeleteReturnResonResp{
		Code:    "000000",
		Message: "",
	}, nil
}
