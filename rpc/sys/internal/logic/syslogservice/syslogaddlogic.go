package syslogservicelogic

import (
	"context"
	"database/sql"
	"time"
	"zero-admin/rpc/model/sysmodel"
	"zero-admin/rpc/sys/sysclient"

	"zero-admin/rpc/sys/internal/svc"

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

func (l *SysLogAddLogic) SysLogAdd(in *sysclient.SysLogAddReq) (*sysclient.SysLogAddResp, error) {
	_, err := l.svcCtx.SysLogModel.Insert(l.ctx, &sysmodel.SysLog{
		UserName:       in.UserName,
		Operation:      in.Operation,
		Method:         in.Method,
		RequestParams:  in.RequestParams,
		ResponseParams: sql.NullString{String: in.ResponseParams, Valid: true},
		Time:           in.Time,
		Ip:             sql.NullString{String: in.Ip, Valid: true},
		OperationTime:  time.Now(),
	})

	if err != nil {
		return nil, err
	}

	return &sysclient.SysLogAddResp{}, nil
}
