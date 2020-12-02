package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &types.DeleteReturnResonResp{}, nil
}
