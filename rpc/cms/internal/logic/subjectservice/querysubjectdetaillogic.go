package subjectservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/cms/cmsclient"
	"github.com/feihua/zero-admin/rpc/cms/gen/query"
	"github.com/feihua/zero-admin/rpc/cms/internal/svc"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QuerySubjectDetailLogic 查询专题详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QuerySubjectDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectDetailLogic {
	return &QuerySubjectDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySubjectDetail 查询专题详情
func (l *QuerySubjectDetailLogic) QuerySubjectDetail(in *cmsclient.QuerySubjectDetailReq) (*cmsclient.QuerySubjectDetailResp, error) {
	item, err := query.CmsSubject.WithContext(l.ctx).Where(query.CmsSubject.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("专题不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询专题异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询专题异常")
	}

	data := &cmsclient.QuerySubjectDetailResp{
		Id:              item.ID,                                 // 专题id
		CategoryId:      item.CategoryID,                         // 专题分类id
		Title:           item.Title,                              // 专题标题
		Pic:             item.Pic,                                // 专题主图
		ProductCount:    item.ProductCount,                       // 关联产品数量
		RecommendStatus: item.RecommendStatus,                    // 推荐状态：0->不推荐；1->推荐
		CollectCount:    item.CollectCount,                       // 收藏数
		ReadCount:       item.ReadCount,                          // 阅读数
		CommentCount:    item.CommentCount,                       // 评论数
		AlbumPics:       item.AlbumPics,                          // 画册图片用逗号分割
		Description:     item.Description,                        // 专题内容
		ShowStatus:      item.ShowStatus,                         // 显示状态：0->不显示；1->显示
		Content:         item.Content,                            // 专题内容
		ForwardCount:    item.ForwardCount,                       // 转发数
		CategoryName:    item.CategoryName,                       // 专题分类名称
		CreateBy:        item.CreateBy,                           // 创建者
		CreateTime:      time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:        item.UpdateBy,                           // 更新者
		UpdateTime:      time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	return data, nil
}
