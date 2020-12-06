package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandDeleteLogic {
	return &HomeBrandDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandDeleteLogic) HomeBrandDelete(in *sms.HomeBrandDeleteReq) (*sms.HomeBrandDeleteResp, error) {
	// todo: add your logic here and delete this line

	return &sms.HomeBrandDeleteResp{}, nil
}
