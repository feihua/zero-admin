package syslogservicelogic

import (
	"context"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// SysLogListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:09
*/
type SysLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SysLogList 操作日志列表
func (l *SysLogListLogic) SysLogList(in *sysclient.SysLogListReq) (*sysclient.SysLogListResp, error) {
	all, err := l.svcCtx.SysLogModel.FindAll(l.ctx, in)
	count, _ := l.svcCtx.SysLogModel.Count(l.ctx, in)

	if err != nil {
		return nil, err
	}
	var list []*sysclient.SysLogListData
	for _, log := range *all {
		list = append(list, &sysclient.SysLogListData{
			Id:             log.Id,
			UserName:       log.UserName,
			Operation:      log.Operation,
			Method:         log.Method,
			RequestParams:  log.RequestParams,
			Time:           log.Time,
			Ip:             log.Ip.String,
			ResponseParams: log.ResponseParams.String,
			OperationTime:  log.OperationTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &sysclient.SysLogListResp{
		Total: count,
		List:  list,
	}, nil

}
