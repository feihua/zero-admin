package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/oms/omsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.Oms.OrderReturnReasonAdd(l.ctx, &omsclient.OrderReturnReasonAddReq{
		Name:   req.Name,
		Sort:   req.Sort,
		Status: req.Status,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.Errorf("添加退货原因地址参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加退货原因失败")
	}

	return &types.AddReturnResonResp{
		Code:    "000000",
		Message: "",
	}, nil
}
