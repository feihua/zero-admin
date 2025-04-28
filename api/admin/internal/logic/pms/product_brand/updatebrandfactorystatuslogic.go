package product_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateBrandFactoryStatusLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 16:55
*/
type UpdateBrandFactoryStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBrandFactoryStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandFactoryStatusLogic {
	return &UpdateBrandFactoryStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateBrandFactoryStatus 批量更新厂家制造商状态
func (l *UpdateBrandFactoryStatusLogic) UpdateBrandFactoryStatus(req *types.UpdateProductBrandStatusReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.BrandService.UpdateBrandFactoryStatus(l.ctx, &pmsclient.UpdateBrandFactoryStatusReq{
		Ids:           req.Ids,
		FactoryStatus: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量更新厂家制造商状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量更新厂家制造商状态失败")
	}

	return res.Success()
}
