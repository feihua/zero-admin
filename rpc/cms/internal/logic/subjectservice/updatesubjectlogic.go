package subjectservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/model"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

// UpdateSubjectLogic 更新专题
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type UpdateSubjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubjectLogic {
	return &UpdateSubjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateSubject 更新专题
func (l *UpdateSubjectLogic) UpdateSubject(in *cmsclient.UpdateSubjectReq) (*cmsclient.UpdateSubjectResp, error) {
	q := query.CmsSubject.WithContext(l.ctx)

	// 1.根据专题id查询专题是否已存在
	s, err := q.Where(query.CmsSubject.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("专题不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询专题异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询专题异常")
	}

	now := time.Now()
	item := &model.CmsSubject{
		ID:              in.Id,              // 专题id
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
		CreateBy:        s.CreateBy,         // 创建者
		CreateTime:      s.CreateTime,       // 创建时间
		UpdateBy:        in.UpdateBy,        // 更新者
		UpdateTime:      &now,               // 更新时间
	}

	// 2.专题存在时,则直接更新专题
	err = l.svcCtx.DB.Model(&model.CmsSubject{}).WithContext(l.ctx).Where(query.CmsSubject.ID.Eq(in.Id)).Save(item).Error

	if err != nil {
		logc.Errorf(l.ctx, "更新专题失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新专题失败")
	}

	return &cmsclient.UpdateSubjectResp{}, nil
}
