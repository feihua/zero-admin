package logic

import (
	"context"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

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

func (l *AlbumPicDeleteLogic) AlbumPicDelete(in *pms.AlbumPicDeleteReq) (*pms.AlbumPicDeleteResp, error) {
	err := l.svcCtx.PmsAlbumPicModel.Delete(in.Id)

	if err != nil {
		return nil, err
	}

	return &pms.AlbumPicDeleteResp{}, nil
}
