package logic

import (
	"context"
	"time"
	"zero-admin/rpc/model/sysmodel"

	"zero-admin/rpc/sys/internal/svc"
	"zero-admin/rpc/sys/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogAddLogic {
	return &SysLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLogAddLogic) SysLogAdd(in *sys.SysLogAddReq) (*sys.SysLogAddResp, error) {
	_, err := l.svcCtx.SysLogModel.Insert(sysmodel.SysLog{
		UserName:       in.UserName,
		Operation:      in.Operation,
		Method:         in.Method,
		Params:         in.Params,
		Time:           in.Time,
		Ip:             in.Ip,
		CreateBy:       in.CreateBy,
		LastUpdateBy:   in.CreateBy,
		LastUpdateTime: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &sys.SysLogAddResp{}, nil
}
