package product_spec_value

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

// DeleteProductSpecValueLogic 删除商品规格值
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductSpecValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductSpecValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSpecValueLogic {
	return &DeleteProductSpecValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductSpecValue 删除商品规格值
func (l *DeleteProductSpecValueLogic) DeleteProductSpecValue(req *types.DeleteProductSpecValueReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductSpecValueService.DeleteProductSpecValue(l.ctx, &pmsclient.DeleteProductSpecValueReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除商品规格值失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除商品规格值成功",
	}, nil
}
