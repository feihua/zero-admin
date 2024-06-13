package flashpromotionsessionservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFlashPromotionSessionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionSessionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionSessionDetailLogic {
	return &QueryFlashPromotionSessionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询限时购场次表详情
func (l *QueryFlashPromotionSessionDetailLogic) QueryFlashPromotionSessionDetail(in *smsclient.QueryFlashPromotionSessionDetailReq) (*smsclient.QueryFlashPromotionSessionDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryFlashPromotionSessionDetailResp{}, nil
}
