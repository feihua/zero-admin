package albumservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AlbumDeleteLogic 商品相册
/*
Author: LiuFeiHua
Date: 2024/5/8 9:40
*/
type AlbumDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAlbumDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AlbumDeleteLogic {
	return &AlbumDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AlbumDelete 删除商品相册
func (l *AlbumDeleteLogic) AlbumDelete(in *pmsclient.AlbumDeleteReq) (*pmsclient.AlbumDeleteResp, error) {
	q := query.PmsAlbum
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumDeleteResp{}, nil
}
