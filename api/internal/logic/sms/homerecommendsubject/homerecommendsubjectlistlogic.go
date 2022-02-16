package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/rpc/sms/smsclient"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectListLogic {
	return HomeRecommendSubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectListLogic) HomeRecommendSubjectList(req types.ListHomeRecommendSubjectReq) (*types.ListHomeRecommendSubjectResp, error) {
	resp, err := l.svcCtx.Sms.HomeRecommendSubjectList(l.ctx, &smsclient.HomeRecommendSubjectListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("参数: %s,查询人气推荐专题列表异常:%s", string(data), err.Error())
		return nil, errorx.NewDefaultError("查询人气推荐专题失败")
	}

	for _, data := range resp.List {

		fmt.Println(data)
	}
	var list []*types.ListtHomeRecommendSubjectData

	for _, item := range resp.List {
		list = append(list, &types.ListtHomeRecommendSubjectData{
			Id:              item.Id,
			SubjectId:       item.SubjectId,
			SubjectName:     item.SubjectName,
			RecommendStatus: item.RecommendStatus,
			Sort:            item.Sort,
		})
	}

	return &types.ListHomeRecommendSubjectResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询人气推荐专题成功",
	}, nil
}
