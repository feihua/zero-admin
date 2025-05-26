package productbrandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

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
func (l *UpdateBrandRecommendStatusLogic) UpdateBrandRecommendStatus(in *pmsclient.UpdateProductBrandStatusReq) (*pmsclient.UpdateProductBrandStatusResp, error) {
	q := query.PmsProductBrand

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.RecommendStatus, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新商品品牌状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品品牌状态失败")
	}

	return &pmsclient.UpdateProductBrandStatusResp{}, nil
}
