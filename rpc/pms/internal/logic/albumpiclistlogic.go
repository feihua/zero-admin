package logic

import (
	"context"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlbumPicListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumPicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumPicListLogic {
	return &AlbumPicListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumPicListLogic) AlbumPicList(in *pms.AlbumPicListReq) (*pms.AlbumPicListResp, error) {
	// todo: add your logic here and delete this line

	return &pms.AlbumPicListResp{}, nil
}
