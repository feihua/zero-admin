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

// UpdateBrandSortLogic 更新品牌的排序
/*
Author: LiuFeiHua
Date: 2025/6/12 11:38
*/
type UpdateBrandSortLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateBrandSortLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandSortLogic {
	return &UpdateBrandSortLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateBrandSort 更新品牌的排序
func (l *UpdateBrandSortLogic) UpdateBrandSort(in *pmsclient.UpdateProductBrandSortReq) (*pmsclient.UpdateProductBrandStatusResp, error) {
	q := query.PmsProductBrand

	_, err := q.WithContext(l.ctx).Where(q.ID.Eq(in.Id)).UpdateSimple(q.Sort.Value(in.Sort), q.UpdateBy.Value(in.UpdateBy))

	if err != nil {
		logc.Errorf(l.ctx, "更新商品品牌排序失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品品牌排序失败")
	}

	return &pmsclient.UpdateProductBrandStatusResp{}, nil
}
