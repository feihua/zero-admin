package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendSubjectDeleteLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 17:28
*/
type HomeRecommendSubjectDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectDeleteLogic {
	return &HomeRecommendSubjectDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeRecommendSubjectDelete 删除专题推荐
func (l *HomeRecommendSubjectDeleteLogic) HomeRecommendSubjectDelete(in *smsclient.HomeRecommendSubjectDeleteReq) (*smsclient.HomeRecommendSubjectDeleteResp, error) {
	_, err := query.SmsHomeRecommendSubject.WithContext(l.ctx).Where(query.SmsHomeRecommendSubject.ID.In(in.Ids...)).Delete()
	if err != nil {
		return nil, err
	}

	return &smsclient.HomeRecommendSubjectDeleteResp{}, nil
}
