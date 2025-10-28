package noticeservicelogic

import (
	"context"
	"errors"

	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/pkg/time_util"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// QueryNoticeDetailLogic 查询通知公告详情
/*
Author: 刘飞华
Date: 2025/10/27 15:49:44
*/
type QueryNoticeDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryNoticeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryNoticeDetailLogic {
	return &QueryNoticeDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryNoticeDetail 查询通知公告详情
func (l *QueryNoticeDetailLogic) QueryNoticeDetail(in *sysclient.QueryNoticeDetailReq) (*sysclient.QueryNoticeDetailResp, error) {
	item, err := query.SysNotice.WithContext(l.ctx).Where(query.SysNotice.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "通知公告不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errorx.NewDefaultError("通知公告不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询通知公告异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errorx.NewDefaultError("查询通知公告异常")
	}

	data := &sysclient.QueryNoticeDetailResp{
		Id:            item.ID,                                 // 公告ID
		NoticeTitle:   item.NoticeTitle,                        // 公告标题
		NoticeType:    item.NoticeType,                         // 公告类型（1:通知,2:公告）
		NoticeContent: item.NoticeContent,                      // 公告内容
		Status:        item.Status,                             // 公告状态（0:关闭,1:正常 ）
		Remark:        item.Remark,                             // 备注
		CreateBy:      item.CreateBy,                           // 创建者
		CreateTime:    time_util.TimeToStr(item.CreateTime),    // 创建时间
		UpdateBy:      item.UpdateBy,                           // 更新者
		UpdateTime:    time_util.TimeToString(item.UpdateTime), // 更新时间

	}

	return data, nil
}
