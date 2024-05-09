package albumservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AlbumUpdateLogic 商品相册
/*
Author: LiuFeiHua
Date: 2024/5/8 9:40
*/
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

// AlbumUpdate 修改商品相册
func (l *AlbumUpdateLogic) AlbumUpdate(in *pmsclient.AlbumUpdateReq) (*pmsclient.AlbumUpdateResp, error) {
	q := query.PmsAlbum
	_, err := q.WithContext(l.ctx).Updates(&model.PmsAlbum{
		ID:          in.Id,
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
