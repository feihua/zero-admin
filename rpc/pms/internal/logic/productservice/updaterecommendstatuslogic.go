package productservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateRecommendStatusLogic
/*
Author: LiuFeiHua
Date: 2024/5/15 16:10
*/
type UpdateRecommendStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRecommendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRecommendStatusLogic {
	return &UpdateRecommendStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateRecommendStatus 批量推荐商品
func (l *UpdateRecommendStatusLogic) UpdateRecommendStatus(in *pmsclient.UpdateProductStatusReq) (*pmsclient.UpdateProductStatusResp, error) {
	q := query.PmsProduct
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommandStatus, in.Status)

	if err != nil {
		return nil, err
	}

	return &pmsclient.UpdateProductStatusResp{}, nil
}
