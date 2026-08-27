package help

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

// QueryHelpListLogic 查询帮助列表
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryHelpListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHelpListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpListLogic {
	return &QueryHelpListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryHelpList 查询帮助列表
func (l *QueryHelpListLogic) QueryHelpList(req *types.QueryHelpListReq) (resp *types.QueryHelpListResp, err error) {
	result, err := l.svcCtx.HelpService.QueryHelpList(l.ctx, &cmsclient.QueryHelpListReq{
		PageNum:    int32(req.Current),
		PageSize:   int32(req.PageSize),
		CategoryId: req.CategoryId, // 分类id
		Title:      req.Title,      // 标题
		ShowStatus: req.ShowStatus, // 显示状态：0->不显示；1->显示

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字帮助列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryHelpListData

	for _, item := range result.List {
		list = append(list, &types.QueryHelpListData{
			Id:         item.Id,         // 主键id
			CategoryId: item.CategoryId, // 分类id
			Icon:       item.Icon,       // 图标
			Title:      item.Title,      // 标题
			ShowStatus: item.ShowStatus, // 显示状态：0->不显示；1->显示
			ReadCount:  item.ReadCount,  // 阅读量
			Content:    item.Content,    // 内容
			CreateBy:   item.CreateBy,   // 创建者
			CreateTime: item.CreateTime, // 创建时间
			UpdateBy:   item.UpdateBy,   // 更新者
			UpdateTime: item.UpdateTime, // 更新时间
		})
	}

	return &types.QueryHelpListResp{
		Code:     "000000",
		Message:  "查询帮助列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
