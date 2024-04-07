package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCategoryDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductCategoryDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductCategoryDeleteLogic {
	return ProductCategoryDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductCategoryDeleteLogic) ProductCategoryDelete(req types.DeleteProductCategoryReq) (*types.DeleteProductCategoryResp, error) {
	_, err := l.svcCtx.ProductCategoryService.ProductCategoryDelete(l.ctx, &pmsclient.ProductCategoryDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除商品类别异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除商品类别失败")
	}

	return &types.DeleteProductCategoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
