package sys_notice

import (
	"context"

	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryNoticeListLogic 查询通知公告表列表
/*
Author: 刘飞华
Date: 2025/10/27 15:51:14
*/
type QueryNoticeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryNoticeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryNoticeListLogic {
	return &QueryNoticeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryNoticeList 查询通知公告表列表
func (l *QueryNoticeListLogic) QueryNoticeList(req *types.QueryNoticeListReq) (resp *types.QueryNoticeListResp, err error) {
	result, err := l.svcCtx.NoticeService.QueryNoticeList(l.ctx, &sysclient.QueryNoticeListReq{
		PageNum:     req.Current,
		PageSize:    req.PageSize,
		NoticeTitle: req.NoticeTitle, // 公告标题
		NoticeType:  req.NoticeType,  // 公告类型（1:通知,2:公告）
		Status:      req.Status,      // 公告状态（0:关闭,1:正常 ）

	})

	if err != nil {
		logc.Errorf(l.ctx, "查询字通知公告表列表失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.QueryNoticeListData

	for _, item := range result.List {
		list = append(list, &types.QueryNoticeListData{
			Id:            item.Id,            // 公告ID
			NoticeTitle:   item.NoticeTitle,   // 公告标题
			NoticeType:    item.NoticeType,    // 公告类型（1:通知,2:公告）
			NoticeContent: item.NoticeContent, // 公告内容
			Status:        item.Status,        // 公告状态（0:关闭,1:正常 ）
			Remark:        item.Remark,        // 备注
			CreateBy:      item.CreateBy,      // 创建者
			CreateTime:    item.CreateTime,    // 创建时间
			UpdateBy:      item.UpdateBy,      // 更新者
			UpdateTime:    item.UpdateTime,    // 更新时间
		})
	}

	return &types.QueryNoticeListResp{
		Code:     "000000",
		Message:  "查询通知公告表列表成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    result.Total,
	}, nil
}
