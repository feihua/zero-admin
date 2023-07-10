package albumpicservicelogic

import (
	"context"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumPicDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumPicDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumPicDeleteLogic {
	return &AlbumPicDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumPicDeleteLogic) AlbumPicDelete(in *pmsclient.AlbumPicDeleteReq) (*pmsclient.AlbumPicDeleteResp, error) {
	err := l.svcCtx.PmsAlbumPicModel.DeleteByIds(l.ctx, in.Ids)

	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumPicDeleteResp{}, nil
}
