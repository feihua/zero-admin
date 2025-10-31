package home_recommend_subject

import (
	"context"
	"strings"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"

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
	subjectList, err := l.svcCtx.SubjectService.QuerySubjectList(l.ctx, &cmsclient.QuerySubjectListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		Title:           strings.TrimSpace(req.SubjectName), // 专题标题
		RecommendStatus: req.RecommendStatus,                // 推荐状态：0->不推荐；1->推荐
		ShowStatus:      1,                                  // 显示状态：0->不显示；1->显示
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询专题列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询专题失败")
	}

	var list []*types.ListHomeRecommendSubjectData

	for _, item := range subjectList.List {
		list = append(list, &types.ListHomeRecommendSubjectData{
			Id:              item.Id,              // 专题id
			CategoryId:      item.CategoryId,      // 专题分类id
			Title:           item.Title,           // 专题标题
			Pic:             item.Pic,             // 专题主图
			ProductCount:    item.ProductCount,    // 关联产品数量
			RecommendStatus: item.RecommendStatus, // 推荐状态：0->不推荐；1->推荐
			CollectCount:    item.CollectCount,    // 收藏数
			ReadCount:       item.ReadCount,       // 阅读数
			CommentCount:    item.CommentCount,    // 评论数
			AlbumPics:       item.AlbumPics,       // 画册图片用逗号分割
			Description:     item.Description,     // 专题内容
			ShowStatus:      item.ShowStatus,      // 显示状态：0->不显示；1->显示
			Content:         item.Content,         // 专题内容
			ForwardCount:    item.ForwardCount,    // 转发数
			CategoryName:    item.CategoryName,    // 专题分类名称
			CreateBy:        item.CreateBy,        // 创建者
			CreateTime:      item.CreateTime,      // 创建时间
			UpdateBy:        item.UpdateBy,        // 更新者
			UpdateTime:      item.UpdateTime,      // 更新时间
			Sort:            item.Sort,            // 排序
		})
	}

	return &types.ListHomeRecommendSubjectResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    subjectList.Total,
		Code:     "000000",
		Message:  "查询人气推荐专题成功",
	}, nil
}
