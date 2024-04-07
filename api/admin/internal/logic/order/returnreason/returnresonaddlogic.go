package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnResonAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonAddLogic {
	return ReturnResonAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonAddLogic) ReturnResonAdd(req types.AddReturnResonReq) (*types.AddReturnResonResp, error) {
	_, err := l.svcCtx.OrderReturnReasonService.OrderReturnReasonAdd(l.ctx, &omsclient.OrderReturnReasonAddReq{
		Name:   req.Name,
		Sort:   req.Sort,
		Status: req.Status,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加退货原因地址信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加退货原因失败")
	}

	return &types.AddReturnResonResp{
		Code:    "000000",
		Message: "",
	}, nil
}
