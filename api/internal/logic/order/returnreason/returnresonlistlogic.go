package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ReturnResonListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReturnResonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ReturnResonListLogic {
	return ReturnResonListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReturnResonListLogic) ReturnResonList(req types.ListReturnResonReq) (*types.ListReturnResonResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListReturnResonResp{}, nil
}
