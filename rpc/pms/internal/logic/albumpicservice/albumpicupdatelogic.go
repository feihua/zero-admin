package albumpicservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/pmsmodel"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumPicUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumPicUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumPicUpdateLogic {
	return &AlbumPicUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumPicUpdateLogic) AlbumPicUpdate(in *pmsclient.AlbumPicUpdateReq) (*pmsclient.AlbumPicUpdateResp, error) {
	err := l.svcCtx.PmsAlbumPicModel.Update(l.ctx, &pmsmodel.PmsAlbumPic{
		Id:      in.Id,
		AlbumId: in.AlbumId,
		Pic:     in.Pic,
	})
	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumPicUpdateResp{}, nil
}
