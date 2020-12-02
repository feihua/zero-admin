package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeBrandListLogic {
	return HomeBrandListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeBrandListLogic) HomeBrandList(req types.ListHomeBrandReq) (*types.ListHomeBrandResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListHomeBrandResp{}, nil
}
