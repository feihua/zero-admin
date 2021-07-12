package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
		logx.Errorf("根据Id: %d,删除商品类别异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除商品类别失败")
	}

	return &types.DeleteProductCategoryResp{
		Code:    "000000",
		Message: "",
	}, nil
}
