package operatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/model"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"time"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AddOperateLogLogic 添加操作日志
/*
Author: LiuFeiHua
Date: 2023/12/18 17:08
*/
type AddOperateLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOperateLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOperateLogLogic {
	return &AddOperateLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddOperateLog 添加操作日志
func (l *AddOperateLogLogic) AddOperateLog(in *sysclient.AddOperateLogReq) (*sysclient.AddOperateLogResp, error) {
	sysLog := &model.SysOperateLog{
		Title:             in.Title,
		OperationType:     in.OperationType,
		OperationName:     in.OperationName,
		RequestMethod:     in.RequestMethod,
		OperationURL:      in.OperationUrl,
		OperationParams:   in.OperationParams,
		OperationResponse: in.OperationResponse,
		OperationStatus:   in.OperationStatus,
		DeptName:          in.DeptName,
		UseTime:           in.UseTime,
		Browser:           in.Browser,
		Os:                in.Os,
		OperationIP:       in.OperationIp,
		OperationTime:     time.Now(),
	}

	err := query.SysOperateLog.WithContext(l.ctx).Create(sysLog)
	if err != nil {
		logc.Errorf(l.ctx, "添加操作日志失败,参数:%+v,异常:%s", sysLog, err.Error())
		return nil, errors.New("添加操作日志失败")
	}

	return &sysclient.AddOperateLogResp{}, nil
}
