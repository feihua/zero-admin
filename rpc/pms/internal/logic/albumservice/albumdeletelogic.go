package albumservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumDeleteLogic {
	return &AlbumDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumDeleteLogic) AlbumDelete(in *pmsclient.AlbumDeleteReq) (*pmsclient.AlbumDeleteResp, error) {
	err := l.svcCtx.PmsAlbumModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumDeleteResp{}, nil
}
