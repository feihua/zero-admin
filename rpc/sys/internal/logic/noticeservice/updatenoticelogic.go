package noticeservicelogic

import (
	"context"
	"errors"
	"fmt"

	"github.com/feihua/zero-admin/pkg/errorx"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// UpdateNoticeLogic 更新通知公告
/*
Author: 刘飞华
Date: 2025/10/27 15:49:44
*/
type UpdateNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogic {
	return &UpdateNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateNotice 更新通知公告
func (l *UpdateNoticeLogic) UpdateNotice(in *sysclient.UpdateNoticeReq) (*sysclient.UpdateNoticeResp, error) {
	notice := query.SysNotice
	q := notice.WithContext(l.ctx)

	// 1.根据通知公告id查询通知公告是否已存在
	_, err := q.Where(notice.ID.Eq(in.Id)).First()

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		logc.Errorf(l.ctx, "通知公告不存在, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errorx.NewDefaultError("通知公告不存在")
	case err != nil:
		logc.Errorf(l.ctx, "查询通知公告异常, 请求参数：%+v, 异常信息: %s", in, err.Error())
		return nil, errorx.NewDefaultError("查询通知公告异常")
	}

	count, err := q.WithContext(l.ctx).Where(notice.NoticeTitle.Eq(in.NoticeTitle), notice.ID.Neq(in.Id)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "根据公告标题：%s,查询公告信息,异常:%s", in.NoticeTitle, err.Error())
		return nil, errors.New(fmt.Sprintf("添加通知公告失败"))
	}
	if count > 0 {
		return nil, errors.New(fmt.Sprintf("添加通知公告失败,公告标题：%s,已存在", in.NoticeTitle))
	}
	item := &model.SysNotice{
		ID:            in.Id,            // 公告ID
		NoticeTitle:   in.NoticeTitle,   // 公告标题
		NoticeType:    in.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: in.NoticeContent, // 公告内容
		Status:        in.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        in.Remark,        // 备注
		UpdateBy:      in.UpdateBy,      // 更新者

	}

	// 2.通知公告存在时,则直接更新通知公告
	_, err = q.Updates(item)

	if err != nil {
		logc.Errorf(l.ctx, "更新通知公告失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("更新通知公告失败")
	}

	return &sysclient.UpdateNoticeResp{}, nil
}
