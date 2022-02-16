package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

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

func (l *AlbumPicAddLogic) AlbumPicAdd(in *pms.AlbumPicAddReq) (*pms.AlbumPicAddResp, error) {
	_, err := l.svcCtx.PmsAlbumPicModel.Insert(pmsmodel.PmsAlbumPic{
		AlbumId: in.AlbumId,
		Pic:     in.Pic,
	})
	if err != nil {
		return nil, err
	}

	return &pms.AlbumPicAddResp{}, nil
}
