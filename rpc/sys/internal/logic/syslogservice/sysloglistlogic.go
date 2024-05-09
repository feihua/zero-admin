package syslogservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/feihua/zero-admin/rpc/sys/internal/svc"

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
	q := query.SysLog.WithContext(l.ctx)
	if len(in.UserName) > 0 {
		q = q.Where(query.SysLog.UserName.Like("%" + in.UserName + "%"))
	}
	if len(in.Method) > 0 {
		q = q.Where(query.SysLog.UserName.Like("%" + in.Method + "%"))
	}

	offset := (in.Current - 1) * in.PageSize
	result, err := q.Offset(int(offset)).Limit(int(in.PageSize)).Find()
	count, err := q.Count()

	if err != nil {
		return nil, err
	}
	var list []*sysclient.SysLogListData
	for _, log := range result {
		list = append(list, &sysclient.SysLogListData{
			Id:             log.ID,
			UserName:       log.UserName,
			Operation:      log.Operation,
			Method:         log.Method,
			RequestParams:  log.RequestParams,
			Time:           log.Time,
			Ip:             *log.IP,
			ResponseParams: *log.ResponseParams,
			OperationTime:  log.OperationTime.Format("2006-01-02 15:04:05"),
		})
	}

	return &sysclient.SysLogListResp{
		Total: count,
		List:  list,
	}, nil

}
