package attribute

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ProductAttributeDeleteLogic 商品属性
/*
Author: LiuFeiHua
Date: 2024/5/15 14:37
*/
type ProductAttributeDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductAttributeDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAttributeDeleteLogic {
	return &ProductAttributeDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// ProductAttributeDelete 删除商品属性
func (l *ProductAttributeDeleteLogic) ProductAttributeDelete(req *types.DeleteProductAttributeReq) (resp *types.DeleteProductAttributeResp, err error) {
	_, err = l.svcCtx.ProductAttributeService.ProductAttributeDelete(l.ctx, &pmsclient.ProductAttributeDeleteReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除属性异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除属性失败")
	}

	return &types.DeleteProductAttributeResp{
		Code:    "000000",
		Message: "",
	}, nil
}
