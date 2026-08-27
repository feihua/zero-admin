package help_category

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

// QueryHelpCategoryListLogic 查询帮助分类列表
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryHelpCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHelpCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpCategoryListLogic {
	return &QueryHelpCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryHelpCategoryList 查询帮助分类列表
func (l *QueryHelpCategoryListLogic) QueryHelpCategoryList(req *types.QueryHelpCategoryListReq) (resp *types.QueryHelpCategoryListResp, err error) {
	result, err := l.svcCtx.HelpCategoryService.QueryHelpCategoryList(l.ctx, &cmsclient.QueryHelpCategoryListReq{
		PageNum:    int32(req.Current),
		PageSize:   int32(req.PageSize),
		Name:       req.Name,       // 分类名称
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字帮助分类列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryHelpCategoryListData

	for _, item := range result.List {
		list = append(list, &types.QueryHelpCategoryListData{
			Id:         item.Id,         // 主键id
			Name:       item.Name,       // 分类名称
			Icon:       item.Icon,       // 分类图标
			HelpCount:  item.HelpCount,  // 专题数量
			ShowStatus: item.ShowStatus, // 显示状态：0->不显示；1->显示
			Sort:       item.Sort,       // 排序
			CreateBy:   item.CreateBy,   // 创建者
			CreateTime: item.CreateTime, // 创建时间
			UpdateBy:   item.UpdateBy,   // 更新者
			UpdateTime: item.UpdateTime, // 更新时间
		})
	}

	return &types.QueryHelpCategoryListResp{
		Code:     "000000",
		Message:  "查询帮助分类列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
