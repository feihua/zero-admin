package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

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

func (l *AlbumPicUpdateLogic) AlbumPicUpdate(in *pms.AlbumPicUpdateReq) (*pms.AlbumPicUpdateResp, error) {
	err := l.svcCtx.PmsAlbumPicModel.Update(l.ctx, &pmsmodel.PmsAlbumPic{
		Id:      in.Id,
		AlbumId: in.AlbumId,
		Pic:     in.Pic,
	})
	if err != nil {
		return nil, err
	}

	return &pms.AlbumPicUpdateResp{}, nil
}
