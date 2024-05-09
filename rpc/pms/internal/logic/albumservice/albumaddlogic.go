package albumservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AlbumAddLogic 商品相册
/*
Author: LiuFeiHua
Date: 2024/5/8 9:40
*/
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

// AlbumAdd 添加商品相册
func (l *AlbumAddLogic) AlbumAdd(in *pmsclient.AlbumAddReq) (*pmsclient.AlbumAddResp, error) {
	err := query.PmsAlbum.WithContext(l.ctx).Create(&model.PmsAlbum{
		Name:        in.Name,
		CoverPic:    in.CoverPic,
		PicCount:    in.PicCount,
		Sort:        in.Sort,
		Description: in.Description,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumAddResp{}, nil
}
