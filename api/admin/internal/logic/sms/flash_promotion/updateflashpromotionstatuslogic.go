package flash_promotion

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionStatusLogic 更新限时购表状态状态
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type UpdateFlashPromotionStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFlashPromotionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionStatusLogic {
	return &UpdateFlashPromotionStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateFlashPromotionStatus 更新限时购表状态
func (l *UpdateFlashPromotionStatusLogic) UpdateFlashPromotionStatus(req *types.UpdateFlashPromotionStatusReq) (resp *types.UpdateFlashPromotionStatusResp, err error) {
	_, err = l.svcCtx.FlashPromotionService.UpdateFlashPromotionStatus(l.ctx, &smsclient.UpdateFlashPromotionStatusReq{
		Ids:    req.Ids,    // 编号
		Status: req.Status, // 上下线状态

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购表状态失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.UpdateFlashPromotionStatusResp{
		Code:    "000000",
		Message: "更新限时购表状态成功",
	}, nil
}
