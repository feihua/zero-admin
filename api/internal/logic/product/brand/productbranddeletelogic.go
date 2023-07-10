package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

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

func (l *ProductBrandDeleteLogic) ProductBrandDelete(req types.DeleteProductBrandReq) (*types.DeleteProductBrandResp, error) {
	_, err := l.svcCtx.BrandService.BrandDelete(l.ctx, &pmsclient.BrandDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除商品品牌异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除商品品牌失败")
	}

	return &types.DeleteProductBrandResp{
		Code:    "000000",
		Message: "",
	}, nil
}
