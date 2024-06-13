package flashpromotionlogservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFlashPromotionLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionLogDetailLogic {
	return &QueryFlashPromotionLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询限时购通知记录详情
func (l *QueryFlashPromotionLogDetailLogic) QueryFlashPromotionLogDetail(in *smsclient.QueryFlashPromotionLogDetailReq) (*smsclient.QueryFlashPromotionLogDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryFlashPromotionLogDetailResp{}, nil
}
