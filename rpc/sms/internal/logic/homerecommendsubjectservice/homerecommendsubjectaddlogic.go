package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/model"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendSubjectAddLogic 专题推荐
/*
Author: LiuFeiHua
Date: 2024/5/6 17:28
*/
type HomeRecommendSubjectAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectAddLogic {
	return &HomeRecommendSubjectAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeRecommendSubjectAdd 添加推荐专题
func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(in *smsclient.HomeRecommendSubjectAddReq) (*smsclient.HomeRecommendSubjectAddResp, error) {

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

	return &smsclient.HomeRecommendSubjectAddResp{}, nil
}
