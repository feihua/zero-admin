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

// UpdateProductSpecValueLogic 更新商品规格值
/*
Author: LiuFeiHua
Date: 2025/06/16 14:37:37
*/
type UpdateProductSpecValueLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProductSpecValueLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductSpecValueLogic {
	return &UpdateProductSpecValueLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateProductSpecValue 更新商品规格值
func (l *UpdateProductSpecValueLogic) UpdateProductSpecValue(req *types.UpdateProductSpecValueReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.ProductSpecValueService.UpdateProductSpecValue(l.ctx, &pmsclient.UpdateProductSpecValueReq{
		Id:       req.Id,     //
		SpecId:   req.SpecId, // 规格ID
		Value:    req.Value,  // 规格值
		Sort:     req.Sort,   // 排序
		Status:   req.Status, // 状态：0->禁用；1->启用
		UpdateBy: userId,     // 更新人ID
	})

	if err != nil {
		logc.Errorf(l.ctx, "更新商品规格值失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新商品规格值成功",
	}, nil
}
