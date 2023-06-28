package PrefrenceArea

import (
	"context"

	"zero-admin/api/internal/svc"
	"zero-admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PrefrenceAreaDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPrefrenceAreaDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PrefrenceAreaDeleteLogic {
	return &PrefrenceAreaDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PrefrenceAreaDeleteLogic) PrefrenceAreaDelete(req *types.DeletePrefrenceAreaReq) (resp *types.DeletePrefrenceAreaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
