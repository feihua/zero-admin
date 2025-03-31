package postservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdatePostStatusLogic 更新岗位管理状态
/*
Author: LiuFeiHua
Date: 2024/5/30 11:20
*/
type UpdatePostStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostStatusLogic {
	return &UpdatePostStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdatePostStatus 更新岗位管理状态
func (l *UpdatePostStatusLogic) UpdatePostStatus(in *sysclient.UpdatePostStatusReq) (*sysclient.UpdatePostStatusResp, error) {
	q := query.SysPost

	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Update(q.Status, in.Status)

	if err != nil {
		logc.Errorf(l.ctx, "更新岗位管理状态失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("更新岗位管理状态失败")
	}

	return &sysclient.UpdatePostStatusResp{}, nil
}
