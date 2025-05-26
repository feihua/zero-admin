package product_brand

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

// UpdateProductBrandStatusLogic 更新商品品牌状态状态
/*
Author: LiuFeiHua
Date: 2025/05/26 10:33:54
*/
type UpdateProductBrandStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductBrandStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductBrandStatusLogic {
	return &UpdateProductBrandStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductBrandStatus 更新商品品牌状态
func (l *UpdateProductBrandStatusLogic) UpdateProductBrandStatus(req *types.UpdateProductBrandStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductBrandService.UpdateProductBrandStatus(l.ctx, &pmsclient.UpdateProductBrandStatusReq{
		Ids:      req.Ids,    //
		Status:   req.Status, // 状态
		UpdateBy: userId,
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品品牌状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新商品品牌状态成功",
	}, nil
}
