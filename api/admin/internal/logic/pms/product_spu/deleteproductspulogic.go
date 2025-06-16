package product_spu

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteProductSpuLogic 删除商品SPU
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductSpuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductSpuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSpuLogic {
	return &DeleteProductSpuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductSpu 删除商品SPU
func (l *DeleteProductSpuLogic) DeleteProductSpu(req *types.DeleteProductSpuReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductSpuService.UpdateDeleteStatus(l.ctx, &pmsclient.UpdateProductSpuStatusReq{
		Ids:      req.Ids,
		Status:   1,
		UpdateBy: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除商品SPU失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除商品SPU成功",
	}, nil
}
