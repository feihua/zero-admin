package log

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

// QueryLoginLogListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:20
*/
type QueryLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) QueryLoginLogListLogic {
	return QueryLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// QueryLoginLogList 登录日志列表
func (l *QueryLoginLogListLogic) QueryLoginLogList(req *types.ListLoginLogReq) (*types.ListLoginLogResp, error) {
	resp, err := l.svcCtx.LoginLogService.LoginLogList(l.ctx, &sysclient.LoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserName: req.UserName,
		Ip:       req.Ip,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询登录日志列表异常:%s", req, err.Error())
		s, _ := status.FromError(err)
		return nil, errorx.NewDefaultError(s.Message())
	}

	var list []*types.ListLoginLogData

	for _, log := range resp.List {
		list = append(list, &types.ListLoginLogData{
			Id:         log.Id,
			UserName:   log.UserName,
			Status:     log.Status,
			Ip:         log.Ip,
			CreateBy:   log.CreateBy,
			CreateTime: log.CreateTime,
			UpdateBy:   log.UpdateBy,
			UpdateTime: log.UpdateTime,
		})
	}

	return &types.ListLoginLogResp{
		Code:     "000000",
		Message:  "查询登录日志成功",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
