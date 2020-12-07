package logic

import (
	"context"
	"fmt"
	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

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

func (l *HomeRecommendSubjectListLogic) HomeRecommendSubjectList(in *sms.HomeRecommendSubjectListReq) (*sms.HomeRecommendSubjectListResp, error) {
	all, _ := l.svcCtx.SmsHomeRecommendSubjectModel.FindAll(in.Current, in.PageSize)
	//count, _ := l.svcCtx.UserModel.Count()

	var list []*sms.HomeRecommendSubjectListData
	for _, item := range *all {

		list = append(list, &sms.HomeRecommendSubjectListData{
			Id:              item.Id,
			SubjectId:       item.SubjectId,
			SubjectName:     item.SubjectName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	fmt.Println(list)
	return &sms.HomeRecommendSubjectListResp{
		Total: 10,
		List:  list,
	}, nil
}
