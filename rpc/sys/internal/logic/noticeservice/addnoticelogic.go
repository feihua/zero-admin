package noticeservicelogic

import (
	"context"
	"errors"
	"fmt"

	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

// AddNoticeLogic 添加通知公告
/*
Author: 刘飞华
Date: 2025/10/27 15:49:44
*/
type AddNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoticeLogic {
	return &AddNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddNotice 添加通知公告
func (l *AddNoticeLogic) AddNotice(in *sysclient.AddNoticeReq) (*sysclient.AddNoticeResp, error) {
	q := query.SysNotice

	count, err := q.WithContext(l.ctx).Where(query.SysNotice.NoticeTitle.Eq(in.NoticeTitle)).Count()
	if err != nil {
		logc.Errorf(l.ctx, "根据公告标题：%s,查询公告信息,异常:%s", in.NoticeTitle, err.Error())
		return nil, errors.New(fmt.Sprintf("添加通知公告失败"))
	}
	if count > 0 {
		return nil, errors.New(fmt.Sprintf("添加通知公告失败,公告标题：%s,已存在", in.NoticeTitle))
	}

	item := &model.SysNotice{
		NoticeTitle:   in.NoticeTitle,   // 公告标题
		NoticeType:    in.NoticeType,    // 公告类型（1:通知,2:公告）
		NoticeContent: in.NoticeContent, // 公告内容
		Status:        in.Status,        // 公告状态（0:关闭,1:正常 ）
		Remark:        in.Remark,        // 备注
		CreateBy:      in.CreateBy,      // 创建者

	}

	err = q.WithContext(l.ctx).Create(item)
	if err != nil {
		logc.Errorf(l.ctx, "添加通知公告失败,参数:%+v,异常:%s", item, err.Error())
		return nil, errors.New("添加通知公告失败")
	}

	return &sysclient.AddNoticeResp{}, nil
}
