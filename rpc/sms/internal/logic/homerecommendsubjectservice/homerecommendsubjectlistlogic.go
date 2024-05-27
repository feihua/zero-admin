package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendSubjectListLogic 专题推荐
/*
Author: LiuFeiHua
Date: 2024/5/6 17:29
*/
type HomeRecommendSubjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeRecommendSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeRecommendSubjectListLogic {
	return &HomeRecommendSubjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// HomeRecommendSubjectList 推荐专题列表
func (l *HomeRecommendSubjectListLogic) HomeRecommendSubjectList(in *smsclient.HomeRecommendSubjectListReq) (*smsclient.HomeRecommendSubjectListResp, error) {
	q := query.SmsHomeRecommendSubject.WithContext(l.ctx)
	if len(in.SubjectName) > 0 {
		q = q.Where(query.SmsHomeRecommendSubject.SubjectName.Like("%" + in.SubjectName + "%"))
	}
	if in.RecommendStatus != 2 {
		q = q.Where(query.SmsHomeRecommendSubject.RecommendStatus.Eq(in.RecommendStatus))
	}

	result, count, err := q.FindByPage(int((in.Current-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询人气专题推荐列表信息失败,参数:%+v,异常:%s", in, err.Error())
		return nil, err
	}

	var list []*smsclient.HomeRecommendSubjectListData
	for _, item := range result {
		list = append(list, &smsclient.HomeRecommendSubjectListData{
			Id:              item.ID,
			SubjectId:       item.SubjectID,
			SubjectName:     item.SubjectName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	logc.Infof(l.ctx, "查询人气专题推荐列表信息,参数：%+v,响应：%+v", in, list)
	return &smsclient.HomeRecommendSubjectListResp{
		Total: count,
		List:  list,
	}, nil
}
