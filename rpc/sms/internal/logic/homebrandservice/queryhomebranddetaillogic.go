package homebrandservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHomeBrandDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeBrandDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeBrandDetailLogic {
	return &QueryHomeBrandDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询首页推荐品牌表详情
func (l *QueryHomeBrandDetailLogic) QueryHomeBrandDetail(in *smsclient.QueryHomeBrandDetailReq) (*smsclient.QueryHomeBrandDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryHomeBrandDetailResp{}, nil
}
