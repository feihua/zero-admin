package homerecommendsubjectservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sms/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryHomeRecommendSubjectListLogic 查询首页推荐专题表列表
/*
Author: LiuFeiHua
Date: 2024/6/12 17:58
*/
type QueryHomeRecommendSubjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryHomeRecommendSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHomeRecommendSubjectListLogic {
	return &QueryHomeRecommendSubjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryHomeRecommendSubjectList 查询首页推荐专题表列表
func (l *QueryHomeRecommendSubjectListLogic) QueryHomeRecommendSubjectList(in *smsclient.QueryHomeRecommendSubjectListReq) (*smsclient.QueryHomeRecommendSubjectListResp, error) {
	q := query.SmsHomeRecommendSubject.WithContext(l.ctx)
	if len(in.SubjectName) > 0 {
		q = q.Where(query.SmsHomeRecommendSubject.SubjectName.Like("%" + in.SubjectName + "%"))
	}
	if in.RecommendStatus != 2 {
		q = q.Where(query.SmsHomeRecommendSubject.RecommendStatus.Eq(in.RecommendStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

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

	return &smsclient.QueryHomeRecommendSubjectListResp{
		Total: count,
		List:  list,
	}, nil

}
