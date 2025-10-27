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

// QueryNoticeDetailLogic 查询通知公告表详情
/*
Author: 刘飞华
Date: 2025/10/27 15:51:14
*/
type QueryNoticeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryNoticeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryNoticeDetailLogic {
	return &QueryNoticeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryNoticeDetail 查询通知公告表详情
func (l *QueryNoticeDetailLogic) QueryNoticeDetail(req *types.QueryNoticeDetailReq) (resp *types.QueryNoticeDetailResp, err error) {

	detail, err := l.svcCtx.NoticeService.QueryNoticeDetail(l.ctx, &sysclient.QueryNoticeDetailReq{
		Id: req.Id,
	})

	if err != nil {
		logc.Errorf(l.ctx, "查询通知公告表详情失败,参数：%+v,响应：%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	data := types.QueryNoticeDetailData{
		Id:            detail.Id,            // 公告ID
		NoticeTitle:   detail.NoticeTitle,   // 公告标题
		NoticeType:    detail.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: detail.NoticeContent, // 公告内容
		Status:        detail.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        detail.Remark,        // 备注
		CreateBy:      detail.CreateBy,      // 创建者
		CreateTime:    detail.CreateTime,    // 创建时间
		UpdateBy:      detail.UpdateBy,      // 更新者
		UpdateTime:    detail.UpdateTime,    // 更新时间
	}
	return &types.QueryNoticeDetailResp{
		Code:    "000000",
		Message: "查询通知公告表成功",
		Data:    data,
	}, nil
}
