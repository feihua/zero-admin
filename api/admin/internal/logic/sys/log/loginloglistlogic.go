package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// LoginLogListLogic
/*
Author: LiuFeiHua
Date: 2023/12/18 17:20
*/
type LoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogListLogic {
	return LoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// LoginLogList 登录日志列表
func (l *LoginLogListLogic) LoginLogList(req types.ListLoginLogReq) (*types.ListLoginLogResp, error) {
	resp, err := l.svcCtx.LoginLogService.LoginLogList(l.ctx, &sysclient.LoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		UserName: req.UserName,
		Ip:       req.Ip,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询登录日志列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询登录日志失败")
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
