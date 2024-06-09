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

// UpdateFlashPromotionSessionStatusLogic 更新状态
type UpdateFlashPromotionSessionStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateFlashPromotionSessionStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFlashPromotionSessionStatusLogic {
	return &UpdateFlashPromotionSessionStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdateFlashPromotionSessionStatus 更新状态
func (l *UpdateFlashPromotionSessionStatusLogic) UpdateFlashPromotionSessionStatus(req *types.UpdateFlashPromotionSessionStatusReq) (resp *types.UpdateFlashPromotionSessionStatusResp, err error) {
	_, err = l.svcCtx.FlashPromotionSessionService.UpdateFlashPromotionSessionStatus(l.ctx, &smsclient.UpdateFlashPromotionSessionStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %v,更新状态异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新状态失败")
	}

	return &types.UpdateFlashPromotionSessionStatusResp{
		Code:    "000000",
		Message: "更新状态成功",
	}, nil

}
