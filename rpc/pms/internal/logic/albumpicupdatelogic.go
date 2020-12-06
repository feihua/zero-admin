package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pms.AlbumPicUpdateResp{}, nil
}
