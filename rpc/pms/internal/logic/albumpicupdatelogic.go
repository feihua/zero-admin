package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
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
	err := l.svcCtx.PmsAlbumPicModel.Update(pmsmodel.PmsAlbumPic{
		Id:      in.Id,
		AlbumId: in.AlbumId,
		Pic:     in.Pic,
	})
	if err != nil {
		return nil, err
	}

	return &pms.AlbumPicUpdateResp{}, nil
}
