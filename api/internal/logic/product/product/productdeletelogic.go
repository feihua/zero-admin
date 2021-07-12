package logic

import (
	"context"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/pms/pmsclient"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	_, err := l.svcCtx.Pms.ProductDelete(l.ctx, &pmsclient.ProductDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		logx.Errorf("根据Id: %d,删除商品信息异常:%s", req.Id, err.Error())
		return nil, errorx.NewDefaultError("删除商品信息失败")
	}

	return &types.DeleteProductResp{
		Code:    "000000",
		Message: "",
	}, nil
}
