package albumpicservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AlbumPicAddLogic 相册表图片
/*
Author: LiuFeiHua
Date: 2024/5/8 10:24
*/
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

// AlbumPicAdd 添加相册表图片
func (l *AlbumPicAddLogic) AlbumPicAdd(in *pmsclient.AlbumPicAddReq) (*pmsclient.AlbumPicAddResp, error) {
	err := query.PmsAlbumPic.WithContext(l.ctx).Create(&model.PmsAlbumPic{
		AlbumID: in.AlbumId,
		Pic:     in.Pic,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumPicAddResp{}, nil
}
