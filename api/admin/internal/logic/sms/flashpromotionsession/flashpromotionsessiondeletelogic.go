package flashpromotionsession

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// FlashPromotionSessionDeleteLogic 限时购场次
/*
Author: LiuFeiHua
Date: 2024/5/14 13:43
*/
type FlashPromotionSessionDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionDeleteLogic {
	return FlashPromotionSessionDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// FlashPromotionSessionDelete 删除限时购场次
func (l *FlashPromotionSessionDeleteLogic) FlashPromotionSessionDelete(req *types.DeleteFlashPromotionSessionReq) (*types.DeleteFlashPromotionSessionResp, error) {
	_, err := l.svcCtx.FlashPromotionSessionService.DeleteFlashPromotionSession(l.ctx, &smsclient.DeleteFlashPromotionSessionReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Id: %+v,删除限时购场次信息异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("删除限时购场次信息失败")
	}
	return &types.DeleteFlashPromotionSessionResp{
		Code:    "000000",
		Message: "",
	}, nil
}
