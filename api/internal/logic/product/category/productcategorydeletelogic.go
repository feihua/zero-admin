package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

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
	_, err := l.svcCtx.Pms.ProductCategoryDelete(l.ctx, &pmsclient.ProductCategoryDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除商品类别异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除商品类别失败")
	}

	return &types.DeleteProductCategoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
