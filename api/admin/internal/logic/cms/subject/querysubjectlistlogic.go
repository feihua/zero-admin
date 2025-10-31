package subject

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

// QuerySubjectListLogic 查询专题表列表
/*
Author: 刘飞华
Date: 2025/02/04 15:04:17
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

// QuerySubjectList 查询专题表列表
func (l *QuerySubjectListLogic) QuerySubjectList(req *types.QuerySubjectListReq) (resp *types.QuerySubjectListResp, err error) {
	subjectList, err := l.svcCtx.SubjectService.QuerySubjectList(l.ctx, &cmsclient.QuerySubjectListReq{
		PageNum:         req.Current,
		PageSize:        req.PageSize,
		Title:           strings.TrimSpace(req.Title), // 专题标题
		RecommendStatus: req.RecommendStatus,          // 推荐状态：0->不推荐；1->推荐
		ShowStatus:      req.ShowStatus,               // 显示状态：0->不显示；1->显示
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询专题列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询专题失败")
	}

	var list []*types.QuerySubjectListData

	for _, item := range subjectList.List {
		list = append(list, &types.QuerySubjectListData{
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

	return &types.QuerySubjectListResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    subjectList.Total,
		Code:     "000000",
		Message:  "查询专题成功",
	}, nil
}
