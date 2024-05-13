package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

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
	resp, err := l.svcCtx.HomeRecommendSubjectService.HomeRecommendSubjectList(l.ctx, &smsclient.HomeRecommendSubjectListReq{
		Current:         req.Current,
		PageSize:        req.PageSize,
		SubjectName:     req.SubjectName,
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询人气推荐专题列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询人气推荐专题失败")
	}

	var list []*types.ListHomeRecommendSubjectData

	for _, item := range resp.List {
		list = append(list, &types.ListHomeRecommendSubjectData{
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
