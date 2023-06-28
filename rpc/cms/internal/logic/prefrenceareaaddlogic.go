package logic

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaAddLogic {
	return &PrefrenceAreaAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品优选
func (l *PrefrenceAreaAddLogic) PrefrenceAreaAdd(in *cmsclient.PrefrenceAreaAddReq) (*cmsclient.PrefrenceAreaAddResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.PrefrenceAreaAddResp{}, nil
}
