package subject

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QuerySubjectDetailLogic 查询专题表详情
/*
Author: 刘飞华
Date: 2025/02/04 15:04:17
*/
type QuerySubjectDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectDetailLogic {
	return &QuerySubjectDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySubjectDetail 查询专题表详情
func (l *QuerySubjectDetailLogic) QuerySubjectDetail(req *types.QuerySubjectDetailReq) (resp *types.QuerySubjectDetailResp, err error) {

	detail, err := l.svcCtx.SubjectService.QuerySubjectDetail(l.ctx, &cmsclient.QuerySubjectDetailReq{
		Id: req.Id, // 专题id
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询专题表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QuerySubjectDetailData{
		Id:              detail.Id,              // 专题id
		CategoryId:      detail.CategoryId,      // 专题分类id
		Title:           detail.Title,           // 专题标题
		Pic:             detail.Pic,             // 专题主图
		ProductCount:    detail.ProductCount,    // 关联产品数量
		RecommendStatus: detail.RecommendStatus, // 推荐状态：0->不推荐；1->推荐
		CollectCount:    detail.CollectCount,    // 收藏数
		ReadCount:       detail.ReadCount,       // 阅读数
		CommentCount:    detail.CommentCount,    // 评论数
		AlbumPics:       detail.AlbumPics,       // 画册图片用逗号分割
		Description:     detail.Description,     // 专题内容
		ShowStatus:      detail.ShowStatus,      // 显示状态：0->不显示；1->显示
		Content:         detail.Content,         // 专题内容
		ForwardCount:    detail.ForwardCount,    // 转发数
		CategoryName:    detail.CategoryName,    // 专题分类名称
		CreateBy:        detail.CreateBy,        // 创建者
		CreateTime:      detail.CreateTime,      // 创建时间
		UpdateBy:        detail.UpdateBy,        // 更新者
		UpdateTime:      detail.UpdateTime,      // 更新时间
		Sort:            detail.Sort,            // 排序
	}
	return &types.QuerySubjectDetailResp{
		Code:    "000000",
		Message: "查询专题表成功",
		Data:    data,
	}, nil
}
