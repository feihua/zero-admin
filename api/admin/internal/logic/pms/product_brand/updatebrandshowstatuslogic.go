package product_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateBrandShowStatusLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 16:55
*/
type UpdateBrandShowStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBrandShowStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBrandShowStatusLogic {
	return &UpdateBrandShowStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateBrandShowStatus 批量更新显示状态
func (l *UpdateBrandShowStatusLogic) UpdateBrandShowStatus(req *types.UpdateProductBrandStatusReq) (resp *types.UpdateProductBrandStatusResp, err error) {
	_, err = l.svcCtx.BrandService.UpdateBrandShowStatus(l.ctx, &pmsclient.UpdateBrandShowStatusReq{
		Ids:        req.Ids,
		ShowStatus: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "批量更新显示状态失败,参数：%+v,响应：%s", req, err.Error())
		return nil, errorx.NewDefaultError("批量更新显示状态失败")
	}

	return &types.UpdateProductBrandStatusResp{
		Code:    "000000",
		Message: "批量更新显示状态成功",
	}, nil
}
