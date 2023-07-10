package albumpicservicelogic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"
	"zero-admin/rpc/pms/pmsclient"

	"zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumPicAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumPicAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumPicAddLogic {
	return &AlbumPicAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumPicAddLogic) AlbumPicAdd(in *pmsclient.AlbumPicAddReq) (*pmsclient.AlbumPicAddResp, error) {
	_, err := l.svcCtx.PmsAlbumPicModel.Insert(l.ctx, &pmsmodel.PmsAlbumPic{
		AlbumId: in.AlbumId,
		Pic:     in.Pic,
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumPicAddResp{}, nil
}
