package home_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateHomeBrandLogic 首页品牌信息
/*
Author: LiuFeiHua
Date: 2024/5/13 15:53
*/
type UpdateHomeBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateHomeBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomeBrandLogic {
	return &UpdateHomeBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateHomeBrand 修改推荐品牌排序
func (l *UpdateHomeBrandLogic) UpdateHomeBrand(req *types.UpdateHomeBrandReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductBrandService.UpdateBrandSort(l.ctx, &pmsclient.UpdateProductBrandSortReq{
		Id:   req.BrandId,
		Sort: req.Sort,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新首页品牌排序失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
