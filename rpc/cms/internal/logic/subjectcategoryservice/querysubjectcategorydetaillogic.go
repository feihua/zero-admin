package subjectcategoryservicelogic

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

// QuerySubjectCategoryDetailLogic 查询专题分类详情
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QuerySubjectCategoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectCategoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCategoryDetailLogic {
	return &QuerySubjectCategoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySubjectCategoryDetail 查询专题分类详情
func (l *QuerySubjectCategoryDetailLogic) QuerySubjectCategoryDetail(in *cmsclient.QuerySubjectCategoryDetailReq) (*cmsclient.QuerySubjectCategoryDetailResp, error) {
	item, err := query.CmsSubjectCategory.WithContext(l.ctx).Where(query.CmsSubjectCategory.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return nil, errors.New("专题分类不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询专题分类异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errors.New("查询专题分类异常")
	}

	data := &cmsclient.QuerySubjectCategoryDetailResp{
		Id:           item.ID,                                 // 主键ID
		Name:         item.Name,                               // 专题分类名称
		Icon:         item.Icon,                               // 分类图标
		SubjectCount: item.SubjectCount,                       // 专题数量
		ShowStatus:   item.ShowStatus,                         // 显示状态：0->不显示；1->显示
		Sort:         item.Sort,                               // 排序
		CreateBy:     item.CreateBy,                           // 创建者
		CreateTime:   time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:     item.UpdateBy,                           // 更新者
		UpdateTime:   time_util.TimeToString(item.UpdateTime), // 更新时间
	}

	logc.Infof(l.ctx, "查询专题分类详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
