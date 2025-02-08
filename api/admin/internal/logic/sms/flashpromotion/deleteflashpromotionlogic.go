package flashpromotion

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

// DeleteFlashPromotionLogic 删除限时购表
/*
Author: 刘飞华
Date: 2025/02/07 10:11:43
*/
type DeleteFlashPromotionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFlashPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFlashPromotionLogic {
	return &DeleteFlashPromotionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteFlashPromotion 删除限时购表
func (l *DeleteFlashPromotionLogic) DeleteFlashPromotion(req *types.DeleteFlashPromotionReq) (resp *types.DeleteFlashPromotionResp, err error) {
	_, err = l.svcCtx.FlashPromotionService.DeleteFlashPromotion(l.ctx, &smsclient.DeleteFlashPromotionReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "删除限时购表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteFlashPromotionResp{
		Code:    "000000",
		Message: "删除限时购表成功",
	}, nil
}
