package brand

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductBrandDeleteLogic 商品品牌
/*
Author: LiuFeiHua
Date: 2024/5/13 16:54
*/
type ProductBrandDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductBrandDeleteLogic {
	return ProductBrandDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductBrandDelete 删除商品品牌
func (l *ProductBrandDeleteLogic) ProductBrandDelete(req types.DeleteProductBrandReq) (*types.DeleteProductBrandResp, error) {
	_, err := l.svcCtx.BrandService.DeleteBrand(l.ctx, &pmsclient.DeleteBrandReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除商品品牌异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除商品品牌失败")
	}

	return &types.DeleteProductBrandResp{
		Code:    "000000",
		Message: "",
	}, nil
}
