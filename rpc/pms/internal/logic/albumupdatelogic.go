package logic

import (
	"context"
	"go-zero-admin/rpc/model/pmsmodel"

	"go-zero-admin/rpc/pms/internal/svc"
	"go-zero-admin/rpc/pms/pms"

	"github.com/tal-tech/go-zero/core/logx"
)

type AlbumUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumUpdateLogic {
	return &AlbumUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AlbumUpdateLogic) AlbumUpdate(in *pms.AlbumUpdateReq) (*pms.AlbumUpdateResp, error) {
	err := l.svcCtx.PmsAlbumModel.Update(pmsmodel.PmsAlbum{
		Id:          in.Id,
		Name:        in.Name,
		CoverPic:    in.CoverPic,
		PicCount:    in.PicCount,
		Sort:        in.Sort,
		Description: in.Description,
	})
	if err != nil {
		return nil, err
	}

	return &pms.AlbumUpdateResp{}, nil
}
