package logic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/search/internal/svc"
	"github.com/feihua/zero-admin/rpc/search/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Recommend 根据商品id推荐商品
func (l *RecommendLogic) Recommend(in *search.RecommendReq) (*search.SearchResp, error) {
	// todo: add your logic here and delete this line

	return &search.SearchResp{}, nil
}
