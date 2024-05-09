package albumpicservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/model"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AlbumPicUpdateLogic 相册表图片
/*
Author: LiuFeiHua
Date: 2024/5/8 10:28
*/
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

// AlbumPicUpdate 更新相册表图片
func (l *AlbumPicUpdateLogic) AlbumPicUpdate(in *pmsclient.AlbumPicUpdateReq) (*pmsclient.AlbumPicUpdateResp, error) {
	q := query.PmsAlbumPic
	_, err := q.WithContext(l.ctx).Updates(&model.PmsAlbumPic{
		ID:      in.Id,
		AlbumID: in.AlbumId,
		Pic:     in.Pic,
	})

	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumPicUpdateResp{}, nil
}
