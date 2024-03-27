package logic

import (
	"context"
	"encoding/json"
	"github.com/feihua/zero-admin/api/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/internal/svc"
	"github.com/feihua/zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeRecommendSubjectAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeRecommendSubjectAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeRecommendSubjectAddLogic {
	return HomeRecommendSubjectAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeRecommendSubjectAddLogic) HomeRecommendSubjectAdd(req types.AddHomeRecommendSubjectReq) (*types.AddHomeRecommendSubjectResp, error) {
	listResp, _ := l.svcCtx.SubjectService.SubjectListByIds(l.ctx, &cmsclient.SubjectListByIdsReq{
		Ids: req.SubjectIds,
	})

	var list []*smsclient.HomeRecommendSubjectAddData

	for _, item := range listResp.List {
		list = append(list, &smsclient.HomeRecommendSubjectAddData{
			SubjectId:       item.Id,
			SubjectName:     item.Title,
			RecommendStatus: item.RecommendStatus,
			Sort:            1,
		})
	}

	_, err := l.svcCtx.HomeRecommendSubjectService.HomeRecommendSubjectAdd(l.ctx, &smsclient.HomeRecommendSubjectAddReq{
		RecommendSubjectAddData: list,
	})

	if err != nil {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("添加人气推荐专题信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errorx.NewDefaultError("添加人气推荐专题失败")
	}

	return &types.AddHomeRecommendSubjectResp{
		Code:    "000000",
		Message: "添加人气推荐专题成功",
	}, nil
}
