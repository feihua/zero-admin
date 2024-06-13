package flashpromotionproductrelationservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryFlashPromotionProductRelationDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryFlashPromotionProductRelationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryFlashPromotionProductRelationDetailLogic {
	return &QueryFlashPromotionProductRelationDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询商品限时购与商品关系表详情
func (l *QueryFlashPromotionProductRelationDetailLogic) QueryFlashPromotionProductRelationDetail(in *smsclient.QueryFlashPromotionProductRelationDetailReq) (*smsclient.QueryFlashPromotionProductRelationDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryFlashPromotionProductRelationDetailResp{}, nil
}
