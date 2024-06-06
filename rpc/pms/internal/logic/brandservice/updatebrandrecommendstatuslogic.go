package brandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateBrandRecommendStatusLogic 更新品牌的推荐状态
/*
Author: LiuFeiHua
Date: 2024/6/06 17:03
*/
type UpdateBrandRecommendStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandRecommendStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandRecommendStatusLogic {
	return &UpdateBrandRecommendStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateBrandRecommendStatus 更新品牌的推荐状态
func (l *UpdateBrandRecommendStatusLogic) UpdateBrandRecommendStatus(in *pmsclient.UpdateBrandRecommendStatusReq) (*pmsclient.UpdateBrandRecommendStatusResp, error) {
	brand := query.PmsBrand

	_, err := brand.WithContext(l.ctx).Where(brand.ID.In(in.Ids...)).Update(brand.RecommendStatus, in.RecommendStatus)
	if err != nil {
		logc.Errorf(l.ctx, "更新品牌的推荐状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新品牌的推荐状态失败")
	}

	return &pmsclient.UpdateBrandRecommendStatusResp{}, nil
}
