package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddHomeRecommendSubjectLogic 添加首页推荐专题
/*
Author: LiuFeiHua
Date: 2024/6/12 11:30
*/
type AddHomeRecommendSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomeRecommendSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomeRecommendSubjectLogic {
	return &AddHomeRecommendSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddHomeRecommendSubject 添加首页推荐专题
func (l *AddHomeRecommendSubjectLogic) AddHomeRecommendSubject(in *smsclient.AddHomeRecommendSubjectReq) (*smsclient.AddHomeRecommendSubjectResp, error) {
	for _, data := range in.RecommendSubjectAddData {
		q := query.SmsHomeRecommendSubject
		count, _ := q.WithContext(l.ctx).Where(q.SubjectID.Eq(data.SubjectId)).Count()
		if count == 0 {
			_ = q.WithContext(l.ctx).Create(&model.SmsHomeRecommendSubject{
				SubjectID:       data.SubjectId,
				SubjectName:     data.SubjectName,
				RecommendStatus: data.RecommendStatus,
				Sort:            data.Sort,
			})

		}
	}

	return &smsclient.AddHomeRecommendSubjectResp{}, nil
}
