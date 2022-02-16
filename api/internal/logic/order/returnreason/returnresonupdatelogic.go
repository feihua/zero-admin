package logic

import (
	"context"
	"encoding/json"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/oms/omsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Oms.OrderReturnReasonUpdate(l.ctx, &omsclient.OrderReturnReasonUpdateReq{
		Id:     req.Id,
		Name:   req.Name,
		Sort:   req.Sort,
		Status: req.Status,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("更新退货原因信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("更新退货原因失败")
	}

	return &types.UpdateReturnResonResp{
		Code:    "000000",
		Message: "更新退货原因成功",
	}, nil
}
