package logic

import (
	"context"
	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"
	"go-zero-admin/rpc/sys/sysclient"

	"github.com/tal-tech/go-zero/core/logx"
)

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

func (l *LoginLogListLogic) LoginLogList(req types.ListLoginLogReq) (*types.ListLoginLogResp, error) {
	resp, err := l.svcCtx.Sys.LoginLogList(l.ctx, &sysclient.LoginLogListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		return nil, err
	}

	var list []*types.ListLoginLogData

	for _, log := range resp.List {
		list = append(list, &types.ListLoginLogData{
			Id:             log.Id,
			UserName:       log.UserName,
			Status:         log.Status,
			Ip:             log.Ip,
			CreateBy:       log.CreateBy,
			CreateTime:     log.CreateTime,
			LastUpdateBy:   log.LastUpdateBy,
			LastUpdateTime: log.LastUpdateTime,
		})
	}

	return &types.ListLoginLogResp{
		Code:     "000000",
		Message:  "",
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
	}, nil
}
