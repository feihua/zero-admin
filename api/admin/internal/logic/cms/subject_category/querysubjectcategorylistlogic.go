package subject_category

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

// QuerySubjectCategoryListLogic 查询专题分类表列表
/*
Author: 刘飞华
Date: 2025/02/07 10:21:09
*/
type QuerySubjectCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuerySubjectCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuerySubjectCategoryListLogic {
	return &QuerySubjectCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QuerySubjectCategoryList 查询专题分类表列表
func (l *QuerySubjectCategoryListLogic) QuerySubjectCategoryList(req *types.QuerySubjectCategoryListReq) (resp *types.QuerySubjectCategoryListResp, err error) {
	result, err := l.svcCtx.SubjectCategoryService.QuerySubjectCategoryList(l.ctx, &cmsclient.QuerySubjectCategoryListReq{
		PageNum:    req.Current,
		PageSize:   req.PageSize,
		Name:       req.Name,       // 专题分类名称
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字专题分类表列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QuerySubjectCategoryListData

	for _, item := range result.List {
		list = append(list, &types.QuerySubjectCategoryListData{
			Id:           item.Id,           // 主键ID
			Name:         item.Name,         // 专题分类名称
			Icon:         item.Icon,         // 分类图标
			SubjectCount: item.SubjectCount, // 专题数量
			ShowStatus:   item.ShowStatus,   // 显示状态：0->不显示；1->显示
			Sort:         item.Sort,         // 排序
			CreateBy:     item.CreateBy,     // 创建者
			CreateTime:   item.CreateTime,   // 创建时间
			UpdateBy:     item.UpdateBy,     // 更新者
			UpdateTime:   item.UpdateTime,   // 更新时间
		})
	}

	return &types.QuerySubjectCategoryListResp{
		Code:     "000000",
		Message:  "查询专题分类表列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
