package home

import (
	"context"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"

	"github.com/feihua/zero-admin/api/front/internal/svc"
	"github.com/feihua/zero-admin/api/front/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySubjectListLogic 分页获取专题
/*
Author: LiuFeiHua
Date: 2024/5/16 15:11
*/
type QuerySubjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectListLogic {
	return &QuerySubjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySubjectList 分页获取专题
func (l *QuerySubjectListLogic) QuerySubjectList(req *types.SubjectListReq) (resp *types.SubjectListResp, err error) {
	var subjectLists []types.SubjectList
	listResp, err := l.svcCtx.HomeRecommendSubjectService.QueryHomeRecommendSubjectList(l.ctx, &smsclient.QueryHomeRecommendSubjectListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		RecommendStatus: 1, //推荐状态：0->不推荐;1->推荐
	})

	var homeRecommendSubjectIdLists []int64
	for _, item := range listResp.List {
		homeRecommendSubjectIdLists = append(homeRecommendSubjectIdLists, item.SubjectId)
	}

	subjectListResp, _ := l.svcCtx.SubjectService.SubjectListByIds(l.ctx, &cmsclient.SubjectListByIdsReq{Ids: homeRecommendSubjectIdLists})
	for _, item := range subjectListResp.List {
		subjectLists = append(subjectLists, types.SubjectList{
			CategoryId:      item.CategoryId,
			Title:           item.Title,
			Pic:             item.Pic,
			ProductCount:    item.ProductCount,
			RecommendStatus: item.RecommendStatus,
			CollectCount:    item.CollectCount,
			ReadCount:       item.ReadCount,
			CommentCount:    item.CommentCount,
			AlbumPics:       item.AlbumPics,
			Description:     item.Description,
			ShowStatus:      item.ShowStatus,
			Content:         item.Content,
			ForwardCount:    item.ForwardCount,
			CategoryName:    item.CategoryName,
		})
	}

	return &types.SubjectListResp{
		Code:    0,
		Message: "分页获取专题成功",
		Data:    subjectLists,
	}, nil
}
