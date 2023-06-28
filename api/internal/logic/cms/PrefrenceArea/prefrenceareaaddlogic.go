package PrefrenceArea

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrefrenceAreaAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaAddLogic {
	return &PrefrenceAreaAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrefrenceAreaAddLogic) PrefrenceAreaAdd(req *types.AddPrefrenceAreaReq) (resp *types.AddPrefrenceAreaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
