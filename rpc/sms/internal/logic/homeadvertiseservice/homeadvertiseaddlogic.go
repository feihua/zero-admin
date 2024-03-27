package homeadvertiseservicelogic

import (
	"context"
	"database/sql"
	"github.com/feihua/zero-admin/rpc/model/smsmodel"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"time"

	"github.com/feihua/zero-admin/rpc/sms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseAddLogic {
	return &HomeAdvertiseAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(in *smsclient.HomeAdvertiseAddReq) (*smsclient.HomeAdvertiseAddResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	note := sql.NullString{String: in.Note, Valid: true}
	_, err := l.svcCtx.SmsHomeAdvertiseModel.Insert(l.ctx, &smsmodel.SmsHomeAdvertise{
		Name:       in.Name,
		Type:       in.Type,
		Pic:        in.Pic,
		StartTime:  StartTime,
		EndTime:    EndTime,
		Status:     in.Status,
		ClickCount: in.ClickCount,
		OrderCount: in.OrderCount,
		Url:        in.Url,
		Note:       note,
		Sort:       in.Sort,
	})
	if err != nil {
		return nil, err
	}

	return &smsclient.HomeAdvertiseAddResp{}, nil
}
