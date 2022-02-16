package logic

import (
	"context"
	"zero-admin/rpc/model/pmsmodel"

	"zero-admin/rpc/pms/internal/svc"
	"zero-admin/rpc/pms/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type AlbumAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumAddLogic {
	return &AlbumAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumAddLogic) AlbumAdd(in *pms.AlbumAddReq) (*pms.AlbumAddResp, error) {
	_, err := l.svcCtx.PmsAlbumModel.Insert(pmsmodel.PmsAlbum{
		Name:        in.Name,
		CoverPic:    in.CoverPic,
		PicCount:    in.PicCount,
		Sort:        in.Sort,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pms.AlbumAddResp{}, nil
}
