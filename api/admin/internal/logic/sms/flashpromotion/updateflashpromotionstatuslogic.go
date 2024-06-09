package flashpromotion

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateFlashPromotionStatusLogic 更新限时购上下线
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

// UpdateFlashPromotionStatus 更新限时购上下线
func (l *UpdateFlashPromotionStatusLogic) UpdateFlashPromotionStatus(req *types.UpdateFlashPromotionStatusReq) (resp *types.UpdateFlashPromotionStatusResp, err error) {
	_, err = l.svcCtx.FlashPromotionService.UpdateFlashPromotionStatus(l.ctx, &smsclient.UpdateFlashPromotionStatusReq{
		Ids:    req.Ids,
		Status: req.Status,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据Ids: %v,更新限时购上下线异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("更新限时购上下线失败")
	}

	return &types.UpdateFlashPromotionStatusResp{
		Code:    "000000",
		Message: "更新限时购上下线成功",
	}, nil
}
