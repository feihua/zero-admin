package logic

import (
	"context"
	"zero-admin/api/internal/common/errorx"
	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"
	"zero-admin/rpc/sys/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SysLogListLogic {
	return SysLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogListLogic) SysLogList(req types.ListSysLogReq) (*types.ListSysLogResp, error) {
	resp, err := l.svcCtx.Sys.SysLogList(l.ctx, &sysclient.SysLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
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
			Params:         log.Params,
			Time:           log.Time,
			Ip:             log.Ip,
			CreateBy:       log.CreateBy,
			CreateTime:     log.CreateTime,
			LastUpdateBy:   log.LastUpdateBy,
			LastUpdateTime: log.LastUpdateTime,
		})
	}

	return &types.ListSysLogResp{
		Code:     "000000",
		Message:  "",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil

}
