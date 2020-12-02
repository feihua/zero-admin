package logic

import (
	"context"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeRecommendSubjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectAddLogic {
	return HomeRecommendSubjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(req types.AddHomeRecommendSubjectReq) (*types.AddHomeRecommendSubjectResp, error) {
	// todo: add your logic here and delete this line

	return &types.AddHomeRecommendSubjectResp{}, nil
}
