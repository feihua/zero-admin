package productbrandservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/internal/svc"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateProductBrandStatusLogic 更新商品品牌
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type UpdateProductBrandStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductBrandStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductBrandStatusLogic {
	return &UpdateProductBrandStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateProductBrandStatus 更新商品品牌状态
func (l *UpdateProductBrandStatusLogic) UpdateProductBrandStatus(in *pmsclient.UpdateProductBrandStatusReq) (*pmsclient.UpdateProductBrandStatusResp, error) {
	q := query.PmsProductBrand

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).UpdateSimple(q.IsEnabled.Value(in.Status), q.UpdateBy.Value(in.UpdateBy))

	if err != nil {
		logc.Errorf(l.ctx, "更新商品品牌状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新商品品牌状态失败")
	}

	return &pmsclient.UpdateProductBrandStatusResp{}, nil
}
