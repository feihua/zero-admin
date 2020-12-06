package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumListLogic {
	return &AlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumListLogic) AlbumList(in *pms.AlbumListReq) (*pms.AlbumListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.AlbumListResp{}, nil
}
