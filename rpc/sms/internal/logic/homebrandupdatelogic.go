package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandUpdateLogic {
	return &HomeBrandUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandUpdateLogic) HomeBrandUpdate(in *sms.HomeBrandUpdateReq) (*sms.HomeBrandUpdateResp, error) {
	// todo: add your logic here and delete this line

	return &sms.HomeBrandUpdateResp{}, nil
}
