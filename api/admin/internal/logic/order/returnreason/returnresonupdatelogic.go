package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnResonUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonUpdateLogic {
	return ReturnResonUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonUpdateLogic) ReturnResonUpdate(req types.UpdateReturnResonReq) (*types.UpdateReturnResonResp, error) {
	_, err := l.svcCtx.OrderReturnReasonService.OrderReturnReasonUpdate(l.ctx, &omsclient.OrderReturnReasonUpdateReq{
		Id:     req.Id,
		Name:   req.Name,
		Sort:   req.Sort,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新退货原因信息失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新退货原因失败")
	}

	return &types.UpdateReturnResonResp{
		Code:    "000000",
		Message: "更新退货原因成功",
	}, nil
}
