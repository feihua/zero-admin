package albumpicservicelogic

import (
	"context"
	"github.com/feihua/zero-admin/rpc/pms/gen/query"
	"github.com/feihua/zero-admin/rpc/pms/pmsclient"

	"github.com/feihua/zero-admin/rpc/pms/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

// AlbumPicDeleteLogic 相册表图片
/*
Author: LiuFeiHua
Date: 2024/5/8 9:35
*/
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

// AlbumPicDelete 删除相册图片
func (l *AlbumPicDeleteLogic) AlbumPicDelete(in *pmsclient.AlbumPicDeleteReq) (*pmsclient.AlbumPicDeleteResp, error) {
	q := query.PmsAlbumPic
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()

	if err != nil {
		return nil, err
	}

	return &pmsclient.AlbumPicDeleteResp{}, nil
}
