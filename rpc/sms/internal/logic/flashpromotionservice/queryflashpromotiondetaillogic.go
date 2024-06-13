package flashpromotionservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFlashPromotionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionDetailLogic {
	return &QueryFlashPromotionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询限时购表详情
func (l *QueryFlashPromotionDetailLogic) QueryFlashPromotionDetail(in *smsclient.QueryFlashPromotionDetailReq) (*smsclient.QueryFlashPromotionDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryFlashPromotionDetailResp{}, nil
}
