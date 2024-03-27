package albumservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/model/pmsmodel"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *AlbumUpdateLogic) AlbumUpdate(in *pmsclient.AlbumUpdateReq) (*pmsclient.AlbumUpdateResp, error) {
	err := l.svcCtx.PmsAlbumModel.Update(l.ctx, &pmsmodel.PmsAlbum{
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

	return &pmsclient.AlbumUpdateResp{}, nil
}
