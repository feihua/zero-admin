package subject

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/common/res"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddSubjectLogic 添加专题表
/*
Author: 刘飞华
Date: 2025/02/04 15:04:17
*/
type AddSubjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectLogic {
	return &AddSubjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// AddSubject 添加专题表
func (l *AddSubjectLogic) AddSubject(req *types.AddSubjectReq) (resp *types.BaseResp, err error) {

	createBy := l.ctx.Value("userName").(string)
	_, err = l.svcCtx.SubjectService.AddSubject(l.ctx, &cmsclient.AddSubjectReq{
		CategoryId:      req.CategoryId,      // 专题分类id
		Title:           req.Title,           // 专题标题
		Pic:             req.Pic,             // 专题主图
		ProductCount:    req.ProductCount,    // 关联产品数量
		RecommendStatus: req.RecommendStatus, // 推荐状态：0->不推荐；1->推荐
		CollectCount:    req.CollectCount,    // 收藏数
		ReadCount:       req.ReadCount,       // 阅读数
		CommentCount:    req.CommentCount,    // 评论数
		AlbumPics:       req.AlbumPics,       // 画册图片用逗号分割
		Description:     req.Description,     // 专题内容
		ShowStatus:      req.ShowStatus,      // 显示状态：0->不显示；1->显示
		Content:         req.Content,         // 专题内容
		ForwardCount:    req.ForwardCount,    // 转发数
		CategoryName:    req.CategoryName,    // 专题分类名称
		CreateBy:        createBy,            // 创建者
	})

	if err != nil {
		logc.Errorf(l.ctx, "添加专题表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return res.Success()
}
