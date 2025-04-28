package flash_promotion

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionLogic 更新限时购表
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type UpdateFlashPromotionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFlashPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionLogic {
	return &UpdateFlashPromotionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateFlashPromotion 更新限时购表
func (l *UpdateFlashPromotionLogic) UpdateFlashPromotion(req *types.UpdateFlashPromotionReq) (resp *types.BaseResp, err error) {

	_, err = l.svcCtx.FlashPromotionService.UpdateFlashPromotion(l.ctx, &smsclient.UpdateFlashPromotionReq{
		Id:        req.Id,        // 编号
		Title:     req.Title,     // 标题
		StartDate: req.StartDate, // 开始日期
		EndDate:   req.EndDate,   // 结束日期
		Status:    req.Status,    // 上下线状态

	})

	if err != nil {
		logc.Errorf(l.ctx, "更新限时购表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
