package operatelogservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOperateLogDetailLogic 查询系统操作日志详情
/*
Author: LiuFeiHua
Date: 2024/5/30 11:11
*/
type QueryOperateLogDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryOperateLogDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryOperateLogDetailLogic {
	return &QueryOperateLogDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// QueryOperateLogDetail 查询系统操作日志详情
func (l *QueryOperateLogDetailLogic) QueryOperateLogDetail(in *sysclient.QueryOperateLogDetailReq) (*sysclient.QueryOperateLogDetailResp, error) {
	log, err := query.SysOperateLog.WithContext(l.ctx).Where(query.SysOperateLog.ID.Eq(in.Id)).First()

	if err != nil {
		logc.Errorf(l.ctx, "查询系统操作日志详情失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("查询系统操作日志详情失败")
	}
	data := &sysclient.QueryOperateLogDetailResp{
		DeptName:          log.DeptName,
		Id:                log.ID,
		OperationIp:       log.OperationIP,
		OperationName:     log.OperationName,
		OperationParams:   log.OperationParams,
		OperationResponse: log.OperationResponse,
		OperationStatus:   log.OperationStatus,
		OperationTime:     log.OperationTime.Format("2006-01-02 15:04:05"),
		OperationType:     log.OperationType,
		OperationUrl:      log.OperationURL,
		RequestMethod:     log.RequestMethod,
		Title:             log.Title,
		UseTime:           log.UseTime,
		Browser:           log.Browser,
		Os:                log.Os,
	}

	logc.Infof(l.ctx, "查询系统操作日志详情,参数：%+v,响应：%+v", in, data)
	return data, nil
}
