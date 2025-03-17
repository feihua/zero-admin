package subjectservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddSubjectLogic 添加专题
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type AddSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSubjectLogic {
	return &AddSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddSubject 添加专题
func (l *AddSubjectLogic) AddSubject(in *cmsclient.AddSubjectReq) (*cmsclient.AddSubjectResp, error) {
	q := query.CmsSubject

	count, err := q.WithContext(l.ctx).Where(q.Title.Eq(in.Title)).Count()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("添加专题失败"))
	}

	if count > 0 {
		return nil, errors.New(fmt.Sprintf("专题名称：%s,已存在", in.Title))
	}
	item := &model.CmsSubject{
		CategoryID:      in.CategoryId,      // 专题分类id
		Title:           in.Title,           // 专题标题
		Pic:             in.Pic,             // 专题主图
		ProductCount:    in.ProductCount,    // 关联产品数量
		RecommendStatus: in.RecommendStatus, // 推荐状态：0->不推荐；1->推荐
		CollectCount:    in.CollectCount,    // 收藏数
		ReadCount:       in.ReadCount,       // 阅读数
		CommentCount:    in.CommentCount,    // 评论数
		AlbumPics:       in.AlbumPics,       // 画册图片用逗号分割
		Description:     in.Description,     // 专题内容
		ShowStatus:      in.ShowStatus,      // 显示状态：0->不显示；1->显示
		Content:         in.Content,         // 专题内容
		ForwardCount:    in.ForwardCount,    // 转发数
		CategoryName:    in.CategoryName,    // 专题分类名称
		CreateBy:        in.CreateBy,        // 创建者
	}

	err = q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加专题失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加专题失败")
	}

	logc.Infof(l.ctx, "添加专题成功,参数：%+v", in)
	return &cmsclient.AddSubjectResp{}, nil
}
