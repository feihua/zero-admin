package log

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryOperateLogListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:20
*/
type QueryOperateLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryOperateLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryOperateLogListLogic {
	return QueryOperateLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryOperateLogList 操作日志列表
func (l *QueryOperateLogListLogic) QueryOperateLogList(req *types.ListSysLogReq) (*types.ListSysLogResp, error) {
	resp, err := l.svcCtx.SysLogService.SysLogList(l.ctx, &sysclient.SysLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserName: req.UserName,
		Method:   req.Method,
	})

	if err != nil {
		return nil, errorx.NewDefaultError("查询操作日志失败")
	}

	var list []*types.ListSysLogData

	for _, log := range resp.List {
		list = append(list, &types.ListSysLogData{
			Id:             log.Id,
			UserName:       log.UserName,
			Operation:      log.Operation,
			Method:         log.Method,
			RequestParams:  log.RequestParams,
			Time:           log.Time,
			Ip:             log.Ip,
			ResponseParams: log.ResponseParams,
			OperationTime:  log.OperationTime,
		})
	}

	return &types.ListSysLogResp{
		Code:     "000000",
		Message:  "查询操作日志列表",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil

}
