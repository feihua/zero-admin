package loginlogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteLoginLogLogic 删除登录日志
/*
Author: LiuFeiHua
Date: 2023/12/18 17:07
*/
type DeleteLoginLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLoginLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLoginLogLogic {
	return &DeleteLoginLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeleteLoginLog 删除登录日志
func (l *DeleteLoginLogLogic) DeleteLoginLog(in *sysclient.DeleteLoginLogReq) (*sysclient.DeleteLoginLogResp, error) {
	q := query.SysLoginLog
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除登录日志失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除登录日志失败")
	}

	return &sysclient.DeleteLoginLogResp{}, nil
}
