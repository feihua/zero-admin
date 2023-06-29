package logic

import (
	"context"

	"zero-admin/rpc/cms/cmsclient"
	"zero-admin/rpc/cms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaProductRelationListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPrefrenceAreaProductRelationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaProductRelationListLogic {
	return &PrefrenceAreaProductRelationListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PrefrenceAreaProductRelationListLogic) PrefrenceAreaProductRelationList(in *cmsclient.PrefrenceAreaProductRelationListReq) (*cmsclient.PrefrenceAreaProductRelationListResp, error) {
	// todo: add your logic here and delete this line

	return &cmsclient.PrefrenceAreaProductRelationListResp{}, nil
}
