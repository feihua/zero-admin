package product_brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductBrandLogic 删除商品品牌
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type DeleteProductBrandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductBrandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductBrandLogic {
	return &DeleteProductBrandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductBrand 删除商品品牌
func (l *DeleteProductBrandLogic) DeleteProductBrand(req *types.DeleteProductBrandReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductBrandService.DeleteProductBrand(l.ctx, &pmsclient.DeleteProductBrandReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除商品品牌失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	_, err = l.svcCtx.HomeBrandService.DeleteHomeBrand(l.ctx, &smsclient.DeleteHomeBrandReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除首页品牌异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}
	return &types.BaseResp{
		Code:    "000000",
		Message: "删除商品品牌成功",
	}, nil
}
