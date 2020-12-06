package logic

import (
	"context"

	"go-zero-admin/rpc/sms/internal/svc"
	"go-zero-admin/rpc/sms/sms"

	"github.com/tal-tech/go-zero/core/logx"
)

type HomeBrandListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeBrandListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeBrandListLogic {
	return &HomeBrandListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeBrandListLogic) HomeBrandList(in *sms.HomeBrandListReq) (*sms.HomeBrandListResp, error) {
	// todo: add your logic here and delete this line

	return &sms.HomeBrandListResp{}, nil
}
