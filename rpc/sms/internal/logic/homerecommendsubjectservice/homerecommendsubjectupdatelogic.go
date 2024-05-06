package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendSubjectUpdateLogic
/*
Author: LiuFeiHua
Date: 2024/5/6 17:29
*/
type HomeRecommendSubjectUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectUpdateLogic {
	return &HomeRecommendSubjectUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeRecommendSubjectUpdate 更新专题
func (l *HomeRecommendSubjectUpdateLogic) HomeRecommendSubjectUpdate(in *smsclient.HomeRecommendSubjectUpdateReq) (*smsclient.HomeRecommendSubjectUpdateResp, error) {

	_, err := query.SmsHomeRecommendSubject.WithContext(l.ctx).Updates(&model.SmsHomeRecommendSubject{
		ID:              in.Id,
		SubjectID:       in.SubjectId,
		SubjectName:     in.SubjectName,
		RecommendStatus: in.RecommendStatus,
		Sort:            in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.HomeRecommendSubjectUpdateResp{}, nil
}
