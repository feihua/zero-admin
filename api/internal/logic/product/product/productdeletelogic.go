package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductDeleteLogic {
	return ProductDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductDeleteLogic) ProductDelete(req types.DeleteProductReq) (*types.DeleteProductResp, error) {
	_, err := l.svcCtx.ProductService.ProductDelete(l.ctx, &pmsclient.ProductDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据Id: %d,删除商品信息异常:%s", req.Ids, err.Error())
		return nil, errorx.NewDefaultError("删除商品信息失败")
	}

	return &types.DeleteProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
