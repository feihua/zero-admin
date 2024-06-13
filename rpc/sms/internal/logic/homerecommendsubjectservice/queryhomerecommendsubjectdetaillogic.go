package homerecommendsubjectservicelogic

import (
	"context"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHomeRecommendSubjectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeRecommendSubjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeRecommendSubjectDetailLogic {
	return &QueryHomeRecommendSubjectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询首页推荐专题表详情
func (l *QueryHomeRecommendSubjectDetailLogic) QueryHomeRecommendSubjectDetail(in *smsclient.QueryHomeRecommendSubjectDetailReq) (*smsclient.QueryHomeRecommendSubjectDetailResp, error) {
	// todo: add your logic here and delete this line

	return &smsclient.QueryHomeRecommendSubjectDetailResp{}, nil
}
