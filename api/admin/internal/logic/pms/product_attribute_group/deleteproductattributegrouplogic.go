package product_attribute_group

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductAttributeGroupLogic 删除商品属性分组
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductAttributeGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductAttributeGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductAttributeGroupLogic {
	return &DeleteProductAttributeGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductAttributeGroup 删除商品属性分组
func (l *DeleteProductAttributeGroupLogic) DeleteProductAttributeGroup(req *types.DeleteProductAttributeGroupReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductAttributeGroupService.DeleteProductAttributeGroup(l.ctx, &pmsclient.DeleteProductAttributeGroupReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除商品属性分组失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除商品属性分组成功",
	}, nil
}
