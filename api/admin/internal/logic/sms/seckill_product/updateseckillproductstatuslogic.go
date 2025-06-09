package seckill_product

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateSeckillProductStatusLogic 更新秒杀商品状态状态
/*
Author: LiuFeiHua
Date: 2025/06/12 10:02:03
*/
type UpdateSeckillProductStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSeckillProductStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSeckillProductStatusLogic {
	return &UpdateSeckillProductStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateSeckillProductStatus 更新秒杀商品状态
func (l *UpdateSeckillProductStatusLogic) UpdateSeckillProductStatus(req *types.UpdateSeckillProductStatusReq) (resp *types.BaseResp, err error) {
	userId, err := common.GetUserId(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.SeckillProductService.UpdateSeckillProductStatus(l.ctx, &smsclient.UpdateSeckillProductStatusReq{
		Ids:      req.Ids,    // ID
		Status:   req.Status, // 状态：0-未上架，1-已上架
		UpdateBy: userId,     // 更新人ID

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新秒杀商品状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.BaseResp{
		Code:    "000000",
		Message: "更新秒杀商品状态成功",
	}, nil
}
