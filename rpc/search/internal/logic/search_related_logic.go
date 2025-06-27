package logic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/search/internal/svc"
	"github.com/feihua/zero-admin/rpc/search/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchRelatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchRelatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchRelatedLogic {
	return &SearchRelatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SearchRelated 获取搜索的相关品牌、分类及筛选属性
func (l *SearchRelatedLogic) SearchRelated(in *search.SearchRelatedReq) (*search.SearchRelatedResp, error) {
	// todo: add your logic here and delete this line

	return &search.SearchRelatedResp{}, nil
}
