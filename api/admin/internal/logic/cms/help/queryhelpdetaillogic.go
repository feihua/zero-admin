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

// QueryHelpDetailLogic 查询帮助详情
/*
Author: 刘飞华
Date: 2026/08/27 10:50:29
*/
type QueryHelpDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHelpDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHelpDetailLogic {
	return &QueryHelpDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryHelpDetail 查询帮助详情
func (l *QueryHelpDetailLogic) QueryHelpDetail(req *types.QueryHelpDetailReq) (resp *types.QueryHelpDetailResp, err error) {

	detail, err := l.svcCtx.HelpService.QueryHelpDetail(l.ctx, &cmsclient.QueryHelpDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询帮助详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryHelpDetailData{
		Id:         detail.Id,         // 主键id
		CategoryId: detail.CategoryId, // 分类id
		Icon:       detail.Icon,       // 图标
		Title:      detail.Title,      // 标题
		ShowStatus: detail.ShowStatus, // 显示状态：0->不显示；1->显示
		ReadCount:  detail.ReadCount,  // 阅读量
		Content:    detail.Content,    // 内容
		CreateBy:   detail.CreateBy,   // 创建者
		CreateTime: detail.CreateTime, // 创建时间
		UpdateBy:   detail.UpdateBy,   // 更新者
		UpdateTime: detail.UpdateTime, // 更新时间
	}
	return &types.QueryHelpDetailResp{
		Code:    "000000",
		Message: "查询帮助成功",
		Data:    data,
	}, nil
}
