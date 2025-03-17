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
)

// QuerySubjectCategoryListLogic 查询专题分类列表
/*
Author: LiuFeiHua
Date: 2025/01/23 15:24:00
*/
type QuerySubjectCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuerySubjectCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCategoryListLogic {
	return &QuerySubjectCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QuerySubjectCategoryList 查询专题分类列表
func (l *QuerySubjectCategoryListLogic) QuerySubjectCategoryList(in *cmsclient.QuerySubjectCategoryListReq) (*cmsclient.QuerySubjectCategoryListResp, error) {
	subjectCategory := query.CmsSubjectCategory
	q := subjectCategory.WithContext(l.ctx)
	if len(in.Name) > 0 {
		q = q.Where(subjectCategory.Name.Like("%" + in.Name + "%"))
	}
	if len(in.Icon) > 0 {
		q = q.Where(subjectCategory.Icon.Like("%" + in.Icon + "%"))
	}
	if in.SubjectCount != 2 {
		q = q.Where(subjectCategory.SubjectCount.Eq(in.SubjectCount))
	}
	if in.ShowStatus != 2 {
		q = q.Where(subjectCategory.ShowStatus.Eq(in.ShowStatus))
	}

	result, count, err := q.FindByPage(int((in.PageNum-1)*in.PageSize), int(in.PageSize))

	if err != nil {
		logc.Errorf(l.ctx, "查询专题分类列表失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询专题分类列表失败")
	}

	var list []*cmsclient.SubjectCategoryListData

	for _, item := range result {
		list = append(list, &cmsclient.SubjectCategoryListData{
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

		})
	}

	logc.Infof(l.ctx, "查询专题分类列表信息,参数：%+v,响应：%+v", in, list)

	return &cmsclient.QuerySubjectCategoryListResp{
		Total: count,
		List:  list,
	}, nil
}
