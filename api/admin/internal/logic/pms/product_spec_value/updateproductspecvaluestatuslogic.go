package product_spec_value

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

// UpdateProductSpecValueStatusLogic 更新商品规格值状态状态
/*
Author: LiuFeiHua
Date: 2025/06/17 16:14:48
*/
type UpdateProductSpecValueStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductSpecValueStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpecValueStatusLogic {
	return &UpdateProductSpecValueStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductSpecValueStatus 更新商品规格值状态
func (l *UpdateProductSpecValueStatusLogic) UpdateProductSpecValueStatus(req *types.UpdateProductSpecValueStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductSpecValueService.UpdateProductSpecValueStatus(l.ctx, &pmsclient.UpdateProductSpecValueStatusReq{
		Ids:      req.Ids,    //
		Status:   req.Status, // 状态：0->禁用；1->启用
		UpdateBy: userId,     // 更新人ID

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品规格值状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新商品规格值状态成功",
	}, nil
}
