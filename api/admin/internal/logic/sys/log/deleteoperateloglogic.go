package log

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// DeleteOperateLogLogic 删除操作日志
/*
Author: LiuFeiHua
Date: 2023/12/18 17:20
*/
type DeleteOperateLogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOperateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteOperateLogLogic {
	return DeleteOperateLogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeleteOperateLog 删除操作日志
func (l *DeleteOperateLogLogic) DeleteOperateLog(req *types.DeleteOperateLogReq) (*types.DeleteOperateLogResp, error) {
	_, err := l.svcCtx.Operatelogservice.DeleteOperateLog(l.ctx, &sysclient.DeleteOperateLogReq{
		Ids: req.Ids,
	})

	if err != nil {
		logc.Errorf(l.ctx, "根据OperateLog ids: %+v,删除操作日志异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	return &types.DeleteOperateLogResp{
		Code:    "000000",
		Message: "删除操作日志成功",
	}, nil
}
