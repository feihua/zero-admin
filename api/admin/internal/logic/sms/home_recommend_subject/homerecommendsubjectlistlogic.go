package home_recommend_subject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// HomeRecommendSubjectListLogic 人气推荐专题
/*
Author: LiuFeiHua
Date: 2024/5/14 9:43
*/
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

// HomeRecommendSubjectList 查询人气推荐专题
func (l *HomeRecommendSubjectListLogic) HomeRecommendSubjectList(req *types.ListHomeRecommendSubjectReq) (*types.ListHomeRecommendSubjectResp, error) {
	resp, err := l.svcCtx.HomeRecommendSubjectService.QueryHomeRecommendSubjectList(l.ctx, &smsclient.QueryHomeRecommendSubjectListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		SubjectName:     strings.TrimSpace(req.SubjectName),
		RecommendStatus: req.RecommendStatus,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询人气推荐专题列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListHomeRecommendSubjectData

	for _, item := range resp.List {
		list = append(list, &types.ListHomeRecommendSubjectData{
			Id:              item.Id,              //
			SubjectId:       item.SubjectId,       // 专题id
			SubjectName:     item.SubjectName,     // 专题名称
			RecommendStatus: item.RecommendStatus, // 推荐状态：0->不推荐;1->推荐
			Sort:            item.Sort,            // 排序
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
