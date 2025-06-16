package product_spec

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

// DeleteProductSpecLogic 删除商品规格
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type DeleteProductSpecLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteProductSpecLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductSpecLogic {
	return &DeleteProductSpecLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteProductSpec 删除商品规格
func (l *DeleteProductSpecLogic) DeleteProductSpec(req *types.DeleteProductSpecReq) (resp *types.BaseResp, err error) {
	_, err = l.svcCtx.ProductSpecService.DeleteProductSpec(l.ctx, &pmsclient.DeleteProductSpecReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除商品规格失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "删除商品规格成功",
	}, nil
}
